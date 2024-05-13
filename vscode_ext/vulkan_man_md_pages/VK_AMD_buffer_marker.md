# VK_AMD_buffer_marker(3) Manual Page

## Name

VK_AMD_buffer_marker - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

180

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_special_use" class="anchor"></a>Special Use

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-compatibility-specialuse"
  target="_blank" rel="noopener">Developer tools</a>

## <a href="#_contact" class="anchor"></a>Contact

- Daniel Rakos <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_buffer_marker%5D%20@drakos-amd%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_buffer_marker%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>drakos-amd</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2018-01-26

**IP Status**  
No known IP claims.

**Contributors**  
- Matthaeus G. Chajdas, AMD

- Jaakko Konttinen, AMD

- Daniel Rakos, AMD

## <a href="#_description" class="anchor"></a>Description

This extension adds a new operation to execute pipelined writes of small
marker values into a `VkBuffer` object.

The primary purpose of these markers is to facilitate the development of
debugging tools for tracking which pipelined command contributed to
device loss.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdWriteBufferMarkerAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteBufferMarkerAMD.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_AMD_BUFFER_MARKER_EXTENSION_NAME`

- `VK_AMD_BUFFER_MARKER_SPEC_VERSION`

## <a href="#_examples" class="anchor"></a>Examples

None.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-01-26 (Jaakko Konttinen)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_AMD_buffer_marker"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
