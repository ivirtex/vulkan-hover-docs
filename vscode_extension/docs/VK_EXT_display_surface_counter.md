# VK_EXT_display_surface_counter(3) Manual Page

## Name

VK_EXT_display_surface_counter - instance extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

91

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_display.html)  

## <a href="#_contact" class="anchor"></a>Contact

- James Jones <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_display_surface_counter%5D%20@cubanismo%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_display_surface_counter%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>cubanismo</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2016-12-13

**IP Status**  
No known IP claims.

**Contributors**  
- Pierre Boudier, NVIDIA

- James Jones, NVIDIA

- Damien Leone, NVIDIA

- Pierre-Loup Griffais, Valve

- Daniel Vetter, Intel

## <a href="#_description" class="anchor"></a>Description

This extension defines a vertical blanking period counter associated
with display surfaces. It provides a mechanism to query support for such
a counter from a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) object.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetPhysicalDeviceSurfaceCapabilities2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilities2EXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkSurfaceCapabilities2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilities2EXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkSurfaceCounterFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCounterFlagBitsEXT.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkSurfaceCounterFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCounterFlagsEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_DISPLAY_SURFACE_COUNTER_EXTENSION_NAME`

- `VK_EXT_DISPLAY_SURFACE_COUNTER_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES2_EXT`

  - `VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES_2_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-12-13 (James Jones)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_display_surface_counter"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
