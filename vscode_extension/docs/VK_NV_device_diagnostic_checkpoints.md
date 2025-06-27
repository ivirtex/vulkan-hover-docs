# VK_NV_device_diagnostic_checkpoints(3) Manual Page

## Name

VK_NV_device_diagnostic_checkpoints - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

207

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Nuno Subtil <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_device_diagnostic_checkpoints%5D%20@nsubtil%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_device_diagnostic_checkpoints%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>nsubtil</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2018-07-16

**Contributors**  
- Oleg Kuznetsov, NVIDIA

- Alex Dunn, NVIDIA

- Jeff Bolz, NVIDIA

- Eric Werness, NVIDIA

- Daniel Koch, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension allows applications to insert markers in the command
stream and associate them with custom data.

If a device lost error occurs, the application **may** then query the
implementation for the last markers to cross specific
implementation-defined pipeline stages, in order to narrow down which
commands were executing at the time and might have caused the failure.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdSetCheckpointNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCheckpointNV.html)

- [vkGetQueueCheckpointDataNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetQueueCheckpointDataNV.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkCheckpointDataNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCheckpointDataNV.html)

- Extending [VkQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties2.html):

  - [VkQueueFamilyCheckpointPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyCheckpointPropertiesNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_DEVICE_DIAGNOSTIC_CHECKPOINTS_EXTENSION_NAME`

- `VK_NV_DEVICE_DIAGNOSTIC_CHECKPOINTS_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_CHECKPOINT_DATA_NV`

  - `VK_STRUCTURE_TYPE_QUEUE_FAMILY_CHECKPOINT_PROPERTIES_NV`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-07-16 (Nuno Subtil)

  - Internal revisions

- Revision 2, 2018-07-16 (Nuno Subtil)

  - ???

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_device_diagnostic_checkpoints"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
