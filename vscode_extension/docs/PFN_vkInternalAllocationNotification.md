# PFN\_vkInternalAllocationNotification(3) Manual Page

## Name

PFN\_vkInternalAllocationNotification - Application-defined memory allocation notification function



## [](#_c_specification)C Specification

The type of `pfnInternalAllocation` is:

```c++
// Provided by VK_VERSION_1_0
typedef void (VKAPI_PTR *PFN_vkInternalAllocationNotification)(
    void*                                       pUserData,
    size_t                                      size,
    VkInternalAllocationType                    allocationType,
    VkSystemAllocationScope                     allocationScope);
```

## [](#_parameters)Parameters

- `pUserData` is the value specified for [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html)::`pUserData` in the allocator specified by the application.
- `size` is the requested size of an allocation.
- `allocationType` is a [VkInternalAllocationType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInternalAllocationType.html) value specifying the requested type of an allocation.
- `allocationScope` is a [VkSystemAllocationScope](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSystemAllocationScope.html) value specifying the allocation scope of the lifetime of the allocation, as described [here](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-host-allocation-scope).

## [](#_description)Description

This is a purely informational callback.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkInternalAllocationType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInternalAllocationType.html), [VkSystemAllocationScope](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSystemAllocationScope.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#PFN_vkInternalAllocationNotification).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0