# VK\_KHR\_get\_display\_properties2(3) Manual Page

## Name

VK\_KHR\_get\_display\_properties2 - instance extension



## [](#_registered_extension_number)Registered Extension Number

122

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_display](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_display.html)

## [](#_contact)Contact

- James Jones [\[GitHub\]cubanismo](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_get_display_properties2%5D%20%40cubanismo%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_get_display_properties2%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-02-21

**IP Status**

No known IP claims.

**Contributors**

- Ian Elliott, Google
- James Jones, NVIDIA

## [](#_description)Description

This extension provides new queries for device display properties and capabilities that can be easily extended by other extensions, without introducing any further queries. This extension can be considered the `VK_KHR_display` equivalent of the `VK_KHR_get_physical_device_properties2` extension.

## [](#_new_commands)New Commands

- [vkGetDisplayModeProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDisplayModeProperties2KHR.html)
- [vkGetDisplayPlaneCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDisplayPlaneCapabilities2KHR.html)
- [vkGetPhysicalDeviceDisplayPlaneProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceDisplayPlaneProperties2KHR.html)
- [vkGetPhysicalDeviceDisplayProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceDisplayProperties2KHR.html)

## [](#_new_structures)New Structures

- [VkDisplayModeProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeProperties2KHR.html)
- [VkDisplayPlaneCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPlaneCapabilities2KHR.html)
- [VkDisplayPlaneInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPlaneInfo2KHR.html)
- [VkDisplayPlaneProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPlaneProperties2KHR.html)
- [VkDisplayProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayProperties2KHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_GET_DISPLAY_PROPERTIES_2_EXTENSION_NAME`
- `VK_KHR_GET_DISPLAY_PROPERTIES_2_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DISPLAY_MODE_PROPERTIES_2_KHR`
  - `VK_STRUCTURE_TYPE_DISPLAY_PLANE_CAPABILITIES_2_KHR`
  - `VK_STRUCTURE_TYPE_DISPLAY_PLANE_INFO_2_KHR`
  - `VK_STRUCTURE_TYPE_DISPLAY_PLANE_PROPERTIES_2_KHR`
  - `VK_STRUCTURE_TYPE_DISPLAY_PROPERTIES_2_KHR`

## [](#_issues)Issues

1\) What should this extension be named?

**RESOLVED**: `VK_KHR_get_display_properties2`. Other alternatives:

- `VK_KHR_display2`
- One extension, combined with `VK_KHR_surface_capabilites2`.

2\) Should extensible input structs be added for these new functions:

**RESOLVED**:

- [vkGetPhysicalDeviceDisplayProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceDisplayProperties2KHR.html): No. The only current input is a [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html). Other inputs would not make sense.
- [vkGetPhysicalDeviceDisplayPlaneProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceDisplayPlaneProperties2KHR.html): No. The only current input is a [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html). Other inputs would not make sense.
- [vkGetDisplayModeProperties2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDisplayModeProperties2KHR.html): No. The only current inputs are a [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) and a [VkDisplayModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeKHR.html). Other inputs would not make sense.

3\) Should additional display query functions be extended?

**RESOLVED**:

- [vkGetDisplayPlaneSupportedDisplaysKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDisplayPlaneSupportedDisplaysKHR.html): No. Extensions should instead extend [vkGetDisplayPlaneCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDisplayPlaneCapabilitiesKHR.html)().

## [](#_version_history)Version History

- Revision 1, 2017-02-21 (James Jones)
  
  - Initial draft.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_get_display_properties2).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0