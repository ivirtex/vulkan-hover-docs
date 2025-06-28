# VK\_MVK\_macos\_surface(3) Manual Page

## Name

VK\_MVK\_macos\_surface - instance extension



## [](#_registered_extension_number)Registered Extension Number

124

## [](#_revision)Revision

3

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html)

## [](#_deprecation_state)Deprecation State

- *Deprecated* by [VK\_EXT\_metal\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_metal_surface.html) extension

## [](#_contact)Contact

- Bill Hollings [\[GitHub\]billhollings](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_MVK_macos_surface%5D%20%40billhollings%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_MVK_macos_surface%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2020-07-31

**IP Status**

No known IP claims.

**Contributors**

- Bill Hollings, The Brenwill Workshop Ltd.

## [](#_description)Description

The `VK_MVK_macos_surface` extension is an instance extension. It provides a mechanism to create a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) object (defined by the `VK_KHR_surface` extension) based on an `NSView`, the native surface type of macOS, which is underpinned by a [CAMetalLayer](https://registry.khronos.org/vulkan/specs/latest/man/html/CAMetalLayer.html), to support rendering to the surface using Apple’s Metal framework.

## [](#_deprecation_by_vk_ext_metal_surface)Deprecation by `VK_EXT_metal_surface`

The `VK_MVK_macos_surface` extension is considered deprecated and has been superseded by the `VK_EXT_metal_surface` extension.

## [](#_new_commands)New Commands

- [vkCreateMacOSSurfaceMVK](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateMacOSSurfaceMVK.html)

## [](#_new_structures)New Structures

- [VkMacOSSurfaceCreateInfoMVK](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMacOSSurfaceCreateInfoMVK.html)

## [](#_new_bitmasks)New Bitmasks

- [VkMacOSSurfaceCreateFlagsMVK](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMacOSSurfaceCreateFlagsMVK.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_MVK_MACOS_SURFACE_EXTENSION_NAME`
- `VK_MVK_MACOS_SURFACE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_MACOS_SURFACE_CREATE_INFO_MVK`

## [](#_version_history)Version History

- Revision 1, 2017-02-15 (Bill Hollings)
  
  - Initial draft.
- Revision 2, 2017-02-24 (Bill Hollings)
  
  - Minor syntax fix to emphasize firm requirement for `NSView` to be backed by a `CAMetalLayer`.
- Revision 3, 2020-07-31 (Bill Hollings)
  
  - Update documentation on requirements for `NSView`.
  - Mark as deprecated by `VK_EXT_metal_surface`.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_MVK_macos_surface)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0