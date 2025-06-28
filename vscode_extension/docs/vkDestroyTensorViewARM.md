# vkDestroyTensorViewARM(3) Manual Page

## Name

vkDestroyTensorViewARM - Destroy a tensor view object



## [](#_c_specification)C Specification

To destroy a tensor view, call:

```c++
// Provided by VK_ARM_tensors
void vkDestroyTensorViewARM(
    VkDevice                                    device,
    VkTensorViewARM                             tensorView,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that destroys the tensor view.
- `tensorView` is the tensor view to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroyTensorViewARM-tensorView-09750)VUID-vkDestroyTensorViewARM-tensorView-09750  
  All submitted commands that refer to `tensorView` **must** have completed execution
- [](#VUID-vkDestroyTensorViewARM-tensorView-09751)VUID-vkDestroyTensorViewARM-tensorView-09751  
  If `VkAllocationCallbacks` were provided when `tensorView` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyTensorViewARM-tensorView-09752)VUID-vkDestroyTensorViewARM-tensorView-09752  
  If no `VkAllocationCallbacks` were provided when `tensorView` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkDestroyTensorViewARM-device-parameter)VUID-vkDestroyTensorViewARM-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyTensorViewARM-tensorView-parameter)VUID-vkDestroyTensorViewARM-tensorView-parameter  
  If `tensorView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `tensorView` **must** be a valid [VkTensorViewARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewARM.html) handle
- [](#VUID-vkDestroyTensorViewARM-pAllocator-parameter)VUID-vkDestroyTensorViewARM-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyTensorViewARM-tensorView-parent)VUID-vkDestroyTensorViewARM-tensorView-parent  
  If `tensorView` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `tensorView` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkTensorViewARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyTensorViewARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0