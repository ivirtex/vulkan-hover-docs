# vkBindImageMemory(3) Manual Page

## Name

vkBindImageMemory - Bind device memory to an image object



## <a href="#_c_specification" class="anchor"></a>C Specification

To attach memory to a `VkImage` object created without the
`VK_IMAGE_CREATE_DISJOINT_BIT` set, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkBindImageMemory(
    VkDevice                                    device,
    VkImage                                     image,
    VkDeviceMemory                              memory,
    VkDeviceSize                                memoryOffset);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the image and memory.

- `image` is the image.

- `memory` is the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object
  describing the device memory to attach.

- `memoryOffset` is the start offset of the region of `memory` which is
  to be bound to the image. The number of bytes returned in the
  `VkMemoryRequirements`::`size` member in `memory`, starting from
  `memoryOffset` bytes, will be bound to the specified image.

## <a href="#_description" class="anchor"></a>Description

`vkBindImageMemory` is equivalent to passing the same parameters through
[VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryInfo.html) to
[vkBindImageMemory2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindImageMemory2.html).

Valid Usage

- <a href="#VUID-vkBindImageMemory-image-07460"
  id="VUID-vkBindImageMemory-image-07460"></a>
  VUID-vkBindImageMemory-image-07460  
  `image` **must** not have been bound to a memory object

- <a href="#VUID-vkBindImageMemory-image-01045"
  id="VUID-vkBindImageMemory-image-01045"></a>
  VUID-vkBindImageMemory-image-01045  
  `image` **must** not have been created with any sparse memory binding
  flags

- <a href="#VUID-vkBindImageMemory-memoryOffset-01046"
  id="VUID-vkBindImageMemory-memoryOffset-01046"></a>
  VUID-vkBindImageMemory-memoryOffset-01046  
  `memoryOffset` **must** be less than the size of `memory`

- <a href="#VUID-vkBindImageMemory-image-01445"
  id="VUID-vkBindImageMemory-image-01445"></a>
  VUID-vkBindImageMemory-image-01445  
  If `image` requires a dedicated allocation (as reported by
  [vkGetImageMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements2.html) in
  [VkMemoryDedicatedRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedRequirements.html)::`requiresDedicatedAllocation`
  for `image`), `memory` **must** have been created with
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)::`image`
  equal to `image`

- <a href="#VUID-vkBindImageMemory-memory-02628"
  id="VUID-vkBindImageMemory-memory-02628"></a>
  VUID-vkBindImageMemory-memory-02628  
  If the
  [`dedicatedAllocationImageAliasing`](#features-dedicatedAllocationImageAliasing)
  feature is not enabled, and the `VkMemoryAllocateInfo` provided when
  `memory` was allocated included a
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)
  structure in its `pNext` chain, and
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)::`image`
  was not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), then `image` **must**
  equal
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)::`image`
  and `memoryOffset` **must** be zero

- <a href="#VUID-vkBindImageMemory-memory-02629"
  id="VUID-vkBindImageMemory-memory-02629"></a>
  VUID-vkBindImageMemory-memory-02629  
  If the
  [`dedicatedAllocationImageAliasing`](#features-dedicatedAllocationImageAliasing)
  feature is enabled, and the `VkMemoryAllocateInfo` provided when
  `memory` was allocated included a
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)
  structure in its `pNext` chain, and
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)::`image`
  was not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), then `memoryOffset`
  **must** be zero, and `image` **must** be either equal to
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)::`image`
  or an image that was created using the same parameters in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html), with the exception that
  `extent` and `arrayLayers` **may** differ subject to the following
  restrictions: every dimension in the `extent` parameter of the image
  being bound **must** be equal to or smaller than the original image
  for which the allocation was created; and the `arrayLayers` parameter
  of the image being bound **must** be equal to or smaller than the
  original image for which the allocation was created

- <a href="#VUID-vkBindImageMemory-None-01901"
  id="VUID-vkBindImageMemory-None-01901"></a>
  VUID-vkBindImageMemory-None-01901  
  If image was created with the `VK_IMAGE_CREATE_PROTECTED_BIT` bit set,
  the image **must** be bound to a memory object allocated with a memory
  type that reports `VK_MEMORY_PROPERTY_PROTECTED_BIT`

- <a href="#VUID-vkBindImageMemory-None-01902"
  id="VUID-vkBindImageMemory-None-01902"></a>
  VUID-vkBindImageMemory-None-01902  
  If image was created with the `VK_IMAGE_CREATE_PROTECTED_BIT` bit not
  set, the image **must** not be bound to a memory object created with a
  memory type that reports `VK_MEMORY_PROPERTY_PROTECTED_BIT`

- <a href="#VUID-vkBindImageMemory-image-01050"
  id="VUID-vkBindImageMemory-image-01050"></a>
  VUID-vkBindImageMemory-image-01050  
  If `image` was created with
  [VkDedicatedAllocationImageCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDedicatedAllocationImageCreateInfoNV.html)::`dedicatedAllocation`
  equal to `VK_TRUE`, `memory` **must** have been created with
  [VkDedicatedAllocationMemoryAllocateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDedicatedAllocationMemoryAllocateInfoNV.html)::`image`
  equal to an image handle created with identical creation parameters to
  `image` and `memoryOffset` **must** be zero

- <a href="#VUID-vkBindImageMemory-apiVersion-07921"
  id="VUID-vkBindImageMemory-apiVersion-07921"></a>
  VUID-vkBindImageMemory-apiVersion-07921  
  If the [VK_KHR_dedicated_allocation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dedicated_allocation.html)
  extension is not enabled,
  [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html)::`apiVersion`
  is less than Vulkan 1.1, and `image` was not created with
  [VkDedicatedAllocationImageCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDedicatedAllocationImageCreateInfoNV.html)::`dedicatedAllocation`
  equal to `VK_TRUE`, `memory` **must** not have been allocated
  dedicated for a specific buffer or image

- <a href="#VUID-vkBindImageMemory-memory-02728"
  id="VUID-vkBindImageMemory-memory-02728"></a>
  VUID-vkBindImageMemory-memory-02728  
  If the value of
  [VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfo.html)::`handleTypes`
  used to allocate `memory` is not `0`, it **must** include at least one
  of the handles set in
  [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html)::`handleTypes`
  when `image` was created

- <a href="#VUID-vkBindImageMemory-memory-02989"
  id="VUID-vkBindImageMemory-memory-02989"></a>
  VUID-vkBindImageMemory-memory-02989  
  If `memory` was created by a memory import operation, that is not
  [VkImportAndroidHardwareBufferInfoANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportAndroidHardwareBufferInfoANDROID.html)
  with a non-`NULL` `buffer` value, the external handle type of the
  imported memory **must** also have been set in
  [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html)::`handleTypes`
  when `image` was created

- <a href="#VUID-vkBindImageMemory-memory-02990"
  id="VUID-vkBindImageMemory-memory-02990"></a>
  VUID-vkBindImageMemory-memory-02990  
  If `memory` was created with the
  [VkImportAndroidHardwareBufferInfoANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportAndroidHardwareBufferInfoANDROID.html)
  memory import operation with a non-`NULL` `buffer` value,
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID`
  **must** also have been set in
  [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html)::`handleTypes`
  when `image` was created

- <a href="#VUID-vkBindImageMemory-descriptorBufferCaptureReplay-08113"
  id="VUID-vkBindImageMemory-descriptorBufferCaptureReplay-08113"></a>
  VUID-vkBindImageMemory-descriptorBufferCaptureReplay-08113  
  If the `image` was created with the
  `VK_IMAGE_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT` bit set,
  `memory` **must** have been allocated with the
  `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT` bit set

- <a href="#VUID-vkBindImageMemory-image-09202"
  id="VUID-vkBindImageMemory-image-09202"></a>
  VUID-vkBindImageMemory-image-09202  
  If the `image` was created with the
  `VK_IMAGE_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT` bit set,
  `memory` **must** have been allocated with the
  `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT` bit set

- <a href="#VUID-vkBindImageMemory-image-01608"
  id="VUID-vkBindImageMemory-image-01608"></a>
  VUID-vkBindImageMemory-image-01608  
  `image` **must** not have been created with the
  `VK_IMAGE_CREATE_DISJOINT_BIT` set

- <a href="#VUID-vkBindImageMemory-memory-01047"
  id="VUID-vkBindImageMemory-memory-01047"></a>
  VUID-vkBindImageMemory-memory-01047  
  `memory` **must** have been allocated using one of the memory types
  allowed in the `memoryTypeBits` member of the `VkMemoryRequirements`
  structure returned from a call to
  [vkGetImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements.html) with
  `image`

- <a href="#VUID-vkBindImageMemory-memoryOffset-01048"
  id="VUID-vkBindImageMemory-memoryOffset-01048"></a>
  VUID-vkBindImageMemory-memoryOffset-01048  
  `memoryOffset` **must** be an integer multiple of the `alignment`
  member of the `VkMemoryRequirements` structure returned from a call to
  [vkGetImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements.html) with
  `image`

- <a href="#VUID-vkBindImageMemory-size-01049"
  id="VUID-vkBindImageMemory-size-01049"></a>
  VUID-vkBindImageMemory-size-01049  
  The difference of the size of `memory` and `memoryOffset` **must** be
  greater than or equal to the `size` member of the
  [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html) structure returned
  from a call to
  [vkGetImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements.html) with
  the same `image`

- <a href="#VUID-vkBindImageMemory-image-06392"
  id="VUID-vkBindImageMemory-image-06392"></a>
  VUID-vkBindImageMemory-image-06392  
  If `image` was created with
  [VkBufferCollectionImageCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionImageCreateInfoFUCHSIA.html)
  chained to [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`pNext`,
  `memory` **must** be allocated with a
  [VkImportMemoryBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryBufferCollectionFUCHSIA.html)
  chained to [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html)::`pNext`

Valid Usage (Implicit)

- <a href="#VUID-vkBindImageMemory-device-parameter"
  id="VUID-vkBindImageMemory-device-parameter"></a>
  VUID-vkBindImageMemory-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkBindImageMemory-image-parameter"
  id="VUID-vkBindImageMemory-image-parameter"></a>
  VUID-vkBindImageMemory-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-vkBindImageMemory-memory-parameter"
  id="VUID-vkBindImageMemory-memory-parameter"></a>
  VUID-vkBindImageMemory-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)
  handle

- <a href="#VUID-vkBindImageMemory-image-parent"
  id="VUID-vkBindImageMemory-image-parent"></a>
  VUID-vkBindImageMemory-image-parent  
  `image` **must** have been created, allocated, or retrieved from
  `device`

- <a href="#VUID-vkBindImageMemory-memory-parent"
  id="VUID-vkBindImageMemory-memory-parent"></a>
  VUID-vkBindImageMemory-memory-parent  
  `memory` **must** have been created, allocated, or retrieved from
  `device`

Host Synchronization

- Host access to `image` **must** be externally synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html), [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkBindImageMemory"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
