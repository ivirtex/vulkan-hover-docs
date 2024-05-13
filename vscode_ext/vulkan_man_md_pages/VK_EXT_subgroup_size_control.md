# VK_EXT_subgroup_size_control(3) Manual Page

## Name

VK_EXT_subgroup_size_control - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

226

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[Version 1.1](#versions-1.1)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.3-promotions"
  target="_blank" rel="noopener">Vulkan 1.3</a>

## <a href="#_contact" class="anchor"></a>Contact

- Neil Henning <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_subgroup_size_control%5D%20@sheredom%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_subgroup_size_control%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>sheredom</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-03-05

**Contributors**  
- Jeff Bolz, NVIDIA

- Faith Ekstrand, Intel

- Sławek Grajewski, Intel

- Jesse Hall, Google

- Neil Henning, AMD

- Daniel Koch, NVIDIA

- Jeff Leger, Qualcomm

- Graeme Leese, Broadcom

- Allan MacKinnon, Google

- Mariusz Merecki, Intel

- Graham Wihlidal, Electronic Arts

## <a href="#_description" class="anchor"></a>Description

This extension enables an implementation to control the subgroup size by
allowing a varying subgroup size and also specifying a required subgroup
size.

It extends the subgroup support in Vulkan 1.1 to allow an implementation
to expose a varying subgroup size. Previously Vulkan exposed a single
subgroup size per physical device, with the expectation that
implementations will behave as if all subgroups have the same size. Some
implementations **may** dispatch shaders with a varying subgroup size
for different subgroups. As a result they could implicitly split a large
subgroup into smaller subgroups or represent a small subgroup as a
larger subgroup, some of whose invocations were inactive on launch.

To aid developers in understanding the performance characteristics of
their programs, this extension exposes a minimum and maximum subgroup
size that a physical device supports and a pipeline create flag to
enable that pipeline to vary its subgroup size. If enabled, any
`SubgroupSize` decorated variables in the SPIR-V shader modules provided
to pipeline creation **may** vary between the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-minSubgroupSize"
target="_blank" rel="noopener">minimum</a> and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxSubgroupSize"
target="_blank" rel="noopener">maximum</a> subgroup sizes.

An implementation is also optionally allowed to support specifying a
required subgroup size for a given pipeline stage. Implementations
advertise which <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-requiredSubgroupSizeStages"
target="_blank" rel="noopener">stages support a required subgroup
size</a>, and any pipeline of a supported stage can be passed a
[VkPipelineShaderStageRequiredSubgroupSizeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfoEXT.html)
structure to set the subgroup size for that shader stage of the
pipeline. For compute shaders, this requires the developer to query the
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxComputeWorkgroupSubgroups"
target="_blank"
rel="noopener"><code>maxComputeWorkgroupSubgroups</code></a> and ensure
that:

s=WorkGroupSize.x×WorkGroupSize.y×WorkgroupSize.z≤SubgroupSize×maxComputeWorkgroupSubgroups

Developers can also specify a new pipeline shader stage create flag that
requires the implementation to have fully populated subgroups within
local workgroups. This requires the workgroup size in the X dimension to
be a multiple of the subgroup size.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceSubgroupSizeControlFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSubgroupSizeControlFeaturesEXT.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceSubgroupSizeControlPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSubgroupSizeControlPropertiesEXT.html)

- Extending
  [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html),
  [VkShaderCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCreateInfoEXT.html):

  - [VkPipelineShaderStageRequiredSubgroupSizeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfoEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_SUBGROUP_SIZE_CONTROL_EXTENSION_NAME`

- `VK_EXT_SUBGROUP_SIZE_CONTROL_SPEC_VERSION`

- Extending
  [VkPipelineShaderStageCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateFlagBits.html):

  - `VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT_EXT`

  - `VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_SIZE_CONTROL_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_SIZE_CONTROL_PROPERTIES_EXT`

  - `VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_REQUIRED_SUBGROUP_SIZE_CREATE_INFO_EXT`

## <a href="#_promotion_to_vulkan_1_3" class="anchor"></a>Promotion to Vulkan 1.3

Functionality in this extension is included in core Vulkan 1.3, with the
EXT suffix omitted. The original type, enum and command names are still
available as aliases of the core functionality.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-03-05 (Neil Henning)

  - Initial draft

- Revision 2, 2019-07-26 (Faith Ekstrand)

  - Add the missing
    [VkPhysicalDeviceSubgroupSizeControlFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSubgroupSizeControlFeaturesEXT.html)
    for querying subgroup size control features.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_subgroup_size_control"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
