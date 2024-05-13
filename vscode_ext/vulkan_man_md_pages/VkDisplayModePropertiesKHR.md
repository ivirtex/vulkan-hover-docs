# VkDisplayModePropertiesKHR(3) Manual Page

## Name

VkDisplayModePropertiesKHR - Structure describing display mode
properties



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDisplayModePropertiesKHR` structure is defined as:

``` c
// Provided by VK_KHR_display
typedef struct VkDisplayModePropertiesKHR {
    VkDisplayModeKHR              displayMode;
    VkDisplayModeParametersKHR    parameters;
} VkDisplayModePropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `displayMode` is a handle to the display mode described in this
  structure. This handle will be valid for the lifetime of the Vulkan
  instance.

- `parameters` is a
  [VkDisplayModeParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeParametersKHR.html)
  structure describing the display parameters associated with
  `displayMode`.

## <a href="#_description" class="anchor"></a>Description

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_display.html),
[VkDisplayModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeKHR.html),
[VkDisplayModeParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeParametersKHR.html),
[VkDisplayModeProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeProperties2KHR.html),
[vkGetDisplayModePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDisplayModePropertiesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDisplayModePropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
