# vkCreateImage(3) Manual Page

## Name

vkCreateImage - Create a new image object



## [](#_c_specification)C Specification

To create images, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkCreateImage(
    VkDevice                                    device,
    const VkImageCreateInfo*                    pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkImage*                                    pImage);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the image.
- `pCreateInfo` is a pointer to a [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) structure containing parameters to be used to create the image.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pImage` is a pointer to a [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle in which the resulting image object is returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkCreateImage-device-09666)VUID-vkCreateImage-device-09666  
  `device` **must** support at least one queue family with one of the `VK_QUEUE_VIDEO_ENCODE_BIT_KHR`, `VK_QUEUE_VIDEO_DECODE_BIT_KHR`, `VK_QUEUE_OPTICAL_FLOW_BIT_NV`, `VK_QUEUE_SPARSE_BINDING_BIT`, `VK_QUEUE_TRANSFER_BIT`, `VK_QUEUE_COMPUTE_BIT`, or `VK_QUEUE_GRAPHICS_BIT` capabilities
- [](#VUID-vkCreateImage-flags-00939)VUID-vkCreateImage-flags-00939  
  If the `flags` member of `pCreateInfo` includes `VK_IMAGE_CREATE_SPARSE_BINDING_BIT`, and the [`extendedSparseAddressSpace`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-extendedSparseAddressSpace) feature is not enabled, creating this `VkImage` **must** not cause the total required sparse memory for all currently valid sparse resources on the device to exceed `VkPhysicalDeviceLimits`::`sparseAddressSpaceSize`
- [](#VUID-vkCreateImage-flags-09385)VUID-vkCreateImage-flags-09385  
  If the `flags` member of `pCreateInfo` includes `VK_IMAGE_CREATE_SPARSE_BINDING_BIT`, the [`extendedSparseAddressSpace`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-extendedSparseAddressSpace) feature is enabled, and the `usage` member of `pCreateInfo` contains bits not in `VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV`::`extendedSparseImageUsageFlags`, creating this `VkImage` **must** not cause the total required sparse memory for all currently valid sparse resources on the device, excluding `VkBuffer` created with `usage` member of `pCreateInfo` containing bits in `VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV`::`extendedSparseBufferUsageFlags` and `VkImage` created with `usage` member of `pCreateInfo` containing bits in `VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV`::`extendedSparseImageUsageFlags`, to exceed `VkPhysicalDeviceLimits`::`sparseAddressSpaceSize`
- [](#VUID-vkCreateImage-flags-09386)VUID-vkCreateImage-flags-09386  
  If the `flags` member of `pCreateInfo` includes `VK_IMAGE_CREATE_SPARSE_BINDING_BIT` and the [`extendedSparseAddressSpace`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-extendedSparseAddressSpace) feature is enabled, creating this `VkImage` **must** not cause the total required sparse memory for all currently valid sparse resources on the device to exceed `VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV`::`extendedSparseAddressSpaceSize`

<!--THE END-->

- [](#VUID-vkCreateImage-pNext-06389)VUID-vkCreateImage-pNext-06389  
  If a [VkBufferCollectionImageCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionImageCreateInfoFUCHSIA.html) has been chained to `pNext`, `pCreateInfo` **must** match the [Sysmem chosen `VkImageCreateInfo`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#sysmem-chosen-create-infos) excepting members [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`extent` and [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`usage` in the match criteria

Valid Usage (Implicit)

- [](#VUID-vkCreateImage-device-parameter)VUID-vkCreateImage-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateImage-pCreateInfo-parameter)VUID-vkCreateImage-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) structure
- [](#VUID-vkCreateImage-pAllocator-parameter)VUID-vkCreateImage-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateImage-pImage-parameter)VUID-vkCreateImage-pImage-parameter  
  `pImage` **must** be a valid pointer to a [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-vkCreateImage-device-queuecount)VUID-vkCreateImage-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_COMPRESSION_EXHAUSTED_EXT`
- `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateImage).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0