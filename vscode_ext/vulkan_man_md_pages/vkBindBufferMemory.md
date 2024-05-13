# vkBindBufferMemory(3) Manual Page

## Name

vkBindBufferMemory - Bind device memory to a buffer object



## <a href="#_c_specification" class="anchor"></a>C Specification

To attach memory to a buffer object, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkBindBufferMemory(
    VkDevice                                    device,
    VkBuffer                                    buffer,
    VkDeviceMemory                              memory,
    VkDeviceSize                                memoryOffset);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the buffer and memory.

- `buffer` is the buffer to be attached to memory.

- `memory` is a [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object describing
  the device memory to attach.

- `memoryOffset` is the start offset of the region of `memory` which is
  to be bound to the buffer. The number of bytes returned in the
  `VkMemoryRequirements`::`size` member in `memory`, starting from
  `memoryOffset` bytes, will be bound to the specified buffer.

## <a href="#_description" class="anchor"></a>Description

`vkBindBufferMemory` is equivalent to passing the same parameters
through [VkBindBufferMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindBufferMemoryInfo.html) to
[vkBindBufferMemory2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindBufferMemory2.html).

Valid Usage

- <a href="#VUID-vkBindBufferMemory-buffer-07459"
  id="VUID-vkBindBufferMemory-buffer-07459"></a>
  VUID-vkBindBufferMemory-buffer-07459  
  `buffer` **must** not have been bound to a memory object

- <a href="#VUID-vkBindBufferMemory-buffer-01030"
  id="VUID-vkBindBufferMemory-buffer-01030"></a>
  VUID-vkBindBufferMemory-buffer-01030  
  `buffer` **must** not have been created with any sparse memory binding
  flags

- <a href="#VUID-vkBindBufferMemory-memoryOffset-01031"
  id="VUID-vkBindBufferMemory-memoryOffset-01031"></a>
  VUID-vkBindBufferMemory-memoryOffset-01031  
  `memoryOffset` **must** be less than the size of `memory`

- <a href="#VUID-vkBindBufferMemory-memory-01035"
  id="VUID-vkBindBufferMemory-memory-01035"></a>
  VUID-vkBindBufferMemory-memory-01035  
  `memory` **must** have been allocated using one of the memory types
  allowed in the `memoryTypeBits` member of the `VkMemoryRequirements`
  structure returned from a call to `vkGetBufferMemoryRequirements` with
  `buffer`

- <a href="#VUID-vkBindBufferMemory-memoryOffset-01036"
  id="VUID-vkBindBufferMemory-memoryOffset-01036"></a>
  VUID-vkBindBufferMemory-memoryOffset-01036  
  `memoryOffset` **must** be an integer multiple of the `alignment`
  member of the `VkMemoryRequirements` structure returned from a call to
  `vkGetBufferMemoryRequirements` with `buffer`

- <a href="#VUID-vkBindBufferMemory-size-01037"
  id="VUID-vkBindBufferMemory-size-01037"></a>
  VUID-vkBindBufferMemory-size-01037  
  The `size` member of the `VkMemoryRequirements` structure returned
  from a call to `vkGetBufferMemoryRequirements` with `buffer` **must**
  be less than or equal to the size of `memory` minus `memoryOffset`

- <a href="#VUID-vkBindBufferMemory-buffer-01444"
  id="VUID-vkBindBufferMemory-buffer-01444"></a>
  VUID-vkBindBufferMemory-buffer-01444  
  If `buffer` requires a dedicated allocation (as reported by
  [vkGetBufferMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferMemoryRequirements2.html)
  in
  [VkMemoryDedicatedRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedRequirements.html)::`requiresDedicatedAllocation`
  for `buffer`), `memory` **must** have been allocated with
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)::`buffer`
  equal to `buffer`

- <a href="#VUID-vkBindBufferMemory-memory-01508"
  id="VUID-vkBindBufferMemory-memory-01508"></a>
  VUID-vkBindBufferMemory-memory-01508  
  If the `VkMemoryAllocateInfo` provided when `memory` was allocated
  included a
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)
  structure in its `pNext` chain, and
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)::`buffer`
  was not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), then `buffer` **must**
  equal
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)::`buffer`,
  and `memoryOffset` **must** be zero

- <a href="#VUID-vkBindBufferMemory-None-01898"
  id="VUID-vkBindBufferMemory-None-01898"></a>
  VUID-vkBindBufferMemory-None-01898  
  If `buffer` was created with the `VK_BUFFER_CREATE_PROTECTED_BIT` bit
  set, the buffer **must** be bound to a memory object allocated with a
  memory type that reports `VK_MEMORY_PROPERTY_PROTECTED_BIT`

- <a href="#VUID-vkBindBufferMemory-None-01899"
  id="VUID-vkBindBufferMemory-None-01899"></a>
  VUID-vkBindBufferMemory-None-01899  
  If `buffer` was created with the `VK_BUFFER_CREATE_PROTECTED_BIT` bit
  not set, the buffer **must** not be bound to a memory object allocated
  with a memory type that reports `VK_MEMORY_PROPERTY_PROTECTED_BIT`

- <a href="#VUID-vkBindBufferMemory-buffer-01038"
  id="VUID-vkBindBufferMemory-buffer-01038"></a>
  VUID-vkBindBufferMemory-buffer-01038  
  If `buffer` was created with
  [VkDedicatedAllocationBufferCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDedicatedAllocationBufferCreateInfoNV.html)::`dedicatedAllocation`
  equal to `VK_TRUE`, `memory` **must** have been allocated with
  [VkDedicatedAllocationMemoryAllocateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDedicatedAllocationMemoryAllocateInfoNV.html)::`buffer`
  equal to a buffer handle created with identical creation parameters to
  `buffer` and `memoryOffset` **must** be zero

- <a href="#VUID-vkBindBufferMemory-apiVersion-07920"
  id="VUID-vkBindBufferMemory-apiVersion-07920"></a>
  VUID-vkBindBufferMemory-apiVersion-07920  
  If the [VK_KHR_dedicated_allocation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dedicated_allocation.html)
  extension is not enabled,
  [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html)::`apiVersion`
  is less than Vulkan 1.1, and `buffer` was not created with
  [VkDedicatedAllocationBufferCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDedicatedAllocationBufferCreateInfoNV.html)::`dedicatedAllocation`
  equal to `VK_TRUE`, `memory` **must** not have been allocated
  dedicated for a specific buffer or image

- <a href="#VUID-vkBindBufferMemory-memory-02726"
  id="VUID-vkBindBufferMemory-memory-02726"></a>
  VUID-vkBindBufferMemory-memory-02726  
  If the value of
  [VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfo.html)::`handleTypes`
  used to allocate `memory` is not `0`, it **must** include at least one
  of the handles set in
  [VkExternalMemoryBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryBufferCreateInfo.html)::`handleTypes`
  when `buffer` was created

- <a href="#VUID-vkBindBufferMemory-memory-02985"
  id="VUID-vkBindBufferMemory-memory-02985"></a>
  VUID-vkBindBufferMemory-memory-02985  
  If `memory` was allocated by a memory import operation, that is not
  [VkImportAndroidHardwareBufferInfoANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportAndroidHardwareBufferInfoANDROID.html)
  with a non-`NULL` `buffer` value, the external handle type of the
  imported memory **must** also have been set in
  [VkExternalMemoryBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryBufferCreateInfo.html)::`handleTypes`
  when `buffer` was created

- <a href="#VUID-vkBindBufferMemory-memory-02986"
  id="VUID-vkBindBufferMemory-memory-02986"></a>
  VUID-vkBindBufferMemory-memory-02986  
  If `memory` was allocated with the
  [VkImportAndroidHardwareBufferInfoANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportAndroidHardwareBufferInfoANDROID.html)
  memory import operation with a non-`NULL` `buffer` value,
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID`
  **must** also have been set in
  [VkExternalMemoryBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryBufferCreateInfo.html)::`handleTypes`
  when `buffer` was created

- <a href="#VUID-vkBindBufferMemory-bufferDeviceAddress-03339"
  id="VUID-vkBindBufferMemory-bufferDeviceAddress-03339"></a>
  VUID-vkBindBufferMemory-bufferDeviceAddress-03339  
  If the
  [VkPhysicalDeviceBufferDeviceAddressFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceBufferDeviceAddressFeatures.html)::`bufferDeviceAddress`
  feature is enabled and `buffer` was created with the
  `VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT` bit set, `memory` **must**
  have been allocated with the `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT`
  bit set

- <a
  href="#VUID-vkBindBufferMemory-bufferDeviceAddressCaptureReplay-09200"
  id="VUID-vkBindBufferMemory-bufferDeviceAddressCaptureReplay-09200"></a>
  VUID-vkBindBufferMemory-bufferDeviceAddressCaptureReplay-09200  
  If the
  [VkPhysicalDeviceBufferDeviceAddressFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceBufferDeviceAddressFeatures.html)::`bufferDeviceAddressCaptureReplay`
  feature is enabled and `buffer` was created with the
  `VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT` bit set, `memory`
  **must** have been allocated with the
  `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT` bit set

- <a href="#VUID-vkBindBufferMemory-buffer-06408"
  id="VUID-vkBindBufferMemory-buffer-06408"></a>
  VUID-vkBindBufferMemory-buffer-06408  
  If `buffer` was created with
  [VkBufferCollectionBufferCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionBufferCreateInfoFUCHSIA.html)
  chained to [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html)::`pNext`,
  `memory` **must** be allocated with a
  [VkImportMemoryBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryBufferCollectionFUCHSIA.html)
  chained to [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html)::`pNext`

- <a href="#VUID-vkBindBufferMemory-descriptorBufferCaptureReplay-08112"
  id="VUID-vkBindBufferMemory-descriptorBufferCaptureReplay-08112"></a>
  VUID-vkBindBufferMemory-descriptorBufferCaptureReplay-08112  
  If the `buffer` was created with the
  `VK_BUFFER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT` bit set,
  `memory` **must** have been allocated with the
  `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT` bit set

- <a href="#VUID-vkBindBufferMemory-buffer-09201"
  id="VUID-vkBindBufferMemory-buffer-09201"></a>
  VUID-vkBindBufferMemory-buffer-09201  
  If the `buffer` was created with the
  `VK_BUFFER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT` bit set,
  `memory` **must** have been allocated with the
  `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT` bit set

Valid Usage (Implicit)

- <a href="#VUID-vkBindBufferMemory-device-parameter"
  id="VUID-vkBindBufferMemory-device-parameter"></a>
  VUID-vkBindBufferMemory-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkBindBufferMemory-buffer-parameter"
  id="VUID-vkBindBufferMemory-buffer-parameter"></a>
  VUID-vkBindBufferMemory-buffer-parameter  
  `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-vkBindBufferMemory-memory-parameter"
  id="VUID-vkBindBufferMemory-memory-parameter"></a>
  VUID-vkBindBufferMemory-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)
  handle

- <a href="#VUID-vkBindBufferMemory-buffer-parent"
  id="VUID-vkBindBufferMemory-buffer-parent"></a>
  VUID-vkBindBufferMemory-buffer-parent  
  `buffer` **must** have been created, allocated, or retrieved from
  `device`

- <a href="#VUID-vkBindBufferMemory-memory-parent"
  id="VUID-vkBindBufferMemory-memory-parent"></a>
  VUID-vkBindBufferMemory-memory-parent  
  `memory` **must** have been created, allocated, or retrieved from
  `device`

Host Synchronization

- Host access to `buffer` **must** be externally synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkBindBufferMemory"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
