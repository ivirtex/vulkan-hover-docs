# VkGraphicsPipelineShaderGroupsCreateInfoNV(3) Manual Page

## Name

VkGraphicsPipelineShaderGroupsCreateInfoNV - Structure specifying parameters of a newly created multi shader group pipeline



## [](#_c_specification)C Specification

The `VkGraphicsPipelineShaderGroupsCreateInfoNV` structure is defined as:

```c++
// Provided by VK_NV_device_generated_commands
typedef struct VkGraphicsPipelineShaderGroupsCreateInfoNV {
    VkStructureType                             sType;
    const void*                                 pNext;
    uint32_t                                    groupCount;
    const VkGraphicsShaderGroupCreateInfoNV*    pGroups;
    uint32_t                                    pipelineCount;
    const VkPipeline*                           pPipelines;
} VkGraphicsPipelineShaderGroupsCreateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `groupCount` is the number of elements in the `pGroups` array.
- `pGroups` is a pointer to an array of [VkGraphicsShaderGroupCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsShaderGroupCreateInfoNV.html) structures specifying which state of the original [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html) each shader group overrides.
- `pipelineCount` is the number of elements in the `pPipelines` array.
- `pPipelines` is a pointer to an array of graphics `VkPipeline` structures which are referenced within the created pipeline, including all their shader groups.

## [](#_description)Description

When referencing shader groups by index, groups defined in the referenced pipelines are treated as if they were defined as additional entries in `pGroups`. They are appended in the order they appear in the `pPipelines` array and in the `pGroups` array when those pipelines were defined.

The application **must** maintain the lifetime of all such referenced pipelines based on the pipelines that make use of them.

Valid Usage

- [](#VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-groupCount-02879)VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-groupCount-02879  
  `groupCount` **must** be at least `1` and as maximum `VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV`::`maxGraphicsShaderGroupCount`
- [](#VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-groupCount-02880)VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-groupCount-02880  
  The sum of `groupCount` including those groups added from referenced `pPipelines` **must** also be as maximum `VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV`::`maxGraphicsShaderGroupCount`
- [](#VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pGroups-02881)VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pGroups-02881  
  The state of the first element of `pGroups` **must** match its equivalent within the parent’s [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html)
- [](#VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pGroups-02882)VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pGroups-02882  
  Each element of `pGroups` **must** in combination with the rest of the pipeline state yield a valid state configuration
- [](#VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pGroups-02884)VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pGroups-02884  
  All elements of `pGroups` **must** use the same shader stage combinations unless any mesh shader stage is used, then either combination of task and mesh or just mesh shader is valid
- [](#VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pGroups-02885)VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pGroups-02885  
  Mesh and regular primitive shading stages cannot be mixed across `pGroups`
- [](#VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pPipelines-02886)VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pPipelines-02886  
  Each element of `pPipelines` **must** have been created with identical state to the pipeline currently created except the state that can be overridden by [VkGraphicsShaderGroupCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsShaderGroupCreateInfoNV.html)
- [](#VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-deviceGeneratedCommands-02887)VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-deviceGeneratedCommands-02887  
  The [`VkPhysicalDeviceDeviceGeneratedCommandsFeaturesNV`::`deviceGeneratedCommands`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-deviceGeneratedCommandsNV) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-sType-sType)VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_GRAPHICS_PIPELINE_SHADER_GROUPS_CREATE_INFO_NV`
- [](#VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pGroups-parameter)VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pGroups-parameter  
  If `groupCount` is not `0`, `pGroups` **must** be a valid pointer to an array of `groupCount` valid [VkGraphicsShaderGroupCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsShaderGroupCreateInfoNV.html) structures
- [](#VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pPipelines-parameter)VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pPipelines-parameter  
  If `pipelineCount` is not `0`, `pPipelines` **must** be a valid pointer to an array of `pipelineCount` valid [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handles

## [](#_see_also)See Also

[VK\_NV\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands.html), [VkGraphicsShaderGroupCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsShaderGroupCreateInfoNV.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkGraphicsPipelineShaderGroupsCreateInfoNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0