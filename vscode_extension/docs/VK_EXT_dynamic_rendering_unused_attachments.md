# VK\_EXT\_dynamic\_rendering\_unused\_attachments(3) Manual Page

## Name

VK\_EXT\_dynamic\_rendering\_unused\_attachments - device extension



## [](#_registered_extension_number)Registered Extension Number

500

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

         [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
         or  
         [Vulkan Version 1.1](#versions-1.1)  
     and  
     [VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html)  
or  
[Vulkan Version 1.3](#versions-1.3)

## [](#_contact)Contact

- Piers Daniell [\[GitHub\]pdaniell-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_dynamic_rendering_unused_attachments%5D%20%40pdaniell-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_dynamic_rendering_unused_attachments%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_dynamic\_rendering\_unused\_attachments](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_dynamic_rendering_unused_attachments.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-05-22

**IP Status**

No known IP claims.

**Contributors**

- Daniel Story, Nintendo
- Hans-Kristian Arntzen, Valve
- Jan-Harald Fredriksen, Arm
- James Fitzpatrick, Imagination Technologies
- Pan Gao, Huawei Technologies
- Ricardo Garcia, Igalia
- Stu Smith, AMD

## [](#_description)Description

This extension lifts some restrictions in the `VK_KHR_dynamic_rendering` extension to allow render pass instances and bound pipelines within those render pass instances to have an unused attachment specified in one but not the other. It also allows pipelines to use different formats in a render pass as long the attachment is NULL.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceDynamicRenderingUnusedAttachmentsFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDynamicRenderingUnusedAttachmentsFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_DYNAMIC_RENDERING_UNUSED_ATTACHMENTS_EXTENSION_NAME`
- `VK_EXT_DYNAMIC_RENDERING_UNUSED_ATTACHMENTS_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DYNAMIC_RENDERING_UNUSED_ATTACHMENTS_FEATURES_EXT`

## [](#_issues)Issues

None.

## [](#_version_history)Version History

- Revision 1, 2023-05-22 (Piers Daniell)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_dynamic_rendering_unused_attachments).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0