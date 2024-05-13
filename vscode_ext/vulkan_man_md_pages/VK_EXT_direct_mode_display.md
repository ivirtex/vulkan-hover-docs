# VK_EXT_direct_mode_display(3) Manual Page

## Name

VK_EXT_direct_mode_display - instance extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

89

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_display.html)  

## <a href="#_contact" class="anchor"></a>Contact

- James Jones <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_direct_mode_display%5D%20@cubanismo%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_direct_mode_display%20extension*"
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

- Liam Middlebrook, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This is extension, along with related platform extensions, allows
applications to take exclusive control of displays associated with a
native windowing system. This is especially useful for virtual reality
applications that wish to hide HMDs (head mounted displays) from the
native platform’s display management system, desktop, and/or other
applications.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkReleaseDisplayEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkReleaseDisplayEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_DIRECT_MODE_DISPLAY_EXTENSION_NAME`

- `VK_EXT_DIRECT_MODE_DISPLAY_SPEC_VERSION`

## <a href="#_issues" class="anchor"></a>Issues

1\) Should this extension and its related platform-specific extensions
leverage [`VK_KHR_display`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_display.html), or provide separate
equivalent interfaces.

**RESOLVED**: Use [`VK_KHR_display`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_display.html) concepts and
objects. [`VK_KHR_display`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_display.html) can be used to
enumerate all displays on the system, including those attached to/in use
by a window system or native platform, but
[`VK_KHR_display_swapchain`](VK_KHR_display_swapchain.html) will fail to
create a swapchain on in-use displays. This extension and its
platform-specific children will allow applications to grab in-use
displays away from window systems and/or native platforms, allowing them
to be used with
[`VK_KHR_display_swapchain`](VK_KHR_display_swapchain.html).

2\) Are separate calls needed to acquire displays and enable direct
mode?

**RESOLVED**: No, these operations happen in one combined command.
Acquiring a display puts it into direct mode.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-12-13 (James Jones)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_direct_mode_display"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
