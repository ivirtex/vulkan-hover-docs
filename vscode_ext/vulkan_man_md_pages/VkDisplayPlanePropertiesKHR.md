# VkDisplayPlanePropertiesKHR(3) Manual Page

## Name

VkDisplayPlanePropertiesKHR - Structure describing display plane
properties



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDisplayPlanePropertiesKHR` structure is defined as:

``` c
// Provided by VK_KHR_display
typedef struct VkDisplayPlanePropertiesKHR {
    VkDisplayKHR    currentDisplay;
    uint32_t        currentStackIndex;
} VkDisplayPlanePropertiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `currentDisplay` is the handle of the display the plane is currently
  associated with. If the plane is not currently attached to any
  displays, this will be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html).

- `currentStackIndex` is the current z-order of the plane. This will be
  between 0 and the value returned by
  `vkGetPhysicalDeviceDisplayPlanePropertiesKHR` in `pPropertyCount`.

## <a href="#_description" class="anchor"></a>Description

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_display.html),
[VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html),
[VkDisplayPlaneProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlaneProperties2KHR.html),
[vkGetPhysicalDeviceDisplayPlanePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceDisplayPlanePropertiesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDisplayPlanePropertiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
