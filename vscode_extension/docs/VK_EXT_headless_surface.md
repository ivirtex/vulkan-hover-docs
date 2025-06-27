# VK_EXT_headless_surface(3) Manual Page

## Name

VK_EXT_headless_surface - instance extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

257

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Lisa Wu <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_headless_surface%5D%20@chengtianww%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_headless_surface%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>chengtianww</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-03-21

**IP Status**  
No known IP claims.

**Contributors**  
- Ray Smith, Arm

## <a href="#_description" class="anchor"></a>Description

The `VK_EXT_headless_surface` extension is an instance extension. It
provides a mechanism to create [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) objects
independently of any window system or display device. The presentation
operation for a swapchain created from a headless surface is by default
a no-op, resulting in no externally-visible result.

Because there is no real presentation target, future extensions can
layer on top of the headless surface to introduce arbitrary or
customizable sets of restrictions or features. These could include
features like saving to a file or restrictions to emulate a particular
presentation target.

This functionality is expected to be useful for application and driver
development because it allows any platform to expose an arbitrary or
customizable set of restrictions and features of a presentation engine.
This makes it a useful portable test target for applications targeting a
wide range of presentation engines where the actual target presentation
engines might be scarce, unavailable or otherwise undesirable or
inconvenient to use for general Vulkan application development.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCreateHeadlessSurfaceEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateHeadlessSurfaceEXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkHeadlessSurfaceCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHeadlessSurfaceCreateInfoEXT.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkHeadlessSurfaceCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHeadlessSurfaceCreateFlagsEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_HEADLESS_SURFACE_EXTENSION_NAME`

- `VK_EXT_HEADLESS_SURFACE_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_HEADLESS_SURFACE_CREATE_INFO_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-03-21 (Ray Smith)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_headless_surface"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
