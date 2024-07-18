# VkInternalAllocationType(3) Manual Page

## Name

VkInternalAllocationType - Allocation type



## <a href="#_c_specification" class="anchor"></a>C Specification

The `allocationType` parameter to the `pfnInternalAllocation` and
`pfnInternalFree` functions **may** be one of the following values:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkInternalAllocationType {
    VK_INTERNAL_ALLOCATION_TYPE_EXECUTABLE = 0,
} VkInternalAllocationType;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_INTERNAL_ALLOCATION_TYPE_EXECUTABLE` specifies that the allocation
  is intended for execution by the host.

## <a href="#_see_also" class="anchor"></a>See Also

[PFN_vkInternalAllocationNotification](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkInternalAllocationNotification.html),
[PFN_vkInternalFreeNotification](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkInternalFreeNotification.html),
[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkInternalAllocationType"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
