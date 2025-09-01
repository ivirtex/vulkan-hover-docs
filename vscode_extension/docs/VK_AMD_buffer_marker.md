# VK\_AMD\_buffer\_marker(3) Manual Page

## Name

VK\_AMD\_buffer\_marker - device extension



## [](#_registered_extension_number)Registered Extension Number

180

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_api_interactions)API Interactions

- Interacts with VK\_VERSION\_1\_3
- Interacts with VK\_KHR\_synchronization2

## [](#_special_use)Special Use

- [Developer tools](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Daniel Rakos [\[GitHub\]drakos-amd](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_buffer_marker%5D%20%40drakos-amd%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_buffer_marker%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-01-26

**IP Status**

No known IP claims.

**Contributors**

- Matthaeus G. Chajdas, AMD
- Jaakko Konttinen, AMD
- Daniel Rakos, AMD

## [](#_description)Description

This extension adds a new operation to execute pipelined writes of small marker values into a `VkBuffer` object.

The primary purpose of these markers is to facilitate the development of debugging tools for tracking which pipelined command contributed to device loss.

## [](#_new_commands)New Commands

- [vkCmdWriteBufferMarkerAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteBufferMarkerAMD.html)

If [Vulkan Version 1.3](#versions-1.3) or [VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html) is supported:

- [vkCmdWriteBufferMarker2AMD](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteBufferMarker2AMD.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_AMD_BUFFER_MARKER_EXTENSION_NAME`
- `VK_AMD_BUFFER_MARKER_SPEC_VERSION`

## [](#_examples)Examples

None.

## [](#_version_history)Version History

- Revision 1, 2018-01-26 (Jaakko Konttinen)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_AMD_buffer_marker).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0