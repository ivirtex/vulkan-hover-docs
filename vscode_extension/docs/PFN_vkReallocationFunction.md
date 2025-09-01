# PFN\_vkReallocationFunction(3) Manual Page

## Name

PFN\_vkReallocationFunction - Application-defined memory reallocation function



## [](#_c_specification)C Specification

The type of `pfnReallocation` is:

```c++
// Provided by VK_VERSION_1_0
typedef void* (VKAPI_PTR *PFN_vkReallocationFunction)(
    void*                                       pUserData,
    void*                                       pOriginal,
    size_t                                      size,
    size_t                                      alignment,
    VkSystemAllocationScope                     allocationScope);
```

## [](#_parameters)Parameters

- `pUserData` is the value specified for [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html)::`pUserData` in the allocator specified by the application.
- `pOriginal` **must** be either `NULL` or a pointer previously returned by `pfnReallocation` or `pfnAllocation` of a compatible allocator.
- `size` is the size in bytes of the requested allocation.
- `alignment` is the requested alignment of the allocation in bytes and **must** be a power of two.
- `allocationScope` is a [VkSystemAllocationScope](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSystemAllocationScope.html) value specifying the allocation scope of the lifetime of the allocation, as described [here](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-host-allocation-scope).

## [](#_description)Description

If the reallocation was successful, `pfnReallocation` **must** return an allocation with enough space for `size` bytes, and the contents of the original allocation from bytes zero to min(original size, new size) - 1 **must** be preserved in the returned allocation. If `size` is larger than the old size, the contents of the additional space are undefined. If satisfying these requirements involves creating a new allocation, then the old allocation **should** be freed.

If `pOriginal` is `NULL`, then `pfnReallocation` **must** behave equivalently to a call to [PFN\_vkAllocationFunction](https://registry.khronos.org/vulkan/specs/latest/man/html/PFN_vkAllocationFunction.html) with the same parameter values (without `pOriginal`).

If `size` is zero, then `pfnReallocation` **must** behave equivalently to a call to [PFN\_vkFreeFunction](https://registry.khronos.org/vulkan/specs/latest/man/html/PFN_vkFreeFunction.html) with the same `pUserData` parameter value, and `pMemory` equal to `pOriginal`.

If `pOriginal` is non-`NULL`, the implementation **must** ensure that `alignment` is equal to the `alignment` used to originally allocate `pOriginal`.

If this function fails and `pOriginal` is non-`NULL` the application **must** not free the old allocation.

`pfnReallocation` **must** follow the same [rules for return values as `PFN_vkAllocationFunction`](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAllocationFunction_return_rules.html).

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkSystemAllocationScope](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSystemAllocationScope.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#PFN_vkReallocationFunction).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0