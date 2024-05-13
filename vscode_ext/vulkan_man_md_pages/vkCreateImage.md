# vkCreateImage(3) Manual Page

## Name

vkCreateImage - Create a new image object



## <a href="#_c_specification" class="anchor"></a>C Specification

To create images, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkCreateImage(
    VkDevice                                    device,
    const VkImageCreateInfo*                    pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkImage*                                    pImage);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the image.

- `pCreateInfo` is a pointer to a
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) structure containing
  parameters to be used to create the image.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pImage` is a pointer to a [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle in which the
  resulting image object is returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkCreateImage-flags-00939"
  id="VUID-vkCreateImage-flags-00939"></a>
  VUID-vkCreateImage-flags-00939  
  If the `flags` member of `pCreateInfo` includes
  `VK_IMAGE_CREATE_SPARSE_BINDING_BIT`, and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedSparseAddressSpace"
  target="_blank"
  rel="noopener"><code>extendedSparseAddressSpace</code></a> feature is
  not enabled, creating this `VkImage` **must** not cause the total
  required sparse memory for all currently valid sparse resources on the
  device to exceed `VkPhysicalDeviceLimits`::`sparseAddressSpaceSize`

- <a href="#VUID-vkCreateImage-flags-09385"
  id="VUID-vkCreateImage-flags-09385"></a>
  VUID-vkCreateImage-flags-09385  
  If the `flags` member of `pCreateInfo` includes
  `VK_IMAGE_CREATE_SPARSE_BINDING_BIT`, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedSparseAddressSpace"
  target="_blank"
  rel="noopener"><code>extendedSparseAddressSpace</code></a> feature is
  enabled, and the `usage` member of `pCreateInfo` contains bits not in
  `VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV`::`extendedSparseImageUsageFlags`,
  creating this `VkImage` **must** not cause the total required sparse
  memory for all currently valid sparse resources on the device,
  excluding `VkBuffer` created with `usage` member of `pCreateInfo`
  containing bits in
  `VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV`::`extendedSparseBufferUsageFlags`
  and `VkImage` created with `usage` member of `pCreateInfo` containing
  bits in
  `VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV`::`extendedSparseImageUsageFlags`,
  to exceed `VkPhysicalDeviceLimits`::`sparseAddressSpaceSize`

- <a href="#VUID-vkCreateImage-flags-09386"
  id="VUID-vkCreateImage-flags-09386"></a>
  VUID-vkCreateImage-flags-09386  
  If the `flags` member of `pCreateInfo` includes
  `VK_IMAGE_CREATE_SPARSE_BINDING_BIT` and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedSparseAddressSpace"
  target="_blank"
  rel="noopener"><code>extendedSparseAddressSpace</code></a> feature is
  enabled, creating this `VkImage` **must** not cause the total required
  sparse memory for all currently valid sparse resources on the device
  to exceed
  `VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV`::`extendedSparseAddressSpaceSize`

<!-- -->

- <a href="#VUID-vkCreateImage-pNext-06389"
  id="VUID-vkCreateImage-pNext-06389"></a>
  VUID-vkCreateImage-pNext-06389  
  If a
  [VkBufferCollectionImageCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionImageCreateInfoFUCHSIA.html)
  has been chained to `pNext`, `pCreateInfo` **must** match the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#sysmem-chosen-create-infos"
  target="_blank" rel="noopener">Sysmem chosen
  <code>VkImageCreateInfo</code></a> excepting members
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`extent` and
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`usage` in the match
  criteria

Valid Usage (Implicit)

- <a href="#VUID-vkCreateImage-device-parameter"
  id="VUID-vkCreateImage-device-parameter"></a>
  VUID-vkCreateImage-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateImage-pCreateInfo-parameter"
  id="VUID-vkCreateImage-pCreateInfo-parameter"></a>
  VUID-vkCreateImage-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) structure

- <a href="#VUID-vkCreateImage-pAllocator-parameter"
  id="VUID-vkCreateImage-pAllocator-parameter"></a>
  VUID-vkCreateImage-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateImage-pImage-parameter"
  id="VUID-vkCreateImage-pImage-parameter"></a>
  VUID-vkCreateImage-pImage-parameter  
  `pImage` **must** be a valid pointer to a [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html)
  handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_COMPRESSION_EXHAUSTED_EXT`

- `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html),
[VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateImage"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
