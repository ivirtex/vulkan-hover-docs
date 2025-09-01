# VkDisplayPlaneInfo2KHR(3) Manual Page

## Name

VkDisplayPlaneInfo2KHR - Structure defining the intended configuration of a display plane



## [](#_c_specification)C Specification

The `VkDisplayPlaneInfo2KHR` structure is defined as:

```c++
// Provided by VK_KHR_get_display_properties2
typedef struct VkDisplayPlaneInfo2KHR {
    VkStructureType     sType;
    const void*         pNext;
    VkDisplayModeKHR    mode;
    uint32_t            planeIndex;
} VkDisplayPlaneInfo2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `mode` is the display mode the application intends to program when using the specified plane.

## [](#_description)Description

Note

This parameter also implicitly specifies a display.

- `planeIndex` is the plane which the application intends to use with the display.

The members of `VkDisplayPlaneInfo2KHR` correspond to the arguments to [vkGetDisplayPlaneCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDisplayPlaneCapabilitiesKHR.html), with `sType` and `pNext` added for extensibility.

Valid Usage (Implicit)

- [](#VUID-VkDisplayPlaneInfo2KHR-sType-sType)VUID-VkDisplayPlaneInfo2KHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DISPLAY_PLANE_INFO_2_KHR`
- [](#VUID-VkDisplayPlaneInfo2KHR-pNext-pNext)VUID-VkDisplayPlaneInfo2KHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkDisplayPlaneInfo2KHR-mode-parameter)VUID-VkDisplayPlaneInfo2KHR-mode-parameter  
  `mode` **must** be a valid [VkDisplayModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeKHR.html) handle

Host Synchronization

- Host access to `mode` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_KHR\_get\_display\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_display_properties2.html), [VkDisplayModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetDisplayPlaneCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDisplayPlaneCapabilities2KHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDisplayPlaneInfo2KHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0