# VK_GGP_frame_token(3) Manual Page

## Name

VK_GGP_frame_token - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

192

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html)  
and  
[VK_GGP_stream_descriptor_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_GGP_stream_descriptor_surface.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Jean-Francois Roy <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_GGP_frame_token%5D%20@jfroy%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_GGP_frame_token%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>jfroy</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-01-28

**IP Status**  
No known IP claims.

**Contributors**  
- Jean-Francois Roy, Google

- Richard O’Grady, Google

## <a href="#_description" class="anchor"></a>Description

This extension allows an application that uses the
[`VK_KHR_swapchain`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html) extension in combination
with a Google Games Platform surface provided by the
[`VK_GGP_stream_descriptor_surface`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_GGP_stream_descriptor_surface.html)
extension to associate a Google Games Platform frame token with a
present operation.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentInfoKHR.html):

  - [VkPresentFrameTokenGGP](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentFrameTokenGGP.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_GGP_FRAME_TOKEN_EXTENSION_NAME`

- `VK_GGP_FRAME_TOKEN_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PRESENT_FRAME_TOKEN_GGP`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-11-26 (Jean-Francois Roy)

  - Initial revision.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_GGP_frame_token"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
