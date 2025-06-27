# VkDisplayModeCreateInfoKHR(3) Manual Page

## Name

VkDisplayModeCreateInfoKHR - Structure specifying parameters of a newly
created display mode object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDisplayModeCreateInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_display
typedef struct VkDisplayModeCreateInfoKHR {
    VkStructureType                sType;
    const void*                    pNext;
    VkDisplayModeCreateFlagsKHR    flags;
    VkDisplayModeParametersKHR     parameters;
} VkDisplayModeCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use, and **must** be zero.

- `parameters` is a
  [VkDisplayModeParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeParametersKHR.html)
  structure describing the display parameters to use in creating the new
  mode. If the parameters are not compatible with the specified display,
  the implementation **must** return `VK_ERROR_INITIALIZATION_FAILED`.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkDisplayModeCreateInfoKHR-sType-sType"
  id="VUID-VkDisplayModeCreateInfoKHR-sType-sType"></a>
  VUID-VkDisplayModeCreateInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DISPLAY_MODE_CREATE_INFO_KHR`

- <a href="#VUID-VkDisplayModeCreateInfoKHR-pNext-pNext"
  id="VUID-VkDisplayModeCreateInfoKHR-pNext-pNext"></a>
  VUID-VkDisplayModeCreateInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkDisplayModeCreateInfoKHR-flags-zerobitmask"
  id="VUID-VkDisplayModeCreateInfoKHR-flags-zerobitmask"></a>
  VUID-VkDisplayModeCreateInfoKHR-flags-zerobitmask  
  `flags` **must** be `0`

- <a href="#VUID-VkDisplayModeCreateInfoKHR-parameters-parameter"
  id="VUID-VkDisplayModeCreateInfoKHR-parameters-parameter"></a>
  VUID-VkDisplayModeCreateInfoKHR-parameters-parameter  
  `parameters` **must** be a valid
  [VkDisplayModeParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeParametersKHR.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_display.html),
[VkDisplayModeCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeCreateFlagsKHR.html),
[VkDisplayModeParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeParametersKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateDisplayModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDisplayModeKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDisplayModeCreateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
