# VK\_FUCHSIA\_imagepipe\_surface(3) Manual Page

## Name

VK\_FUCHSIA\_imagepipe\_surface - instance extension



## [](#_registered_extension_number)Registered Extension Number

215

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html)

## [](#_contact)Contact

- Craig Stout [\[GitHub\]cdotstout](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_FUCHSIA_imagepipe_surface%5D%20%40cdotstout%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_FUCHSIA_imagepipe_surface%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-07-27

**IP Status**

No known IP claims.

**Contributors**

- Craig Stout, Google
- Ian Elliott, Google
- Jesse Hall, Google

## [](#_description)Description

The `VK_FUCHSIA_imagepipe_surface` extension is an instance extension. It provides a mechanism to create a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) object (defined by the `VK_KHR_surface` extension) that refers to a Fuchsia `imagePipeHandle`.

## [](#_new_commands)New Commands

- [vkCreateImagePipeSurfaceFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateImagePipeSurfaceFUCHSIA.html)

## [](#_new_structures)New Structures

- [VkImagePipeSurfaceCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImagePipeSurfaceCreateInfoFUCHSIA.html)

## [](#_new_bitmasks)New Bitmasks

- [VkImagePipeSurfaceCreateFlagsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImagePipeSurfaceCreateFlagsFUCHSIA.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_FUCHSIA_IMAGEPIPE_SURFACE_EXTENSION_NAME`
- `VK_FUCHSIA_IMAGEPIPE_SURFACE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_IMAGEPIPE_SURFACE_CREATE_INFO_FUCHSIA`

## [](#_version_history)Version History

- Revision 1, 2018-07-27 (Craig Stout)
  
  - Initial draft.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_FUCHSIA_imagepipe_surface).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0