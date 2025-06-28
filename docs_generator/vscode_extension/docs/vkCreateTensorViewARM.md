# vkCreateTensorViewARM(3) Manual Page

## Name

vkCreateTensorViewARM - Create an tensor view from an existing tensor



## [](#_c_specification)C Specification

To create a tensor view, call:

```c++
// Provided by VK_ARM_tensors
VkResult vkCreateTensorViewARM(
    VkDevice                                    device,
    const VkTensorViewCreateInfoARM*            pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkTensorViewARM*                            pView);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the tensor view.
- `pCreateInfo` is a pointer to an instance of the `VkTensorViewCreateInfoARM` structure containing parameters to be used to create the tensor view.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pView` is a pointer to a [VkTensorViewARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewARM.html) handle in which the resulting tensor view object is returned.

## [](#_description)Description

Some of the tensor creation parameters are inherited by the view. In particular, other than format, the tensor view creation inherits all other parameters from the tensor.

The remaining parameters are contained in `pCreateInfo`.

Valid Usage (Implicit)

- [](#VUID-vkCreateTensorViewARM-device-parameter)VUID-vkCreateTensorViewARM-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateTensorViewARM-pCreateInfo-parameter)VUID-vkCreateTensorViewARM-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkTensorViewCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewCreateInfoARM.html) structure
- [](#VUID-vkCreateTensorViewARM-pAllocator-parameter)VUID-vkCreateTensorViewARM-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateTensorViewARM-pView-parameter)VUID-vkCreateTensorViewARM-pView-parameter  
  `pView` **must** be a valid pointer to a [VkTensorViewARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewARM.html) handle
- [](#VUID-vkCreateTensorViewARM-device-queuecount)VUID-vkCreateTensorViewARM-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkTensorViewARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewARM.html), [VkTensorViewCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewCreateInfoARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateTensorViewARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0