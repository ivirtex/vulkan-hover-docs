# vkCreateBuffer(3) Manual Page

## Name

vkCreateBuffer - Create a new buffer object



## <a href="#_c_specification" class="anchor"></a>C Specification

To create buffers, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkCreateBuffer(
    VkDevice                                    device,
    const VkBufferCreateInfo*                   pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkBuffer*                                   pBuffer);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the buffer object.

- `pCreateInfo` is a pointer to a
  [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html) structure containing
  parameters affecting creation of the buffer.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pBuffer` is a pointer to a [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle in which
  the resulting buffer object is returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkCreateBuffer-flags-00911"
  id="VUID-vkCreateBuffer-flags-00911"></a>
  VUID-vkCreateBuffer-flags-00911  
  If the `flags` member of `pCreateInfo` includes
  `VK_BUFFER_CREATE_SPARSE_BINDING_BIT`, and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedSparseAddressSpace"
  target="_blank"
  rel="noopener"><code>extendedSparseAddressSpace</code></a> feature is
  not enabled, creating this `VkBuffer` **must** not cause the total
  required sparse memory for all currently valid sparse resources on the
  device to exceed `VkPhysicalDeviceLimits`::`sparseAddressSpaceSize`

- <a href="#VUID-vkCreateBuffer-flags-09383"
  id="VUID-vkCreateBuffer-flags-09383"></a>
  VUID-vkCreateBuffer-flags-09383  
  If the `flags` member of `pCreateInfo` includes
  `VK_BUFFER_CREATE_SPARSE_BINDING_BIT`, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedSparseAddressSpace"
  target="_blank"
  rel="noopener"><code>extendedSparseAddressSpace</code></a> feature is
  enabled, and the `usage` member of `pCreateInfo` contains bits not in
  `VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV`::`extendedSparseBufferUsageFlags`,
  creating this `VkBuffer` **must** not cause the total required sparse
  memory for all currently valid sparse resources on the device,
  excluding `VkBuffer` created with `usage` member of `pCreateInfo`
  containing bits in
  `VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV`::`extendedSparseBufferUsageFlags`
  and `VkImage` created with `usage` member of `pCreateInfo` containing
  bits in
  `VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV`::`extendedSparseImageUsageFlags`,
  to exceed `VkPhysicalDeviceLimits`::`sparseAddressSpaceSize`

- <a href="#VUID-vkCreateBuffer-flags-09384"
  id="VUID-vkCreateBuffer-flags-09384"></a>
  VUID-vkCreateBuffer-flags-09384  
  If the `flags` member of `pCreateInfo` includes
  `VK_BUFFER_CREATE_SPARSE_BINDING_BIT` and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedSparseAddressSpace"
  target="_blank"
  rel="noopener"><code>extendedSparseAddressSpace</code></a> feature is
  enabled, creating this `VkBuffer` **must** not cause the total
  required sparse memory for all currently valid sparse resources on the
  device to exceed
  `VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV`::`extendedSparseAddressSpaceSize`

<!-- -->

- <a href="#VUID-vkCreateBuffer-pNext-06387"
  id="VUID-vkCreateBuffer-pNext-06387"></a>
  VUID-vkCreateBuffer-pNext-06387  
  If using the [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) for an import operation from a
  [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html) where a
  [VkBufferCollectionBufferCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionBufferCreateInfoFUCHSIA.html)
  has been chained to `pNext`, `pCreateInfo` **must** match the
  [VkBufferConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferConstraintsInfoFUCHSIA.html)::`createInfo`
  used when setting the constraints on the buffer collection with
  [vkSetBufferCollectionBufferConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetBufferCollectionBufferConstraintsFUCHSIA.html)

Valid Usage (Implicit)

- <a href="#VUID-vkCreateBuffer-device-parameter"
  id="VUID-vkCreateBuffer-device-parameter"></a>
  VUID-vkCreateBuffer-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateBuffer-pCreateInfo-parameter"
  id="VUID-vkCreateBuffer-pCreateInfo-parameter"></a>
  VUID-vkCreateBuffer-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html) structure

- <a href="#VUID-vkCreateBuffer-pAllocator-parameter"
  id="VUID-vkCreateBuffer-pAllocator-parameter"></a>
  VUID-vkCreateBuffer-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateBuffer-pBuffer-parameter"
  id="VUID-vkCreateBuffer-pBuffer-parameter"></a>
  VUID-vkCreateBuffer-pBuffer-parameter  
  `pBuffer` **must** be a valid pointer to a [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html)
  handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html),
[VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateBuffer"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
