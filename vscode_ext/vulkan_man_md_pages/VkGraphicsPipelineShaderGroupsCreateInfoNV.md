# VkGraphicsPipelineShaderGroupsCreateInfoNV(3) Manual Page

## Name

VkGraphicsPipelineShaderGroupsCreateInfoNV - Structure specifying
parameters of a newly created multi shader group pipeline



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkGraphicsPipelineShaderGroupsCreateInfoNV` structure is defined
as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `groupCount` is the number of elements in the `pGroups` array.

- `pGroups` is a pointer to an array of
  [VkGraphicsShaderGroupCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsShaderGroupCreateInfoNV.html)
  structures specifying which state of the original
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html) each
  shader group overrides.

- `pipelineCount` is the number of elements in the `pPipelines` array.

- `pPipelines` is a pointer to an array of graphics `VkPipeline`
  structures which are referenced within the created pipeline, including
  all their shader groups.

## <a href="#_description" class="anchor"></a>Description

When referencing shader groups by index, groups defined in the
referenced pipelines are treated as if they were defined as additional
entries in `pGroups`. They are appended in the order they appear in the
`pPipelines` array and in the `pGroups` array when those pipelines were
defined.

The application **must** maintain the lifetime of all such referenced
pipelines based on the pipelines that make use of them.

Valid Usage

- <a
  href="#VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-groupCount-02879"
  id="VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-groupCount-02879"></a>
  VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-groupCount-02879  
  `groupCount` **must** be at least `1` and as maximum
  `VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV`::`maxGraphicsShaderGroupCount`

- <a
  href="#VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-groupCount-02880"
  id="VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-groupCount-02880"></a>
  VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-groupCount-02880  
  The sum of `groupCount` including those groups added from referenced
  `pPipelines` **must** also be as maximum
  `VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV`::`maxGraphicsShaderGroupCount`

- <a href="#VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pGroups-02881"
  id="VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pGroups-02881"></a>
  VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pGroups-02881  
  The state of the first element of `pGroups` **must** match its
  equivalent within the parent’s
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html)

- <a href="#VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pGroups-02882"
  id="VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pGroups-02882"></a>
  VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pGroups-02882  
  Each element of `pGroups` **must** in combination with the rest of the
  pipeline state yield a valid state configuration

- <a href="#VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pGroups-02884"
  id="VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pGroups-02884"></a>
  VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pGroups-02884  
  All elements of `pGroups` **must** use the same shader stage
  combinations unless any mesh shader stage is used, then either
  combination of task and mesh or just mesh shader is valid

- <a href="#VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pGroups-02885"
  id="VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pGroups-02885"></a>
  VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pGroups-02885  
  Mesh and regular primitive shading stages cannot be mixed across
  `pGroups`

- <a
  href="#VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pPipelines-02886"
  id="VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pPipelines-02886"></a>
  VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pPipelines-02886  
  Each element of `pPipelines` **must** have been created with identical
  state to the pipeline currently created except the state that can be
  overridden by
  [VkGraphicsShaderGroupCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsShaderGroupCreateInfoNV.html)

- <a
  href="#VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-deviceGeneratedCommands-02887"
  id="VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-deviceGeneratedCommands-02887"></a>
  VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-deviceGeneratedCommands-02887  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-deviceGeneratedCommands"
  target="_blank" rel="noopener"><code>deviceGeneratedCommands</code></a>
  feature **must** be enabled

Valid Usage (Implicit)

- <a href="#VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-sType-sType"
  id="VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-sType-sType"></a>
  VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_GRAPHICS_PIPELINE_SHADER_GROUPS_CREATE_INFO_NV`

- <a
  href="#VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pGroups-parameter"
  id="VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pGroups-parameter"></a>
  VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pGroups-parameter  
  If `groupCount` is not `0`, `pGroups` **must** be a valid pointer to
  an array of `groupCount` valid
  [VkGraphicsShaderGroupCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsShaderGroupCreateInfoNV.html)
  structures

- <a
  href="#VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pPipelines-parameter"
  id="VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pPipelines-parameter"></a>
  VUID-VkGraphicsPipelineShaderGroupsCreateInfoNV-pPipelines-parameter  
  If `pipelineCount` is not `0`, `pPipelines` **must** be a valid
  pointer to an array of `pipelineCount` valid
  [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) handles

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_generated_commands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands.html),
[VkGraphicsShaderGroupCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsShaderGroupCreateInfoNV.html),
[VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkGraphicsPipelineShaderGroupsCreateInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
