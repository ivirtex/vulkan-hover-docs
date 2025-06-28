# VK\_GGP\_stream\_descriptor\_surface(3) Manual Page

## Name

VK\_GGP\_stream\_descriptor\_surface - instance extension



## [](#_registered_extension_number)Registered Extension Number

50

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html)

## [](#_contact)Contact

- Jean-Francois Roy [\[GitHub\]jfroy](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_GGP_stream_descriptor_surface%5D%20%40jfroy%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_GGP_stream_descriptor_surface%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-01-28

**IP Status**

No known IP claims.

**Contributors**

- Jean-Francois Roy, Google
- Brad Grantham, Google
- Connor Smith, Google
- Cort Stratton, Google
- Hai Nguyen, Google
- Ian Elliott, Google
- Jesse Hall, Google
- Jim Ray, Google
- Katherine Wu, Google
- Kaye Mason, Google
- Kuangye Guo, Google
- Mark Segal, Google
- Nicholas Vining, Google
- Paul Lalonde, Google
- Richard O’Grady, Google

## [](#_description)Description

The `VK_GGP_stream_descriptor_surface` extension is an instance extension. It provides a mechanism to create a [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html) object (defined by the `VK_KHR_surface` extension) that refers to a Google Games Platform `GgpStreamDescriptor`.

## [](#_new_commands)New Commands

- [vkCreateStreamDescriptorSurfaceGGP](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateStreamDescriptorSurfaceGGP.html)

## [](#_new_structures)New Structures

- [VkStreamDescriptorSurfaceCreateInfoGGP](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStreamDescriptorSurfaceCreateInfoGGP.html)

## [](#_new_bitmasks)New Bitmasks

- [VkStreamDescriptorSurfaceCreateFlagsGGP](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStreamDescriptorSurfaceCreateFlagsGGP.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_GGP_STREAM_DESCRIPTOR_SURFACE_EXTENSION_NAME`
- `VK_GGP_STREAM_DESCRIPTOR_SURFACE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_STREAM_DESCRIPTOR_SURFACE_CREATE_INFO_GGP`

## [](#_version_history)Version History

- Revision 1, 2018-11-26 (Jean-Francois Roy)
  
  - Initial revision.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_GGP_stream_descriptor_surface)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0