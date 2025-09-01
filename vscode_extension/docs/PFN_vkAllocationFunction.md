# PFN\_vkAllocationFunction(3) Manual Page

## Name

PFN\_vkAllocationFunction - Application-defined memory allocation function



## [](#_c_specification)C Specification

The type of `pfnAllocation` is:

```c++
// Provided by VK_VERSION_1_0
typedef void* (VKAPI_PTR *PFN_vkAllocationFunction)(
    void*                                       pUserData,
    size_t                                      size,
    size_t                                      alignment,
    VkSystemAllocationScope                     allocationScope);
```

## [](#_parameters)Parameters

- `pUserData` is the value specified for [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html)::`pUserData` in the allocator specified by the application.
- `size` is the size in bytes of the requested allocation.
- `alignment` is the requested alignment of the allocation in bytes and **must** be a power of two.
- `allocationScope` is a [VkSystemAllocationScope](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSystemAllocationScope.html) value specifying the allocation scope of the lifetime of the allocation, as described [here](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-host-allocation-scope).

## [](#_description)Description

If `pfnAllocation` is unable to allocate the requested memory, it **must** return `NULL`. If the allocation was successful, it **must** return a valid pointer to memory allocation containing at least `size` bytes, and with the pointer value being a multiple of `alignment`.

Note

Correct Vulkan operation **cannot** be assumed if the application does not follow these rules.

For example, `pfnAllocation` (or `pfnReallocation`) could cause termination of running Vulkan instance(s) on a failed allocation for debugging purposes, either directly or indirectly. In these circumstances, it **cannot** be assumed that any part of any affected [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) objects are going to operate correctly (even [vkDestroyInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyInstance.html)), and the application **must** ensure it cleans up properly via other means (e.g. process termination).

If `pfnAllocation` returns `NULL`, and if the implementation is unable to continue correct processing of the current command without the requested allocation, it **must** treat this as a runtime error, and generate `VK_ERROR_OUT_OF_HOST_MEMORY` at the appropriate time for the command in which the condition was detected, as described in [Return Codes](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fundamentals-errorcodes).

If the implementation is able to continue correct processing of the current command without the requested allocation, then it **may** do so, and **must** not generate `VK_ERROR_OUT_OF_HOST_MEMORY` as a result of this failed allocation.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkSystemAllocationScope](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSystemAllocationScope.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#PFN_vkAllocationFunction).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0