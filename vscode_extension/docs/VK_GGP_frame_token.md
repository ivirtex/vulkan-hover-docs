# VK\_GGP\_frame\_token(3) Manual Page

## Name

VK\_GGP\_frame\_token - device extension



## [](#_registered_extension_number)Registered Extension Number

192

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html)  
and  
[VK\_GGP\_stream\_descriptor\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_GGP_stream_descriptor_surface.html)

## [](#_contact)Contact

- Jean-Francois Roy [\[GitHub\]jfroy](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_GGP_frame_token%5D%20%40jfroy%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_GGP_frame_token%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-01-28

**IP Status**

No known IP claims.

**Contributors**

- Jean-Francois Roy, Google
- Richard O’Grady, Google

## [](#_description)Description

This extension allows an application that uses the `VK_KHR_swapchain` extension in combination with a Google Games Platform surface provided by the `VK_GGP_stream_descriptor_surface` extension to associate a Google Games Platform frame token with a present operation.

## [](#_new_structures)New Structures

- Extending [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html):
  
  - [VkPresentFrameTokenGGP](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentFrameTokenGGP.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_GGP_FRAME_TOKEN_EXTENSION_NAME`
- `VK_GGP_FRAME_TOKEN_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PRESENT_FRAME_TOKEN_GGP`

## [](#_version_history)Version History

- Revision 1, 2018-11-26 (Jean-Francois Roy)
  
  - Initial revision.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_GGP_frame_token).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0