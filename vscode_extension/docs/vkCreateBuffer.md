# vkCreateBuffer(3) Manual Page

## Name

vkCreateBuffer - Create a new buffer object



## [](#_c_specification)C Specification

To create buffers, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkCreateBuffer(
    VkDevice                                    device,
    const VkBufferCreateInfo*                   pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkBuffer*                                   pBuffer);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the buffer object.
- `pCreateInfo` is a pointer to a [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html) structure containing parameters affecting creation of the buffer.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pBuffer` is a pointer to a [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle in which the resulting buffer object is returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkCreateBuffer-device-09664)VUID-vkCreateBuffer-device-09664  
  `device` **must** support at least one queue family with one of the `VK_QUEUE_VIDEO_ENCODE_BIT_KHR`, `VK_QUEUE_VIDEO_DECODE_BIT_KHR`, `VK_QUEUE_SPARSE_BINDING_BIT`, `VK_QUEUE_TRANSFER_BIT`, `VK_QUEUE_COMPUTE_BIT`, or `VK_QUEUE_GRAPHICS_BIT` capabilities
- [](#VUID-vkCreateBuffer-flags-00911)VUID-vkCreateBuffer-flags-00911  
  If the `flags` member of `pCreateInfo` includes `VK_BUFFER_CREATE_SPARSE_BINDING_BIT`, and the [`extendedSparseAddressSpace`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-extendedSparseAddressSpace) feature is not enabled, creating this `VkBuffer` **must** not cause the total required sparse memory for all currently valid sparse resources on the device to exceed `VkPhysicalDeviceLimits`::`sparseAddressSpaceSize`
- [](#VUID-vkCreateBuffer-flags-09383)VUID-vkCreateBuffer-flags-09383  
  If the `flags` member of `pCreateInfo` includes `VK_BUFFER_CREATE_SPARSE_BINDING_BIT`, the [`extendedSparseAddressSpace`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-extendedSparseAddressSpace) feature is enabled, and the `usage` member of `pCreateInfo` contains bits not in `VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV`::`extendedSparseBufferUsageFlags`, creating this `VkBuffer` **must** not cause the total required sparse memory for all currently valid sparse resources on the device, excluding `VkBuffer` created with `usage` member of `pCreateInfo` containing bits in `VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV`::`extendedSparseBufferUsageFlags` and `VkImage` created with `usage` member of `pCreateInfo` containing bits in `VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV`::`extendedSparseImageUsageFlags`, to exceed `VkPhysicalDeviceLimits`::`sparseAddressSpaceSize`
- [](#VUID-vkCreateBuffer-flags-09384)VUID-vkCreateBuffer-flags-09384  
  If the `flags` member of `pCreateInfo` includes `VK_BUFFER_CREATE_SPARSE_BINDING_BIT` and the [`extendedSparseAddressSpace`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-extendedSparseAddressSpace) feature is enabled, creating this `VkBuffer` **must** not cause the total required sparse memory for all currently valid sparse resources on the device to exceed `VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV`::`extendedSparseAddressSpaceSize`

<!--THE END-->

- [](#VUID-vkCreateBuffer-pNext-06387)VUID-vkCreateBuffer-pNext-06387  
  If using the [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) for an import operation from a [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionFUCHSIA.html) where a [VkBufferCollectionBufferCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionBufferCreateInfoFUCHSIA.html) has been chained to `pNext`, `pCreateInfo` **must** match the [VkBufferConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferConstraintsInfoFUCHSIA.html)::`createInfo` used when setting the constraints on the buffer collection with [vkSetBufferCollectionBufferConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetBufferCollectionBufferConstraintsFUCHSIA.html)

Valid Usage (Implicit)

- [](#VUID-vkCreateBuffer-device-parameter)VUID-vkCreateBuffer-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateBuffer-pCreateInfo-parameter)VUID-vkCreateBuffer-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html) structure
- [](#VUID-vkCreateBuffer-pAllocator-parameter)VUID-vkCreateBuffer-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateBuffer-pBuffer-parameter)VUID-vkCreateBuffer-pBuffer-parameter  
  `pBuffer` **must** be a valid pointer to a [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-vkCreateBuffer-device-queuecount)VUID-vkCreateBuffer-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateBuffer).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0