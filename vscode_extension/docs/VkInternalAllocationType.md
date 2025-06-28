# VkInternalAllocationType(3) Manual Page

## Name

VkInternalAllocationType - Allocation type



## [](#_c_specification)C Specification

The `allocationType` parameter to the `pfnInternalAllocation` and `pfnInternalFree` functions **may** be one of the following values:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkInternalAllocationType {
    VK_INTERNAL_ALLOCATION_TYPE_EXECUTABLE = 0,
} VkInternalAllocationType;
```

## [](#_description)Description

- `VK_INTERNAL_ALLOCATION_TYPE_EXECUTABLE` specifies that the allocation is intended for execution by the host.

## [](#_see_also)See Also

[PFN\_vkInternalAllocationNotification](https://registry.khronos.org/vulkan/specs/latest/man/html/PFN_vkInternalAllocationNotification.html), [PFN\_vkInternalFreeNotification](https://registry.khronos.org/vulkan/specs/latest/man/html/PFN_vkInternalFreeNotification.html), [VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkInternalAllocationType)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0