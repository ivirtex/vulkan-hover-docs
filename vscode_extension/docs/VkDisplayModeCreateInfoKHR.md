# VkDisplayModeCreateInfoKHR(3) Manual Page

## Name

VkDisplayModeCreateInfoKHR - Structure specifying parameters of a newly created display mode object



## [](#_c_specification)C Specification

The `VkDisplayModeCreateInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_display
typedef struct VkDisplayModeCreateInfoKHR {
    VkStructureType                sType;
    const void*                    pNext;
    VkDisplayModeCreateFlagsKHR    flags;
    VkDisplayModeParametersKHR     parameters;
} VkDisplayModeCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use, and **must** be zero.
- `parameters` is a [VkDisplayModeParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeParametersKHR.html) structure describing the display parameters to use in creating the new mode. If the parameters are not compatible with the specified display, the implementation **must** return `VK_ERROR_INITIALIZATION_FAILED`.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkDisplayModeCreateInfoKHR-sType-sType)VUID-VkDisplayModeCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DISPLAY_MODE_CREATE_INFO_KHR`
- [](#VUID-VkDisplayModeCreateInfoKHR-pNext-pNext)VUID-VkDisplayModeCreateInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkDisplayModeCreateInfoKHR-flags-zerobitmask)VUID-VkDisplayModeCreateInfoKHR-flags-zerobitmask  
  `flags` **must** be `0`
- [](#VUID-VkDisplayModeCreateInfoKHR-parameters-parameter)VUID-VkDisplayModeCreateInfoKHR-parameters-parameter  
  `parameters` **must** be a valid [VkDisplayModeParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeParametersKHR.html) structure

## [](#_see_also)See Also

[VK\_KHR\_display](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_display.html), [VkDisplayModeCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeCreateFlagsKHR.html), [VkDisplayModeParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeParametersKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateDisplayModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDisplayModeKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDisplayModeCreateInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0