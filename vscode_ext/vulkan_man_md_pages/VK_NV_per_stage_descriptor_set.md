# VK_NV_per_stage_descriptor_set(3) Manual Page

## Name

VK_NV_per_stage_descriptor_set - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

517

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_maintenance6](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance6.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Piers Daniell <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_per_stage_descriptor_set%5D%20@pdaniell-nv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_per_stage_descriptor_set%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>pdaniell-nv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-10-16

**IP Status**  
No known IP claims.

**Contributors**  
- Daniel Story, Nintendo

## <a href="#_description" class="anchor"></a>Description

This extension introduces a new descriptor set layout creation flag that
allows bindings in a descriptor set to be scoped to each shader stage.
This means that shaders bound at the same time **may** use completely
different descriptor set layouts without any restrictions on
compatibility, and that the descriptor limits that would otherwise apply
to the union of all stages together instead apply to each stage
individually. It also means that descriptors shared by multiple stages
**must** be bound to each stage or set of stages that use a unique
descriptor set layout using their specific per stage descriptor set
layout(s).

This extension also allows each of the new descriptor binding functions
from VK_KHR_maintenance6 to have their
[VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) member be optionally set to
[VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), in which case the pipeline layout
information is taken from a
[VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html) structure
in the `pNext` chain. This enables descriptors to be directly bound
using descriptor set layouts without applications needing to create and
manage [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) objects at command
recording time.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDevicePerStageDescriptorSetFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePerStageDescriptorSetFeaturesNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_PER_STAGE_DESCRIPTOR_SET_EXTENSION_NAME`

- `VK_NV_PER_STAGE_DESCRIPTOR_SET_SPEC_VERSION`

- Extending
  [VkDescriptorSetLayoutCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateFlagBits.html):

  - `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PER_STAGE_BIT_NV`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PER_STAGE_DESCRIPTOR_SET_FEATURES_NV`

## <a href="#_issues" class="anchor"></a>Issues

None

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-10-16 (Piers Daniell)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_per_stage_descriptor_set"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
