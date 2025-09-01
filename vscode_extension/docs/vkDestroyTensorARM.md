# vkDestroyTensorARM(3) Manual Page

## Name

vkDestroyTensorARM - Destroy a tensor object



## [](#_c_specification)C Specification

To destroy a tensor, call:

```c++
// Provided by VK_ARM_tensors
void vkDestroyTensorARM(
    VkDevice                                    device,
    VkTensorARM                                 tensor,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that destroys the tensor.
- `tensor` is the tensor to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroyTensorARM-tensor-09730)VUID-vkDestroyTensorARM-tensor-09730  
  All submitted commands that refer to `tensor`, either directly or via a `VkTensorViewARM`, **must** have completed execution
- [](#VUID-vkDestroyTensorARM-tensor-09731)VUID-vkDestroyTensorARM-tensor-09731  
  If `VkAllocationCallbacks` were provided when `tensor` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyTensorARM-tensor-09732)VUID-vkDestroyTensorARM-tensor-09732  
  If no `VkAllocationCallbacks` were provided when `tensor` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkDestroyTensorARM-device-parameter)VUID-vkDestroyTensorARM-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyTensorARM-tensor-parameter)VUID-vkDestroyTensorARM-tensor-parameter  
  If `tensor` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `tensor` **must** be a valid [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html) handle
- [](#VUID-vkDestroyTensorARM-pAllocator-parameter)VUID-vkDestroyTensorARM-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyTensorARM-tensor-parent)VUID-vkDestroyTensorARM-tensor-parent  
  If `tensor` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `tensor` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyTensorARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0