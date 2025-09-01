# vkDestroyDeferredOperationKHR(3) Manual Page

## Name

vkDestroyDeferredOperationKHR - Destroy a deferred operation handle



## [](#_c_specification)C Specification

When a deferred operation is completed, the application **can** destroy the tracking object by calling:

```c++
// Provided by VK_KHR_deferred_host_operations
void vkDestroyDeferredOperationKHR(
    VkDevice                                    device,
    VkDeferredOperationKHR                      operation,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the device which owns `operation`.
- `operation` is the completed operation to be destroyed.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroyDeferredOperationKHR-operation-03434)VUID-vkDestroyDeferredOperationKHR-operation-03434  
  If `VkAllocationCallbacks` were provided when `operation` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyDeferredOperationKHR-operation-03435)VUID-vkDestroyDeferredOperationKHR-operation-03435  
  If no `VkAllocationCallbacks` were provided when `operation` was created, `pAllocator` **must** be `NULL`
- [](#VUID-vkDestroyDeferredOperationKHR-operation-03436)VUID-vkDestroyDeferredOperationKHR-operation-03436  
  `operation` **must** be completed

Valid Usage (Implicit)

- [](#VUID-vkDestroyDeferredOperationKHR-device-parameter)VUID-vkDestroyDeferredOperationKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyDeferredOperationKHR-operation-parameter)VUID-vkDestroyDeferredOperationKHR-operation-parameter  
  If `operation` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `operation` **must** be a valid [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html) handle
- [](#VUID-vkDestroyDeferredOperationKHR-pAllocator-parameter)VUID-vkDestroyDeferredOperationKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyDeferredOperationKHR-operation-parent)VUID-vkDestroyDeferredOperationKHR-operation-parent  
  If `operation` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `operation` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_KHR\_deferred\_host\_operations](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_deferred_host_operations.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyDeferredOperationKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0