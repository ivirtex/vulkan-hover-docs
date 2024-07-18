# VK_KHR_get_display_properties2(3) Manual Page

## Name

VK_KHR_get_display_properties2 - instance extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

122

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_display.html)  

## <a href="#_contact" class="anchor"></a>Contact

- James Jones <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_get_display_properties2%5D%20@cubanismo%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_get_display_properties2%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>cubanismo</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-02-21

**IP Status**  
No known IP claims.

**Contributors**  
- Ian Elliott, Google

- James Jones, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension provides new queries for device display properties and
capabilities that can be easily extended by other extensions, without
introducing any further queries. This extension can be considered the
[`VK_KHR_display`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_display.html) equivalent of the
[`VK_KHR_get_physical_device_properties2`](VK_KHR_get_physical_device_properties2.html)
extension.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetDisplayModeProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDisplayModeProperties2KHR.html)

- [vkGetDisplayPlaneCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDisplayPlaneCapabilities2KHR.html)

- [vkGetPhysicalDeviceDisplayPlaneProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceDisplayPlaneProperties2KHR.html)

- [vkGetPhysicalDeviceDisplayProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceDisplayProperties2KHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkDisplayModeProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeProperties2KHR.html)

- [VkDisplayPlaneCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlaneCapabilities2KHR.html)

- [VkDisplayPlaneInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlaneInfo2KHR.html)

- [VkDisplayPlaneProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlaneProperties2KHR.html)

- [VkDisplayProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayProperties2KHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_GET_DISPLAY_PROPERTIES_2_EXTENSION_NAME`

- `VK_KHR_GET_DISPLAY_PROPERTIES_2_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_DISPLAY_MODE_PROPERTIES_2_KHR`

  - `VK_STRUCTURE_TYPE_DISPLAY_PLANE_CAPABILITIES_2_KHR`

  - `VK_STRUCTURE_TYPE_DISPLAY_PLANE_INFO_2_KHR`

  - `VK_STRUCTURE_TYPE_DISPLAY_PLANE_PROPERTIES_2_KHR`

  - `VK_STRUCTURE_TYPE_DISPLAY_PROPERTIES_2_KHR`

## <a href="#_issues" class="anchor"></a>Issues

1\) What should this extension be named?

**RESOLVED**: `VK_KHR_get_display_properties2`. Other alternatives:

- `VK_KHR_display2`

- One extension, combined with `VK_KHR_surface_capabilites2`.

2\) Should extensible input structs be added for these new functions:

**RESOLVED**:

- [vkGetPhysicalDeviceDisplayProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceDisplayProperties2KHR.html):
  No. The only current input is a
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html). Other inputs would not make
  sense.

- [vkGetPhysicalDeviceDisplayPlaneProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceDisplayPlaneProperties2KHR.html):
  No. The only current input is a
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html). Other inputs would not make
  sense.

- [vkGetDisplayModeProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDisplayModeProperties2KHR.html):
  No. The only current inputs are a
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) and a
  [VkDisplayModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeKHR.html). Other inputs would not make
  sense.

3\) Should additional display query functions be extended?

**RESOLVED**:

- [vkGetDisplayPlaneSupportedDisplaysKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDisplayPlaneSupportedDisplaysKHR.html):
  No. Extensions should instead extend
  [vkGetDisplayPlaneCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDisplayPlaneCapabilitiesKHR.html)().

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-02-21 (James Jones)

  - Initial draft.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_get_display_properties2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
