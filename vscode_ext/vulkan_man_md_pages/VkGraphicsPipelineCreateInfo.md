# VkGraphicsPipelineCreateInfo(3) Manual Page

## Name

VkGraphicsPipelineCreateInfo - Structure specifying parameters of a
newly created graphics pipeline



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkGraphicsPipelineCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkGraphicsPipelineCreateInfo {
    VkStructureType                                  sType;
    const void*                                      pNext;
    VkPipelineCreateFlags                            flags;
    uint32_t                                         stageCount;
    const VkPipelineShaderStageCreateInfo*           pStages;
    const VkPipelineVertexInputStateCreateInfo*      pVertexInputState;
    const VkPipelineInputAssemblyStateCreateInfo*    pInputAssemblyState;
    const VkPipelineTessellationStateCreateInfo*     pTessellationState;
    const VkPipelineViewportStateCreateInfo*         pViewportState;
    const VkPipelineRasterizationStateCreateInfo*    pRasterizationState;
    const VkPipelineMultisampleStateCreateInfo*      pMultisampleState;
    const VkPipelineDepthStencilStateCreateInfo*     pDepthStencilState;
    const VkPipelineColorBlendStateCreateInfo*       pColorBlendState;
    const VkPipelineDynamicStateCreateInfo*          pDynamicState;
    VkPipelineLayout                                 layout;
    VkRenderPass                                     renderPass;
    uint32_t                                         subpass;
    VkPipeline                                       basePipelineHandle;
    int32_t                                          basePipelineIndex;
} VkGraphicsPipelineCreateInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html) specifying
  how the pipeline will be generated.

- `stageCount` is the number of entries in the `pStages` array.

- `pStages` is a pointer to an array of `stageCount`
  [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html)
  structures describing the set of the shader stages to be included in
  the graphics pipeline.

- `pVertexInputState` is a pointer to a
  [VkPipelineVertexInputStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineVertexInputStateCreateInfo.html)
  structure. It is ignored if the pipeline includes a mesh shader stage.
  It **can** be `NULL` if the pipeline is created with the
  `VK_DYNAMIC_STATE_VERTEX_INPUT_EXT` dynamic state set.

- `pInputAssemblyState` is a pointer to a
  [VkPipelineInputAssemblyStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineInputAssemblyStateCreateInfo.html)
  structure which determines input assembly behavior for vertex shading,
  as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#drawing"
  target="_blank" rel="noopener">Drawing Commands</a>.
  <span id="pipelines-pInputAssemblyState-null"></span> If the
  [`VK_EXT_extended_dynamic_state3`](VK_EXT_extended_dynamic_state3.html)
  extension is enabled, it **can** be `NULL` if the pipeline is created
  with both `VK_DYNAMIC_STATE_PRIMITIVE_RESTART_ENABLE`, and
  `VK_DYNAMIC_STATE_PRIMITIVE_TOPOLOGY` dynamic states set and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-dynamicPrimitiveTopologyUnrestricted"
  target="_blank"
  rel="noopener"><code>dynamicPrimitiveTopologyUnrestricted</code></a>
  is `VK_TRUE`. It is ignored if the pipeline includes a mesh shader
  stage.

- `pTessellationState` is a pointer to a
  [VkPipelineTessellationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineTessellationStateCreateInfo.html)
  structure defining tessellation state used by tessellation shaders. It
  **can** be `NULL` if the pipeline is created with the
  `VK_DYNAMIC_STATE_PATCH_CONTROL_POINTS_EXT` dynamic state set.

- `pViewportState` is a pointer to a
  [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html)
  structure defining viewport state used when rasterization is enabled.
  <span id="pipelines-pViewportState-null"></span> If the
  [`VK_EXT_extended_dynamic_state3`](VK_EXT_extended_dynamic_state3.html)
  extension is enabled, it **can** be `NULL` if the pipeline is created
  with both `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT`, and
  `VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT` dynamic states set.

- `pRasterizationState` is a pointer to a
  [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)
  structure defining rasterization state.
  <span id="pipelines-pRasterizationState-null"></span> If the
  [`VK_EXT_extended_dynamic_state3`](VK_EXT_extended_dynamic_state3.html)
  extension is enabled, it **can** be `NULL` if the pipeline is created
  with all of `VK_DYNAMIC_STATE_DEPTH_CLAMP_ENABLE_EXT`,
  `VK_DYNAMIC_STATE_RASTERIZER_DISCARD_ENABLE`,
  `VK_DYNAMIC_STATE_POLYGON_MODE_EXT`, `VK_DYNAMIC_STATE_CULL_MODE`,
  `VK_DYNAMIC_STATE_FRONT_FACE`, `VK_DYNAMIC_STATE_DEPTH_BIAS_ENABLE`,
  `VK_DYNAMIC_STATE_DEPTH_BIAS`, and `VK_DYNAMIC_STATE_LINE_WIDTH`
  dynamic states set.

- `pMultisampleState` is a pointer to a
  [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)
  structure defining multisample state used when rasterization is
  enabled. <span id="pipelines-pMultisampleState-null"></span> If the
  [`VK_EXT_extended_dynamic_state3`](VK_EXT_extended_dynamic_state3.html)
  extension is enabled, it **can** be `NULL` if the pipeline is created
  with all of `VK_DYNAMIC_STATE_RASTERIZATION_SAMPLES_EXT`,
  `VK_DYNAMIC_STATE_SAMPLE_MASK_EXT`, and
  `VK_DYNAMIC_STATE_ALPHA_TO_COVERAGE_ENABLE_EXT` dynamic states set,
  and either <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-alphaToOne"
  target="_blank" rel="noopener">alphaToOne</a> is disabled on the
  device or `VK_DYNAMIC_STATE_ALPHA_TO_ONE_ENABLE_EXT` is set, in which
  case
  [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)::`sampleShadingEnable`
  is assumed to be `VK_FALSE`.

- `pDepthStencilState` is a pointer to a
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
  structure defining depth/stencil state used when rasterization is
  enabled for depth or stencil attachments accessed during rendering.
  <span id="pipelines-pDepthStencilState-null"></span> If the
  [`VK_EXT_extended_dynamic_state3`](VK_EXT_extended_dynamic_state3.html)
  extension is enabled, it **can** be `NULL` if the pipeline is created
  with all of `VK_DYNAMIC_STATE_DEPTH_TEST_ENABLE`,
  `VK_DYNAMIC_STATE_DEPTH_WRITE_ENABLE`,
  `VK_DYNAMIC_STATE_DEPTH_COMPARE_OP`,
  `VK_DYNAMIC_STATE_DEPTH_BOUNDS_TEST_ENABLE`,
  `VK_DYNAMIC_STATE_STENCIL_TEST_ENABLE`, `VK_DYNAMIC_STATE_STENCIL_OP`,
  and `VK_DYNAMIC_STATE_DEPTH_BOUNDS` dynamic states set.

- `pColorBlendState` is a pointer to a
  [VkPipelineColorBlendStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateInfo.html)
  structure defining color blend state used when rasterization is
  enabled for any color attachments accessed during rendering.
  <span id="pipelines-pColorBlendState-null"></span> If the
  [`VK_EXT_extended_dynamic_state3`](VK_EXT_extended_dynamic_state3.html)
  extension is enabled, it **can** be `NULL` if the pipeline is created
  with all of `VK_DYNAMIC_STATE_LOGIC_OP_ENABLE_EXT`,
  `VK_DYNAMIC_STATE_LOGIC_OP_EXT`,
  `VK_DYNAMIC_STATE_COLOR_BLEND_ENABLE_EXT`,
  `VK_DYNAMIC_STATE_COLOR_BLEND_EQUATION_EXT`,
  `VK_DYNAMIC_STATE_COLOR_WRITE_MASK_EXT`, and
  `VK_DYNAMIC_STATE_BLEND_CONSTANTS` dynamic states set.

- `pDynamicState` is a pointer to a
  [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)
  structure defining which properties of the pipeline state object are
  dynamic and **can** be changed independently of the pipeline state.
  This **can** be `NULL`, which means no state in the pipeline is
  considered dynamic.

- `layout` is the description of binding locations used by both the
  pipeline and descriptor sets used with the pipeline.

- `renderPass` is a handle to a render pass object describing the
  environment in which the pipeline will be used. The pipeline **must**
  only be used with a render pass instance compatible with the one
  provided. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-compatibility"
  target="_blank" rel="noopener">Render Pass Compatibility</a> for more
  information.

- `subpass` is the index of the subpass in the render pass where this
  pipeline will be used.

- `basePipelineHandle` is a pipeline to derive from.

- `basePipelineIndex` is an index into the `pCreateInfos` parameter to
  use as a pipeline to derive from.

## <a href="#_description" class="anchor"></a>Description

The parameters `basePipelineHandle` and `basePipelineIndex` are
described in more detail in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-pipeline-derivatives"
target="_blank" rel="noopener">Pipeline Derivatives</a>.

If any shader stage fails to compile, the compile log will be reported
back to the application, and `VK_ERROR_INVALID_SHADER_NV` will be
generated.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>With <a
href="VK_EXT_extended_dynamic_state3.html"><code>VK_EXT_extended_dynamic_state3</code></a>,
it is possible that many of the
<code>VkGraphicsPipelineCreateInfo</code> members above
<strong>can</strong> be <code>NULL</code> because all their state is
dynamic and therefore ignored. This is optional so the application
<strong>can</strong> still use a valid pointer if it needs to set the
<code>pNext</code> or <code>flags</code> fields to specify state for
other extensions.</p></td>
</tr>
</tbody>
</table>

The state required for a graphics pipeline is divided into <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-vertex-input"
target="_blank" rel="noopener">vertex input state</a>, <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization shader state</a>, <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
target="_blank" rel="noopener">fragment shader state</a>, and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
target="_blank" rel="noopener">fragment output state</a>.

Vertex Input State

Vertex input state is defined by:

- [VkPipelineVertexInputStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineVertexInputStateCreateInfo.html)

- [VkPipelineInputAssemblyStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineInputAssemblyStateCreateInfo.html)

If this pipeline specifies <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization state</a> either
directly or by including it as a pipeline library and its `pStages`
includes a vertex shader, this state **must** be specified to create a
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-complete"
target="_blank" rel="noopener">complete graphics pipeline</a>.

If a pipeline includes
`VK_GRAPHICS_PIPELINE_LIBRARY_VERTEX_INPUT_INTERFACE_BIT_EXT` in
[VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
either explicitly or as a default, and either the conditions requiring
this state for a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-complete"
target="_blank" rel="noopener">complete graphics pipeline</a> are met or
this pipeline does not specify <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization state</a> in any way,
that pipeline **must** specify this state directly.

Pre-Rasterization Shader State

Pre-rasterization shader state is defined by:

- [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html)
  entries for:

  - Vertex shaders

  - Tessellation control shaders

  - Tessellation evaluation shaders

  - Geometry shaders

  - Task shaders

  - Mesh shaders

- Within the [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html), all descriptor
  sets with pre-rasterization shader bindings if
  `VK_PIPELINE_LAYOUT_CREATE_INDEPENDENT_SETS_BIT_EXT` was specified.

  - If `VK_PIPELINE_LAYOUT_CREATE_INDEPENDENT_SETS_BIT_EXT` was not
    specified, the full pipeline layout must be specified.

- [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html)

- [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)

- [VkPipelineTessellationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineTessellationStateCreateInfo.html)

- [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html) and `subpass` parameter

- The `viewMask` parameter of
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)
  (formats are ignored)

- [VkPipelineDiscardRectangleStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDiscardRectangleStateCreateInfoEXT.html)

- [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html)

This state **must** be specified to create a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-complete"
target="_blank" rel="noopener">complete graphics pipeline</a>.

If either the `pNext` chain includes a
[VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)
structure with
`VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT`
included in `flags`, or it is not specified and would default to include
that value, this state **must** be specified in the pipeline.

Fragment Shader State

Fragment shader state is defined by:

- A
  [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html)
  entry for the fragment shader

- Within the [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html), all descriptor
  sets with fragment shader bindings if
  `VK_PIPELINE_LAYOUT_CREATE_INDEPENDENT_SETS_BIT_EXT` was specified.

  - If `VK_PIPELINE_LAYOUT_CREATE_INDEPENDENT_SETS_BIT_EXT` was not
    specified, the full pipeline layout must be specified.

- [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)
  if <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-sampleshading"
  target="_blank" rel="noopener">sample shading</a> is enabled or
  `renderpass` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)

- [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html) and `subpass` parameter

- The `viewMask` parameter of
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)
  (formats are ignored)

- [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html)

- [VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html)

- [VkPipelineRepresentativeFragmentTestStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRepresentativeFragmentTestStateCreateInfoNV.html)

- Inclusion/omission of the
  `VK_PIPELINE_RASTERIZATION_STATE_CREATE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`
  flag

- Inclusion/omission of the
  `VK_PIPELINE_RASTERIZATION_STATE_CREATE_FRAGMENT_DENSITY_MAP_ATTACHMENT_BIT_EXT`
  flag

- [VkRenderingInputAttachmentIndexInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInputAttachmentIndexInfoKHR.html)

If a pipeline specifies <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization state</a> either
directly or by including it as a pipeline library and
`rasterizerDiscardEnable` is set to `VK_FALSE` or
`VK_DYNAMIC_STATE_RASTERIZER_DISCARD_ENABLE` is used, this state
**must** be specified to create a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-complete"
target="_blank" rel="noopener">complete graphics pipeline</a>.

If a pipeline includes
`VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT` in
[VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
either explicitly or as a default, and either the conditions requiring
this state for a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-complete"
target="_blank" rel="noopener">complete graphics pipeline</a> are met or
this pipeline does not specify <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization state</a> in any way,
that pipeline **must** specify this state directly.

Fragment Output State

Fragment output state is defined by:

- [VkPipelineColorBlendStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateInfo.html)

- [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html) and `subpass` parameter

- [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)

- [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)

- [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoAMD.html)

- [VkAttachmentSampleCountInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoNV.html)

- [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)

- Inclusion/omission of the
  `VK_PIPELINE_CREATE_COLOR_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT` and
  `VK_PIPELINE_CREATE_DEPTH_STENCIL_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT`
  flags

- Inclusion/omission of the
  `VK_PIPELINE_CREATE_2_ENABLE_LEGACY_DITHERING_BIT_EXT` flag

- [VkRenderingAttachmentLocationInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingAttachmentLocationInfoKHR.html)

If a pipeline specifies <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization state</a> either
directly or by including it as a pipeline library and
`rasterizerDiscardEnable` is set to `VK_FALSE` or
`VK_DYNAMIC_STATE_RASTERIZER_DISCARD_ENABLE` is used, this state
**must** be specified to create a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-complete"
target="_blank" rel="noopener">complete graphics pipeline</a>.

If a pipeline includes
`VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_OUTPUT_INTERFACE_BIT_EXT` in
[VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
either explicitly or as a default, and either the conditions requiring
this state for a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-complete"
target="_blank" rel="noopener">complete graphics pipeline</a> are met or
this pipeline does not specify <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization state</a> in any way,
that pipeline **must** specify this state directly.

Dynamic State

Dynamic state values set via `pDynamicState` **must** be ignored if the
state they correspond to is not otherwise statically set by one of the
state subsets used to create the pipeline. Additionally, setting dynamic
state values **must** not modify whether state in a linked library is
static or dynamic; this is set and unchangeable when the library is
created. For example, if a pipeline only included <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization shader state</a>, then
any dynamic state value corresponding to depth or stencil testing has no
effect. Any linked library that has dynamic state enabled that same
dynamic state **must** also be enabled in all the other linked libraries
to which that dynamic state applies.

Complete Graphics Pipelines

A complete graphics pipeline always includes <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization shader state</a>, with
other subsets included depending on that state as specified in the above
sections.

Graphics Pipeline Library Layouts

If different subsets are linked together with pipeline layouts created
with `VK_PIPELINE_LAYOUT_CREATE_INDEPENDENT_SETS_BIT_EXT`, the final
effective pipeline layout is effectively the union of the linked
pipeline layouts. When binding descriptor sets for this pipeline, the
pipeline layout used **must** be compatible with this union. This
pipeline layout **can** be overridden when linking with
`VK_PIPELINE_CREATE_LINK_TIME_OPTIMIZATION_BIT_EXT` by providing a
[VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) that is <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-compatibility"
target="_blank" rel="noopener">compatible</a> with this union other than
`VK_PIPELINE_LAYOUT_CREATE_INDEPENDENT_SETS_BIT_EXT`, or when linking
without `VK_PIPELINE_CREATE_LINK_TIME_OPTIMIZATION_BIT_EXT` by providing
a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) that is fully <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-compatibility"
target="_blank" rel="noopener">compatible</a> with this union.

If the `pNext` chain includes a
[VkPipelineCreateFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlags2CreateInfoKHR.html)
structure,
[VkPipelineCreateFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlags2CreateInfoKHR.html)::`flags`
from that structure is used instead of `flags` from this structure.

Valid Usage

- <a href="#VUID-VkGraphicsPipelineCreateInfo-None-09497"
  id="VUID-VkGraphicsPipelineCreateInfo-None-09497"></a>
  VUID-VkGraphicsPipelineCreateInfo-None-09497  
  If the `pNext` chain does not include a
  [VkPipelineCreateFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlags2CreateInfoKHR.html)
  structure, `flags` must be a valid combination of
  [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html) values

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-07984"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-07984"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-07984  
  If `flags` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, and
  `basePipelineIndex` is -1, `basePipelineHandle` **must** be a valid
  graphics `VkPipeline` handle

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-07985"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-07985"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-07985  
  If `flags` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, and
  `basePipelineHandle` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `basePipelineIndex` **must** be a valid index into the calling
  command’s `pCreateInfos` parameter

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-07986"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-07986"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-07986  
  If `flags` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag,
  `basePipelineIndex` **must** be -1 or `basePipelineHandle` **must** be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkGraphicsPipelineCreateInfo-layout-07987"
  id="VUID-VkGraphicsPipelineCreateInfo-layout-07987"></a>
  VUID-VkGraphicsPipelineCreateInfo-layout-07987  
  If a push constant block is declared in a shader, a push constant
  range in `layout` **must** match both the shader stage and range

- <a href="#VUID-VkGraphicsPipelineCreateInfo-layout-07988"
  id="VUID-VkGraphicsPipelineCreateInfo-layout-07988"></a>
  VUID-VkGraphicsPipelineCreateInfo-layout-07988  
  If a [resource variables](#interfaces-resources) is declared in a
  shader, a descriptor slot in `layout` **must** match the shader stage

- <a href="#VUID-VkGraphicsPipelineCreateInfo-layout-07990"
  id="VUID-VkGraphicsPipelineCreateInfo-layout-07990"></a>
  VUID-VkGraphicsPipelineCreateInfo-layout-07990  
  If a [resource variables](#interfaces-resources) is declared in a
  shader, and the descriptor type is not
  `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`, a descriptor slot in `layout`
  **must** match the descriptor type

- <a href="#VUID-VkGraphicsPipelineCreateInfo-layout-07991"
  id="VUID-VkGraphicsPipelineCreateInfo-layout-07991"></a>
  VUID-VkGraphicsPipelineCreateInfo-layout-07991  
  If a [resource variables](#interfaces-resources) is declared in a
  shader as an array, a descriptor slot in `layout` **must** match the
  descriptor count

- <a href="#VUID-VkGraphicsPipelineCreateInfo-stage-02096"
  id="VUID-VkGraphicsPipelineCreateInfo-stage-02096"></a>
  VUID-VkGraphicsPipelineCreateInfo-stage-02096  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> the
  `stage` member of one element of `pStages` **must** be
  `VK_SHADER_STAGE_VERTEX_BIT` or `VK_SHADER_STAGE_MESH_BIT_EXT`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pStages-02095"
  id="VUID-VkGraphicsPipelineCreateInfo-pStages-02095"></a>
  VUID-VkGraphicsPipelineCreateInfo-pStages-02095  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> the
  geometric shader stages provided in `pStages` **must** be either from
  the mesh shading pipeline (`stage` is `VK_SHADER_STAGE_TASK_BIT_EXT`
  or `VK_SHADER_STAGE_MESH_BIT_EXT`) or from the primitive shading
  pipeline (`stage` is `VK_SHADER_STAGE_VERTEX_BIT`,
  `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT`,
  `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`, or
  `VK_SHADER_STAGE_GEOMETRY_BIT`)

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pStages-09631"
  id="VUID-VkGraphicsPipelineCreateInfo-pStages-09631"></a>
  VUID-VkGraphicsPipelineCreateInfo-pStages-09631  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> and
  `pStages` contains both `VK_SHADER_STAGE_TASK_BIT_EXT` and
  `VK_SHADER_STAGE_MESH_BIT_EXT`, then the mesh shader’s entry point
  **must** not declare a variable with a `DrawIndex` `BuiltIn`
  decoration

- <a href="#VUID-VkGraphicsPipelineCreateInfo-TaskNV-07063"
  id="VUID-VkGraphicsPipelineCreateInfo-TaskNV-07063"></a>
  VUID-VkGraphicsPipelineCreateInfo-TaskNV-07063  
  The shader stages for `VK_SHADER_STAGE_TASK_BIT_EXT` or
  `VK_SHADER_STAGE_MESH_BIT_EXT` **must** use either the `TaskNV` and
  `MeshNV` `Execution` `Model` or the `TaskEXT` and `MeshEXT`
  `Execution` `Model`, but **must** not use both

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pStages-00729"
  id="VUID-VkGraphicsPipelineCreateInfo-pStages-00729"></a>
  VUID-VkGraphicsPipelineCreateInfo-pStages-00729  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> and
  `pStages` includes a tessellation control shader stage, it **must**
  include a tessellation evaluation shader stage

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pStages-00730"
  id="VUID-VkGraphicsPipelineCreateInfo-pStages-00730"></a>
  VUID-VkGraphicsPipelineCreateInfo-pStages-00730  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> and
  `pStages` includes a tessellation evaluation shader stage, it **must**
  include a tessellation control shader stage

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pStages-09022"
  id="VUID-VkGraphicsPipelineCreateInfo-pStages-09022"></a>
  VUID-VkGraphicsPipelineCreateInfo-pStages-09022  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> and
  `pStages` includes a tessellation control shader stage, and the
  [`VK_EXT_extended_dynamic_state3`](VK_EXT_extended_dynamic_state3.html)
  extension is not enabled or the
  `VK_DYNAMIC_STATE_PATCH_CONTROL_POINTS_EXT` dynamic state is not set,
  `pTessellationState` **must** be a valid pointer to a valid
  [VkPipelineTessellationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineTessellationStateCreateInfo.html)
  structure

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pTessellationState-09023"
  id="VUID-VkGraphicsPipelineCreateInfo-pTessellationState-09023"></a>
  VUID-VkGraphicsPipelineCreateInfo-pTessellationState-09023  
  If `pTessellationState` is not `NULL` it **must** be a pointer to a
  valid
  [VkPipelineTessellationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineTessellationStateCreateInfo.html)
  structure

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pStages-00732"
  id="VUID-VkGraphicsPipelineCreateInfo-pStages-00732"></a>
  VUID-VkGraphicsPipelineCreateInfo-pStages-00732  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> and
  `pStages` includes tessellation shader stages, the shader code of at
  least one stage **must** contain an `OpExecutionMode` instruction
  specifying the type of subdivision in the pipeline

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pStages-00733"
  id="VUID-VkGraphicsPipelineCreateInfo-pStages-00733"></a>
  VUID-VkGraphicsPipelineCreateInfo-pStages-00733  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> and
  `pStages` includes tessellation shader stages, and the shader code of
  both stages contain an `OpExecutionMode` instruction specifying the
  type of subdivision in the pipeline, they **must** both specify the
  same subdivision mode

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pStages-00734"
  id="VUID-VkGraphicsPipelineCreateInfo-pStages-00734"></a>
  VUID-VkGraphicsPipelineCreateInfo-pStages-00734  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> and
  `pStages` includes tessellation shader stages, the shader code of at
  least one stage **must** contain an `OpExecutionMode` instruction
  specifying the output patch size in the pipeline

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pStages-00735"
  id="VUID-VkGraphicsPipelineCreateInfo-pStages-00735"></a>
  VUID-VkGraphicsPipelineCreateInfo-pStages-00735  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> and
  `pStages` includes tessellation shader stages, and the shader code of
  both contain an `OpExecutionMode` instruction specifying the out patch
  size in the pipeline, they **must** both specify the same patch size

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pStages-08888"
  id="VUID-VkGraphicsPipelineCreateInfo-pStages-08888"></a>
  VUID-VkGraphicsPipelineCreateInfo-pStages-08888  
  If the pipeline is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> and
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-vertex-input"
  target="_blank" rel="noopener">vertex input state</a> and `pStages`
  includes tessellation shader stages, and either
  `VK_DYNAMIC_STATE_PRIMITIVE_TOPOLOGY` dynamic state is not enabled or
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-dynamicPrimitiveTopologyUnrestricted"
  target="_blank"
  rel="noopener"><code>dynamicPrimitiveTopologyUnrestricted</code></a>
  is `VK_FALSE`, the `topology` member of `pInputAssembly` **must** be
  `VK_PRIMITIVE_TOPOLOGY_PATCH_LIST`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-topology-08889"
  id="VUID-VkGraphicsPipelineCreateInfo-topology-08889"></a>
  VUID-VkGraphicsPipelineCreateInfo-topology-08889  
  If the pipeline is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> and
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-vertex-input"
  target="_blank" rel="noopener">vertex input state</a> and the
  `topology` member of `pInputAssembly` is
  `VK_PRIMITIVE_TOPOLOGY_PATCH_LIST`, and either
  `VK_DYNAMIC_STATE_PRIMITIVE_TOPOLOGY` dynamic state is not enabled or
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-dynamicPrimitiveTopologyUnrestricted"
  target="_blank"
  rel="noopener"><code>dynamicPrimitiveTopologyUnrestricted</code></a>
  is `VK_FALSE`, then `pStages` **must** include tessellation shader
  stages

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-TessellationEvaluation-07723"
  id="VUID-VkGraphicsPipelineCreateInfo-TessellationEvaluation-07723"></a>
  VUID-VkGraphicsPipelineCreateInfo-TessellationEvaluation-07723  
  If the pipeline is being created with a `TessellationEvaluation`
  `Execution` `Model`, no `Geometry` `Execution` `Model`, uses the
  `PointMode` `Execution` `Mode`, and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderTessellationAndGeometryPointSize"
  target="_blank"
  rel="noopener"><code>shaderTessellationAndGeometryPointSize</code></a>
  is enabled, a `PointSize` decorated variable **must** be written to if
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance5"
  target="_blank" rel="noopener"><code>maintenance5</code></a> is not
  enabled

- <a href="#VUID-VkGraphicsPipelineCreateInfo-topology-08773"
  id="VUID-VkGraphicsPipelineCreateInfo-topology-08773"></a>
  VUID-VkGraphicsPipelineCreateInfo-topology-08773  
  If the pipeline is being created with a `Vertex` `Execution` `Model`
  and no `TessellationEvaluation` or `Geometry` `Execution` `Model`, and
  the `topology` member of `pInputAssembly` is
  `VK_PRIMITIVE_TOPOLOGY_POINT_LIST`, and either
  `VK_DYNAMIC_STATE_PRIMITIVE_TOPOLOGY` dynamic state is not enabled or
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-dynamicPrimitiveTopologyUnrestricted"
  target="_blank"
  rel="noopener"><code>dynamicPrimitiveTopologyUnrestricted</code></a>
  is `VK_FALSE`, a `PointSize` decorated variable **must** be written to
  if <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance5"
  target="_blank" rel="noopener"><code>maintenance5</code></a> is not
  enabled

- <a href="#VUID-VkGraphicsPipelineCreateInfo-maintenance5-08775"
  id="VUID-VkGraphicsPipelineCreateInfo-maintenance5-08775"></a>
  VUID-VkGraphicsPipelineCreateInfo-maintenance5-08775  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance5"
  target="_blank" rel="noopener"><code>maintenance5</code></a> is
  enabled and a `PointSize` decorated variable is written to, all
  execution paths **must** write to a `PointSize` decorated variable

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-TessellationEvaluation-07724"
  id="VUID-VkGraphicsPipelineCreateInfo-TessellationEvaluation-07724"></a>
  VUID-VkGraphicsPipelineCreateInfo-TessellationEvaluation-07724  
  If the pipeline is being created with a `TessellationEvaluation`
  `Execution` `Model`, no `Geometry` `Execution` `Model`, uses the
  `PointMode` `Execution` `Mode`, and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderTessellationAndGeometryPointSize"
  target="_blank"
  rel="noopener"><code>shaderTessellationAndGeometryPointSize</code></a>
  is not enabled, a `PointSize` decorated variable **must** not be
  written to

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-shaderTessellationAndGeometryPointSize-08776"
  id="VUID-VkGraphicsPipelineCreateInfo-shaderTessellationAndGeometryPointSize-08776"></a>
  VUID-VkGraphicsPipelineCreateInfo-shaderTessellationAndGeometryPointSize-08776  
  If the pipeline is being created with a `Geometry` `Execution`
  `Model`, uses the `OutputPoints` `Execution` `Mode`, and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderTessellationAndGeometryPointSize"
  target="_blank"
  rel="noopener"><code>shaderTessellationAndGeometryPointSize</code></a>
  is enabled, a `PointSize` decorated variable **must** be written to
  for every vertex emitted if <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance5"
  target="_blank" rel="noopener"><code>maintenance5</code></a> is not
  enabled

- <a href="#VUID-VkGraphicsPipelineCreateInfo-Geometry-07726"
  id="VUID-VkGraphicsPipelineCreateInfo-Geometry-07726"></a>
  VUID-VkGraphicsPipelineCreateInfo-Geometry-07726  
  If the pipeline is being created with a `Geometry` `Execution`
  `Model`, uses the `OutputPoints` `Execution` `Mode`, and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderTessellationAndGeometryPointSize"
  target="_blank"
  rel="noopener"><code>shaderTessellationAndGeometryPointSize</code></a>
  is not enabled, a `PointSize` decorated variable **must** not be
  written to

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pStages-00738"
  id="VUID-VkGraphicsPipelineCreateInfo-pStages-00738"></a>
  VUID-VkGraphicsPipelineCreateInfo-pStages-00738  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> and
  `pStages` includes a geometry shader stage, and does not include any
  tessellation shader stages, its shader code **must** contain an
  `OpExecutionMode` instruction specifying an input primitive type that
  is <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-geometry-execution"
  target="_blank" rel="noopener">compatible</a> with the primitive
  topology specified in `pInputAssembly`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pStages-00739"
  id="VUID-VkGraphicsPipelineCreateInfo-pStages-00739"></a>
  VUID-VkGraphicsPipelineCreateInfo-pStages-00739  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> and
  `pStages` includes a geometry shader stage, and also includes
  tessellation shader stages, its shader code **must** contain an
  `OpExecutionMode` instruction specifying an input primitive type that
  is <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-geometry-execution"
  target="_blank" rel="noopener">compatible</a> with the primitive
  topology that is output by the tessellation stages

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pStages-00740"
  id="VUID-VkGraphicsPipelineCreateInfo-pStages-00740"></a>
  VUID-VkGraphicsPipelineCreateInfo-pStages-00740  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> and
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>, it includes
  both a fragment shader and a geometry shader, and the fragment shader
  code reads from an input variable that is decorated with
  `PrimitiveId`, then the geometry shader code **must** write to a
  matching output variable, decorated with `PrimitiveId`, in all
  execution paths

- <a href="#VUID-VkGraphicsPipelineCreateInfo-PrimitiveId-06264"
  id="VUID-VkGraphicsPipelineCreateInfo-PrimitiveId-06264"></a>
  VUID-VkGraphicsPipelineCreateInfo-PrimitiveId-06264  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, it
  includes a mesh shader and the fragment shader code reads from an
  input variable that is decorated with `PrimitiveId`, then the mesh
  shader code **must** write to a matching output variable, decorated
  with `PrimitiveId`, in all execution paths

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-06038"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-06038"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-06038  
  If `renderPass` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and the
  pipeline is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> the fragment
  shader **must** not read from any input attachment that is defined as
  `VK_ATTACHMENT_UNUSED` in `subpass`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pStages-00742"
  id="VUID-VkGraphicsPipelineCreateInfo-pStages-00742"></a>
  VUID-VkGraphicsPipelineCreateInfo-pStages-00742  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> and
  multiple pre-rasterization shader stages are included in `pStages`,
  the shader code for the entry points identified by those `pStages` and
  the rest of the state identified by this structure **must** adhere to
  the pipeline linking rules described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces"
  target="_blank" rel="noopener">Shader Interfaces</a> chapter

- <a href="#VUID-VkGraphicsPipelineCreateInfo-None-04889"
  id="VUID-VkGraphicsPipelineCreateInfo-None-04889"></a>
  VUID-VkGraphicsPipelineCreateInfo-None-04889  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> and
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>, the fragment
  shader and last <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader stage</a> and
  any relevant state **must** adhere to the pipeline linking rules
  described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces"
  target="_blank" rel="noopener">Shader Interfaces</a> chapter

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-06041"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-06041"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-06041  
  If `renderPass` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and the
  pipeline is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  then for each color attachment in the subpass, if the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#potential-format-features"
  target="_blank" rel="noopener">potential format features</a> of the
  format of the corresponding attachment description do not contain
  `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BLEND_BIT`, then the `blendEnable`
  member of the corresponding element of the `pAttachments` member of
  `pColorBlendState` **must** be `VK_FALSE`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-07609"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-07609"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-07609  
  If `renderPass` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the
  pipeline is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  the `pColorBlendState` pointer is not `NULL`, the `attachmentCount`
  member of `pColorBlendState` is not ignored, and the subpass uses
  color attachments, the `attachmentCount` member of `pColorBlendState`
  **must** be equal to the `colorAttachmentCount` used to create
  `subpass`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04130"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04130"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04130  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  `pViewportState->pViewports` is not dynamic, then
  `pViewportState->pViewports` **must** be a valid pointer to an array
  of `pViewportState->viewportCount` valid `VkViewport` structures

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04131"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04131"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04131  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  `pViewportState->pScissors` is not dynamic, then
  `pViewportState->pScissors` **must** be a valid pointer to an array of
  `pViewportState->scissorCount` `VkRect2D` structures

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-00749"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-00749"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-00749  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-wideLines"
  target="_blank" rel="noopener"><code>wideLines</code></a> feature is
  not enabled, and no element of the `pDynamicStates` member of
  `pDynamicState` is `VK_DYNAMIC_STATE_LINE_WIDTH`, the `lineWidth`
  member of `pRasterizationState` **must** be `1.0`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-rasterizerDiscardEnable-09024"
  id="VUID-VkGraphicsPipelineCreateInfo-rasterizerDiscardEnable-09024"></a>
  VUID-VkGraphicsPipelineCreateInfo-rasterizerDiscardEnable-09024  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  the `VK_DYNAMIC_STATE_RASTERIZER_DISCARD_ENABLE` dynamic state is
  enabled or the `rasterizerDiscardEnable` member of
  `pRasterizationState` is `VK_FALSE`, and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-pViewportState-null"
  target="_blank" rel="noopener">related dynamic state is not set</a>,
  `pViewportState` **must** be a valid pointer to a valid
  [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html)
  structure

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pViewportState-09025"
  id="VUID-VkGraphicsPipelineCreateInfo-pViewportState-09025"></a>
  VUID-VkGraphicsPipelineCreateInfo-pViewportState-09025  
  If `pViewportState` is not `NULL` it **must** be a valid pointer to a
  valid
  [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html)
  structure

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pMultisampleState-09026"
  id="VUID-VkGraphicsPipelineCreateInfo-pMultisampleState-09026"></a>
  VUID-VkGraphicsPipelineCreateInfo-pMultisampleState-09026  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  and the
  [`VK_EXT_extended_dynamic_state3`](VK_EXT_extended_dynamic_state3.html)
  extension is not enabled or any of the
  `VK_DYNAMIC_STATE_RASTERIZATION_SAMPLES_EXT`,
  `VK_DYNAMIC_STATE_SAMPLE_MASK_EXT`, or
  `VK_DYNAMIC_STATE_ALPHA_TO_COVERAGE_ENABLE_EXT` dynamic states is not
  set, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-alphaToOne"
  target="_blank" rel="noopener">alphaToOne</a> is enabled on the device
  and `VK_DYNAMIC_STATE_ALPHA_TO_ONE_ENABLE_EXT` is not set,
  `pMultisampleState` **must** be a valid pointer to a valid
  [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)
  structure

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pMultisampleState-09027"
  id="VUID-VkGraphicsPipelineCreateInfo-pMultisampleState-09027"></a>
  VUID-VkGraphicsPipelineCreateInfo-pMultisampleState-09027  
  If `pMultisampleState` is not `NULL` it **must** be a valid pointer to
  a valid
  [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)
  structure

- <a href="#VUID-VkGraphicsPipelineCreateInfo-alphaToCoverageEnable-08891"
  id="VUID-VkGraphicsPipelineCreateInfo-alphaToCoverageEnable-08891"></a>
  VUID-VkGraphicsPipelineCreateInfo-alphaToCoverageEnable-08891  
  If the pipeline is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>, the
  [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)::`alphaToCoverageEnable`
  is not ignored and is `VK_TRUE`, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-fragmentoutput"
  target="_blank" rel="noopener">Fragment Output Interface</a> **must**
  contain a variable for the alpha `Component` word in `Location` 0 at
  `Index` 0

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-09028"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-09028"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-09028  
  If `renderPass` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the
  pipeline is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>, and
  `subpass` uses a depth/stencil attachment, and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-pDepthStencilState-null"
  target="_blank" rel="noopener">related dynamic state is not set</a>,
  `pDepthStencilState` **must** be a valid pointer to a valid
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
  structure

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDepthStencilState-09029"
  id="VUID-VkGraphicsPipelineCreateInfo-pDepthStencilState-09029"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDepthStencilState-09029  
  If `pDepthStencilState` is not `NULL` it **must** be a valid pointer
  to a valid
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
  structure

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-09030"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-09030"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-09030  
  If `renderPass` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the
  pipeline is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  and `subpass` uses color attachments, and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-pColorBlendState-null"
  target="_blank" rel="noopener">related dynamic state is not set</a>,
  `pColorBlendState` **must** be a valid pointer to a valid
  [VkPipelineColorBlendStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateInfo.html)
  structure

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-00754"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-00754"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-00754  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, the
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-depthBiasClamp"
  target="_blank" rel="noopener"><code>depthBiasClamp</code></a> feature
  is not enabled, no element of the `pDynamicStates` member of
  `pDynamicState` is `VK_DYNAMIC_STATE_DEPTH_BIAS`, and the
  `depthBiasEnable` member of `pRasterizationState` is `VK_TRUE`, the
  `depthBiasClamp` member of `pRasterizationState` **must** be `0.0`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-02510"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-02510"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-02510  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>, the
  [`VK_EXT_depth_range_unrestricted`](VK_EXT_depth_range_unrestricted.html)
  extension is not enabled and no element of the `pDynamicStates` member
  of `pDynamicState` is `VK_DYNAMIC_STATE_DEPTH_BOUNDS`, and the
  `depthBoundsTestEnable` member of `pDepthStencilState` is `VK_TRUE`,
  the `minDepthBounds` and `maxDepthBounds` members of
  `pDepthStencilState` **must** be between `0.0` and `1.0`, inclusive

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07610"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07610"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07610  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  and `rasterizationSamples` and `sampleLocationsInfo` are not dynamic,
  and
  [VkPipelineSampleLocationsStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineSampleLocationsStateCreateInfoEXT.html)::`sampleLocationsEnable`
  included in the `pNext` chain of `pMultisampleState` is `VK_TRUE`,
  `sampleLocationsInfo.sampleLocationGridSize.width` **must** evenly
  divide
  [VkMultisamplePropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisamplePropertiesEXT.html)::`sampleLocationGridSize.width`
  as returned by
  [vkGetPhysicalDeviceMultisamplePropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceMultisamplePropertiesEXT.html)
  with a `samples` parameter equaling `rasterizationSamples`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07611"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07611"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07611  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  and `rasterizationSamples` and `sampleLocationsInfo` are not dynamic,
  and
  [VkPipelineSampleLocationsStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineSampleLocationsStateCreateInfoEXT.html)::`sampleLocationsEnable`
  the included in the `pNext` chain of `pMultisampleState` is `VK_TRUE`
  or `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_ENABLE_EXT` is used,
  `sampleLocationsInfo.sampleLocationGridSize.height` **must** evenly
  divide
  [VkMultisamplePropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisamplePropertiesEXT.html)::`sampleLocationGridSize.height`
  as returned by
  [vkGetPhysicalDeviceMultisamplePropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceMultisamplePropertiesEXT.html)
  with a `samples` parameter equaling `rasterizationSamples`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07612"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07612"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07612  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  and `rasterizationSamples` and `sampleLocationsInfo` are not dynamic,
  and
  [VkPipelineSampleLocationsStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineSampleLocationsStateCreateInfoEXT.html)::`sampleLocationsEnable`
  included in the `pNext` chain of `pMultisampleState` is `VK_TRUE` or
  `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_ENABLE_EXT` is used,
  `sampleLocationsInfo.sampleLocationsPerPixel` **must** equal
  `rasterizationSamples`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-sampleLocationsEnable-01524"
  id="VUID-VkGraphicsPipelineCreateInfo-sampleLocationsEnable-01524"></a>
  VUID-VkGraphicsPipelineCreateInfo-sampleLocationsEnable-01524  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>, and the
  `sampleLocationsEnable` member of a
  [VkPipelineSampleLocationsStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineSampleLocationsStateCreateInfoEXT.html)
  structure included in the `pNext` chain of `pMultisampleState` is
  `VK_TRUE`, the fragment shader code **must** not statically use the
  extended instruction `InterpolateAtSample`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-multisampledRenderToSingleSampled-06853"
  id="VUID-VkGraphicsPipelineCreateInfo-multisampledRenderToSingleSampled-06853"></a>
  VUID-VkGraphicsPipelineCreateInfo-multisampledRenderToSingleSampled-06853  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  and none of the
  [`VK_AMD_mixed_attachment_samples`](VK_AMD_mixed_attachment_samples.html)
  extension, the
  [`VK_NV_framebuffer_mixed_samples`](VK_NV_framebuffer_mixed_samples.html)
  extension, or the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multisampledRenderToSingleSampled"
  target="_blank"
  rel="noopener"><code>multisampledRenderToSingleSampled</code></a>
  feature are enabled, `rasterizationSamples` is not dynamic, and if
  `subpass` uses color and/or depth/stencil attachments, then the
  `rasterizationSamples` member of `pMultisampleState` **must** be the
  same as the sample count for those subpass attachments

- <a href="#VUID-VkGraphicsPipelineCreateInfo-subpass-01505"
  id="VUID-VkGraphicsPipelineCreateInfo-subpass-01505"></a>
  VUID-VkGraphicsPipelineCreateInfo-subpass-01505  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  and the
  [`VK_AMD_mixed_attachment_samples`](VK_AMD_mixed_attachment_samples.html)
  extension is enabled, `rasterizationSamples` is not dynamic, and if
  `subpass` uses color and/or depth/stencil attachments, then the
  `rasterizationSamples` member of `pMultisampleState` **must** equal
  the maximum of the sample counts of those subpass attachments

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-06854"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-06854"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-06854  
  If `renderPass` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the
  [`VK_EXT_multisampled_render_to_single_sampled`](VK_EXT_multisampled_render_to_single_sampled.html)
  extension is enabled, `rasterizationSamples` is not dynamic, and
  `subpass` has a
  [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html)
  structure included in the
  [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html)::`pNext` chain
  with `multisampledRenderToSingleSampledEnable` equal to `VK_TRUE`,
  then the `rasterizationSamples` member of `pMultisampleState` **must**
  be equal to
  [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html)::`rasterizationSamples`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-subpass-01411"
  id="VUID-VkGraphicsPipelineCreateInfo-subpass-01411"></a>
  VUID-VkGraphicsPipelineCreateInfo-subpass-01411  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  the
  [`VK_NV_framebuffer_mixed_samples`](VK_NV_framebuffer_mixed_samples.html)
  extension is enabled, `rasterizationSamples` is not dynamic, and if
  `subpass` has a depth/stencil attachment and depth test, stencil test,
  or depth bounds test are enabled, then the `rasterizationSamples`
  member of `pMultisampleState` **must** be the same as the sample count
  of the depth/stencil attachment

- <a href="#VUID-VkGraphicsPipelineCreateInfo-subpass-01412"
  id="VUID-VkGraphicsPipelineCreateInfo-subpass-01412"></a>
  VUID-VkGraphicsPipelineCreateInfo-subpass-01412  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  the
  [`VK_NV_framebuffer_mixed_samples`](VK_NV_framebuffer_mixed_samples.html)
  extension is enabled, `rasterizationSamples` is not dynamic, and if
  `subpass` has any color attachments, then the `rasterizationSamples`
  member of `pMultisampleState` **must** be greater than or equal to the
  sample count for those subpass attachments

- <a href="#VUID-VkGraphicsPipelineCreateInfo-coverageReductionMode-02722"
  id="VUID-VkGraphicsPipelineCreateInfo-coverageReductionMode-02722"></a>
  VUID-VkGraphicsPipelineCreateInfo-coverageReductionMode-02722  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  the
  [`VK_NV_coverage_reduction_mode`](VK_NV_coverage_reduction_mode.html)
  extension is enabled, and `rasterizationSamples` is not dynamic, the
  coverage reduction mode specified by
  [VkPipelineCoverageReductionStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageReductionStateCreateInfoNV.html)::`coverageReductionMode`,
  the `rasterizationSamples` member of `pMultisampleState` and the
  sample counts for the color and depth/stencil attachments (if the
  subpass has them) **must** be a valid combination returned by
  `vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-subpass-00758"
  id="VUID-VkGraphicsPipelineCreateInfo-subpass-00758"></a>
  VUID-VkGraphicsPipelineCreateInfo-subpass-00758  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  `rasterizationSamples` is not dynamic, and `subpass` does not use any
  color and/or depth/stencil attachments, then the
  `rasterizationSamples` member of `pMultisampleState` **must** follow
  the rules for a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-noattachments"
  target="_blank" rel="noopener">zero-attachment subpass</a>

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-06046"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-06046"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-06046  
  If `renderPass` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `subpass` **must** be a valid subpass within `renderPass`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-06047"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-06047"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-06047  
  If `renderPass` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the
  pipeline is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>,
  `subpass` viewMask is not `0`, and `multiviewTessellationShader` is
  not enabled, then `pStages` **must** not include tessellation shaders

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-06048"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-06048"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-06048  
  If `renderPass` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the
  pipeline is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>,
  `subpass` viewMask is not `0`, and `multiviewGeometryShader` is not
  enabled, then `pStages` **must** not include a geometry shader

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-06050"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-06050"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-06050  
  If `renderPass` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and the
  pipeline is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  `subpass` viewMask is not `0`, then all of the shaders in the pipeline
  **must** not include variables decorated with the `Layer` built-in
  decoration in their interfaces

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-07064"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-07064"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-07064  
  If `renderPass` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the
  pipeline is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>,
  `subpass` viewMask is not `0`, and `multiviewMeshShader` is not
  enabled, then `pStages` **must** not include a mesh shader

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-00764"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-00764"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-00764  
  `flags` **must** not contain the `VK_PIPELINE_CREATE_DISPATCH_BASE`
  flag

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pStages-01565"
  id="VUID-VkGraphicsPipelineCreateInfo-pStages-01565"></a>
  VUID-VkGraphicsPipelineCreateInfo-pStages-01565  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> and an input
  attachment was referenced by an `aspectMask` at `renderPass` creation
  time, the fragment shader **must** only read from the aspects that
  were specified for that input attachment

- <a href="#VUID-VkGraphicsPipelineCreateInfo-layout-01688"
  id="VUID-VkGraphicsPipelineCreateInfo-layout-01688"></a>
  VUID-VkGraphicsPipelineCreateInfo-layout-01688  
  The number of resources in `layout` accessible to each shader stage
  that is used by the pipeline **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxPerStageResources`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-01715"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-01715"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-01715  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  no element of the `pDynamicStates` member of `pDynamicState` is
  `VK_DYNAMIC_STATE_VIEWPORT_W_SCALING_NV`, and the
  `viewportWScalingEnable` member of a
  [VkPipelineViewportWScalingStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportWScalingStateCreateInfoNV.html)
  structure, included in the `pNext` chain of `pViewportState`, is
  `VK_TRUE`, the `pViewportWScalings` member of the
  [VkPipelineViewportWScalingStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportWScalingStateCreateInfoNV.html)
  **must** be a pointer to an array of
  [VkPipelineViewportWScalingStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportWScalingStateCreateInfoNV.html)::`viewportCount`
  valid [VkViewportWScalingNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewportWScalingNV.html) structures

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04056"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04056"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04056  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  no element of the `pDynamicStates` member of `pDynamicState` is
  `VK_DYNAMIC_STATE_EXCLUSIVE_SCISSOR_NV`, and if
  `pViewportState->pNext` chain includes a
  [VkPipelineViewportExclusiveScissorStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportExclusiveScissorStateCreateInfoNV.html)
  structure, and if its `exclusiveScissorCount` member is not `0`, then
  its `pExclusiveScissors` member **must** be a valid pointer to an
  array of `exclusiveScissorCount` [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html) structures

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07854"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07854"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07854  
  If `VK_DYNAMIC_STATE_EXCLUSIVE_SCISSOR_ENABLE_NV` is included in the
  `pDynamicStates` array then the implementation **must** support at
  least `specVersion` `2` of the
  [`VK_NV_scissor_exclusive`](VK_NV_scissor_exclusive.html) extension

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04057"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04057"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04057  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  no element of the `pDynamicStates` member of `pDynamicState` is
  `VK_DYNAMIC_STATE_VIEWPORT_SHADING_RATE_PALETTE_NV`, and if
  `pViewportState->pNext` chain includes a
  [VkPipelineViewportShadingRateImageStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportShadingRateImageStateCreateInfoNV.html)
  structure, then its `pShadingRatePalettes` member **must** be a valid
  pointer to an array of `viewportCount` valid
  [VkShadingRatePaletteNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShadingRatePaletteNV.html) structures

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04058"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04058"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04058  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  no element of the `pDynamicStates` member of `pDynamicState` is
  `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_EXT`, and if `pNext` chain
  includes a
  [VkPipelineDiscardRectangleStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDiscardRectangleStateCreateInfoEXT.html)
  structure, and if its `discardRectangleCount` member is not `0`, then
  its `pDiscardRectangles` member **must** be a valid pointer to an
  array of `discardRectangleCount` [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html) structures

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07855"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07855"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07855  
  If `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_ENABLE_EXT` is included in the
  `pDynamicStates` array then the implementation **must** support at
  least `specVersion` `2` of the
  [`VK_EXT_discard_rectangles`](VK_EXT_discard_rectangles.html)
  extension

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07856"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07856"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07856  
  If `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_MODE_EXT` is included in the
  `pDynamicStates` array then the implementation **must** support at
  least `specVersion` `2` of the
  [`VK_EXT_discard_rectangles`](VK_EXT_discard_rectangles.html)
  extension

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pStages-02097"
  id="VUID-VkGraphicsPipelineCreateInfo-pStages-02097"></a>
  VUID-VkGraphicsPipelineCreateInfo-pStages-02097  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-vertex-input"
  target="_blank" rel="noopener">vertex input state</a>, and
  `pVertexInputState` is not dynamic, then `pVertexInputState` **must**
  be a valid pointer to a valid
  [VkPipelineVertexInputStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineVertexInputStateCreateInfo.html)
  structure

- <a href="#VUID-VkGraphicsPipelineCreateInfo-Input-07904"
  id="VUID-VkGraphicsPipelineCreateInfo-Input-07904"></a>
  VUID-VkGraphicsPipelineCreateInfo-Input-07904  
  If the pipeline is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-vertex-input"
  target="_blank" rel="noopener">vertex input state</a> and
  `pVertexInputState` is not dynamic, then all variables with the
  `Input` storage class decorated with `Location` in the `Vertex`
  `Execution` `Model` `OpEntryPoint` **must** contain a location in
  [VkVertexInputAttributeDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputAttributeDescription.html)::`location`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-Input-08733"
  id="VUID-VkGraphicsPipelineCreateInfo-Input-08733"></a>
  VUID-VkGraphicsPipelineCreateInfo-Input-08733  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-vertex-input"
  target="_blank" rel="noopener">vertex input state</a> and
  `pVertexInputState` is not dynamic, then the numeric type associated
  with all `Input` variables of the corresponding `Location` in the
  `Vertex` `Execution` `Model` `OpEntryPoint` **must** be the same as
  [VkVertexInputAttributeDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputAttributeDescription.html)::`format`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pVertexInputState-08929"
  id="VUID-VkGraphicsPipelineCreateInfo-pVertexInputState-08929"></a>
  VUID-VkGraphicsPipelineCreateInfo-pVertexInputState-08929  
  If the pipeline is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-vertex-input"
  target="_blank" rel="noopener">vertex input state</a> and
  `pVertexInputState` is not dynamic, and
  [VkVertexInputAttributeDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputAttributeDescription.html)::`format`
  has a 64-bit component, then the scalar width associated with all
  `Input` variables of the corresponding `Location` in the `Vertex`
  `Execution` `Model` `OpEntryPoint` **must** be 64-bit

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pVertexInputState-08930"
  id="VUID-VkGraphicsPipelineCreateInfo-pVertexInputState-08930"></a>
  VUID-VkGraphicsPipelineCreateInfo-pVertexInputState-08930  
  If the pipeline is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-vertex-input"
  target="_blank" rel="noopener">vertex input state</a> and
  `pVertexInputState` is not dynamic, and the scalar width associated
  with a `Location` decorated `Input` variable in the `Vertex`
  `Execution` `Model` `OpEntryPoint` is 64-bit, then the corresponding
  [VkVertexInputAttributeDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputAttributeDescription.html)::`format`
  **must** have a 64-bit component

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pVertexInputState-09198"
  id="VUID-VkGraphicsPipelineCreateInfo-pVertexInputState-09198"></a>
  VUID-VkGraphicsPipelineCreateInfo-pVertexInputState-09198  
  If the pipeline is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-vertex-input"
  target="_blank" rel="noopener">vertex input state</a> and
  `pVertexInputState` is not dynamic, and
  [VkVertexInputAttributeDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputAttributeDescription.html)::`format`
  has a 64-bit component, then all `Input` variables at the
  corresponding `Location` in the `Vertex` `Execution` `Model`
  `OpEntryPoint` **must** not use components that are not present in the
  format

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-dynamicPrimitiveTopologyUnrestricted-09031"
  id="VUID-VkGraphicsPipelineCreateInfo-dynamicPrimitiveTopologyUnrestricted-09031"></a>
  VUID-VkGraphicsPipelineCreateInfo-dynamicPrimitiveTopologyUnrestricted-09031  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-vertex-input"
  target="_blank" rel="noopener">vertex input state</a>, and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-pInputAssemblyState-null"
  target="_blank" rel="noopener">related dynamic state is not set</a>,
  `pInputAssemblyState` **must** be a valid pointer to a valid
  [VkPipelineInputAssemblyStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineInputAssemblyStateCreateInfo.html)
  structure

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pInputAssemblyState-09032"
  id="VUID-VkGraphicsPipelineCreateInfo-pInputAssemblyState-09032"></a>
  VUID-VkGraphicsPipelineCreateInfo-pInputAssemblyState-09032  
  If `pInputAssemblyState` is not `NULL` it **must** be a valid pointer
  to a valid
  [VkPipelineInputAssemblyStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineInputAssemblyStateCreateInfo.html)
  structure

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pStages-02317"
  id="VUID-VkGraphicsPipelineCreateInfo-pStages-02317"></a>
  VUID-VkGraphicsPipelineCreateInfo-pStages-02317  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, the
  `Xfb` execution mode **can** be specified by no more than one shader
  stage in `pStages`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pStages-02318"
  id="VUID-VkGraphicsPipelineCreateInfo-pStages-02318"></a>
  VUID-VkGraphicsPipelineCreateInfo-pStages-02318  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  any shader stage in `pStages` specifies `Xfb` execution mode it
  **must** be the last <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader stage</a>

- <a href="#VUID-VkGraphicsPipelineCreateInfo-rasterizationStream-02319"
  id="VUID-VkGraphicsPipelineCreateInfo-rasterizationStream-02319"></a>
  VUID-VkGraphicsPipelineCreateInfo-rasterizationStream-02319  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  a
  [VkPipelineRasterizationStateStreamCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateStreamCreateInfoEXT.html)::`rasterizationStream`
  value other than zero is specified, all variables in the output
  interface of the entry point being compiled decorated with `Position`,
  `PointSize`, `ClipDistance`, or `CullDistance` **must** be decorated
  with identical `Stream` values that match the `rasterizationStream`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-rasterizationStream-02320"
  id="VUID-VkGraphicsPipelineCreateInfo-rasterizationStream-02320"></a>
  VUID-VkGraphicsPipelineCreateInfo-rasterizationStream-02320  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  [VkPipelineRasterizationStateStreamCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateStreamCreateInfoEXT.html)::`rasterizationStream`
  is zero, or not specified, all variables in the output interface of
  the entry point being compiled decorated with `Position`, `PointSize`,
  `ClipDistance`, or `CullDistance` **must** be decorated with a
  `Stream` value of zero, or **must** not specify the `Stream`
  decoration

- <a href="#VUID-VkGraphicsPipelineCreateInfo-geometryStreams-02321"
  id="VUID-VkGraphicsPipelineCreateInfo-geometryStreams-02321"></a>
  VUID-VkGraphicsPipelineCreateInfo-geometryStreams-02321  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  the last <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader stage</a> is a
  geometry shader, and that geometry shader uses the `GeometryStreams`
  capability, then
  `VkPhysicalDeviceTransformFeedbackFeaturesEXT`::`geometryStreams`
  feature **must** be enabled

- <a href="#VUID-VkGraphicsPipelineCreateInfo-None-02322"
  id="VUID-VkGraphicsPipelineCreateInfo-None-02322"></a>
  VUID-VkGraphicsPipelineCreateInfo-None-02322  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  there are any mesh shader stages in the pipeline there **must** not be
  any shader stage in the pipeline with a `Xfb` execution mode

- <a href="#VUID-VkGraphicsPipelineCreateInfo-lineRasterizationMode-02766"
  id="VUID-VkGraphicsPipelineCreateInfo-lineRasterizationMode-02766"></a>
  VUID-VkGraphicsPipelineCreateInfo-lineRasterizationMode-02766  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> and
  at least one of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a> or
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>, and
  `pMultisampleState` is not `NULL`, the `lineRasterizationMode` member
  of a
  [VkPipelineRasterizationLineStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationLineStateCreateInfoKHR.html)
  structure included in the `pNext` chain of `pRasterizationState` is
  `VK_LINE_RASTERIZATION_MODE_BRESENHAM_KHR` or
  `VK_LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH_KHR`, then the
  `alphaToCoverageEnable`, `alphaToOneEnable`, and `sampleShadingEnable`
  members of `pMultisampleState` **must** all be `VK_FALSE`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-stippledLineEnable-02767"
  id="VUID-VkGraphicsPipelineCreateInfo-stippledLineEnable-02767"></a>
  VUID-VkGraphicsPipelineCreateInfo-stippledLineEnable-02767  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, the
  `stippledLineEnable` member of
  [VkPipelineRasterizationLineStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationLineStateCreateInfoKHR.html)
  is `VK_TRUE`, and no element of the `pDynamicStates` member of
  `pDynamicState` is `VK_DYNAMIC_STATE_LINE_STIPPLE_EXT`, then the
  `lineStippleFactor` member of
  [VkPipelineRasterizationLineStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationLineStateCreateInfoKHR.html)
  **must** be in the range \[1,256\]

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-03372"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-03372"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-03372  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_ANY_HIT_SHADERS_BIT_KHR`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-03373"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-03373"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-03373  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_CLOSEST_HIT_SHADERS_BIT_KHR`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-03374"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-03374"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-03374  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_MISS_SHADERS_BIT_KHR`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-03375"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-03375"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-03375  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_NO_NULL_INTERSECTION_SHADERS_BIT_KHR`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-03376"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-03376"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-03376  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-03377"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-03377"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-03377  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_AABBS_BIT_KHR`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-03577"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-03577"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-03577  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-04947"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-04947"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-04947  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_ALLOW_MOTION_BIT_NV`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-03378"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-03378"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-03378  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState"
  target="_blank" rel="noopener"><code>extendedDynamicState</code></a>
  feature is not enabled, and the value of
  [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkApplicationInfo.html)::`apiVersion` used to
  create the [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) is less than Version 1.3
  there **must** be no element of the `pDynamicStates` member of
  `pDynamicState` set to `VK_DYNAMIC_STATE_CULL_MODE`,
  `VK_DYNAMIC_STATE_FRONT_FACE`, `VK_DYNAMIC_STATE_PRIMITIVE_TOPOLOGY`,
  `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT`,
  `VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT`,
  `VK_DYNAMIC_STATE_VERTEX_INPUT_BINDING_STRIDE`,
  `VK_DYNAMIC_STATE_DEPTH_TEST_ENABLE`,
  `VK_DYNAMIC_STATE_DEPTH_WRITE_ENABLE`,
  `VK_DYNAMIC_STATE_DEPTH_COMPARE_OP`,
  `VK_DYNAMIC_STATE_DEPTH_BOUNDS_TEST_ENABLE`,
  `VK_DYNAMIC_STATE_STENCIL_TEST_ENABLE`, or
  `VK_DYNAMIC_STATE_STENCIL_OP`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-03379"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-03379"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-03379  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT` is included in the
  `pDynamicStates` array then `viewportCount` **must** be zero

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-03380"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-03380"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-03380  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  `VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT` is included in the
  `pDynamicStates` array then `scissorCount` **must** be zero

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04132"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04132"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04132  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT` is included in the
  `pDynamicStates` array then `VK_DYNAMIC_STATE_VIEWPORT` **must** not
  be present

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04133"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04133"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04133  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  `VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT` is included in the
  `pDynamicStates` array then `VK_DYNAMIC_STATE_SCISSOR` **must** not be
  present

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07065"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07065"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07065  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  includes a mesh shader, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_PRIMITIVE_TOPOLOGY`, or
  `VK_DYNAMIC_STATE_VERTEX_INPUT_BINDING_STRIDE`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04868"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04868"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04868  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState2"
  target="_blank" rel="noopener"><code>extendedDynamicState2</code></a>
  feature is not enabled, and the value of
  [VkApplicationInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkApplicationInfo.html)::`apiVersion` used to
  create the [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) is less than Version 1.3
  there **must** be no element of the `pDynamicStates` member of
  `pDynamicState` set to `VK_DYNAMIC_STATE_DEPTH_BIAS_ENABLE`,
  `VK_DYNAMIC_STATE_PRIMITIVE_RESTART_ENABLE`, or
  `VK_DYNAMIC_STATE_RASTERIZER_DISCARD_ENABLE`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04869"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04869"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04869  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState2LogicOp"
  target="_blank"
  rel="noopener"><code>extendedDynamicState2LogicOp</code></a> feature
  is not enabled, there **must** be no element of the `pDynamicStates`
  member of `pDynamicState` set to `VK_DYNAMIC_STATE_LOGIC_OP_EXT`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04870"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04870"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04870  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState2PatchControlPoints"
  target="_blank"
  rel="noopener"><code>extendedDynamicState2PatchControlPoints</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_PATCH_CONTROL_POINTS_EXT`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07066"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07066"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07066  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  includes a mesh shader, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_PRIMITIVE_RESTART_ENABLE`, or
  `VK_DYNAMIC_STATE_PATCH_CONTROL_POINTS_EXT`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-02877"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-02877"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-02877  
  If `flags` includes `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV`,
  then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-deviceGeneratedCommands"
  target="_blank" rel="noopener"><code>deviceGeneratedCommands</code></a>
  feature **must** be enabled

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-02966"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-02966"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-02966  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> and
  `flags` includes `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV`, then
  all stages **must** not specify `Xfb` execution mode

- <a href="#VUID-VkGraphicsPipelineCreateInfo-libraryCount-06648"
  id="VUID-VkGraphicsPipelineCreateInfo-libraryCount-06648"></a>
  VUID-VkGraphicsPipelineCreateInfo-libraryCount-06648  
  If the pipeline is not created with a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-complete"
  target="_blank" rel="noopener">complete set of state</a>, or
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`libraryCount`
  is not `0`,
  [VkGraphicsPipelineShaderGroupsCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineShaderGroupsCreateInfoNV.html)::`groupCount`
  and
  [VkGraphicsPipelineShaderGroupsCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineShaderGroupsCreateInfoNV.html)::`pipelineCount`
  **must** be `0`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-libraryCount-06649"
  id="VUID-VkGraphicsPipelineCreateInfo-libraryCount-06649"></a>
  VUID-VkGraphicsPipelineCreateInfo-libraryCount-06649  
  If the pipeline is created with a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-complete"
  target="_blank" rel="noopener">complete set of state</a>, and
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`libraryCount`
  is `0`, and the `pNext` chain includes an instance of
  [VkGraphicsPipelineShaderGroupsCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineShaderGroupsCreateInfoNV.html),
  [VkGraphicsPipelineShaderGroupsCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineShaderGroupsCreateInfoNV.html)::`groupCount`
  **must** be greater than `0`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-pipelineCreationCacheControl-02878"
  id="VUID-VkGraphicsPipelineCreateInfo-pipelineCreationCacheControl-02878"></a>
  VUID-VkGraphicsPipelineCreateInfo-pipelineCreationCacheControl-02878  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelineCreationCacheControl"
  target="_blank"
  rel="noopener"><code>pipelineCreationCacheControl</code></a> feature
  is not enabled, `flags` **must** not include
  `VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT` or
  `VK_PIPELINE_CREATE_EARLY_RETURN_ON_FAILURE_BIT`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-pipelineProtectedAccess-07368"
  id="VUID-VkGraphicsPipelineCreateInfo-pipelineProtectedAccess-07368"></a>
  VUID-VkGraphicsPipelineCreateInfo-pipelineProtectedAccess-07368  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelineProtectedAccess"
  target="_blank" rel="noopener"><code>pipelineProtectedAccess</code></a>
  feature is not enabled, `flags` **must** not include
  `VK_PIPELINE_CREATE_NO_PROTECTED_ACCESS_BIT_EXT` or
  `VK_PIPELINE_CREATE_PROTECTED_ACCESS_ONLY_BIT_EXT`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-07369"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-07369"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-07369  
  `flags` **must** not include both
  `VK_PIPELINE_CREATE_NO_PROTECTED_ACCESS_BIT_EXT` and
  `VK_PIPELINE_CREATE_PROTECTED_ACCESS_ONLY_BIT_EXT`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04494"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04494"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04494  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> or
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> and
  `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR` is not included in
  `pDynamicState->pDynamicStates`,
  [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html)::`fragmentSize.width`
  **must** be greater than or equal to `1`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04495"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04495"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04495  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> or
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> and
  `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR` is not included in
  `pDynamicState->pDynamicStates`,
  [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html)::`fragmentSize.height`
  **must** be greater than or equal to `1`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04496"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04496"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04496  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> or
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> and
  `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR` is not included in
  `pDynamicState->pDynamicStates`,
  [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html)::`fragmentSize.width`
  **must** be a power-of-two value

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04497"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04497"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04497  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> or
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> and
  `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR` is not included in
  `pDynamicState->pDynamicStates`,
  [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html)::`fragmentSize.height`
  **must** be a power-of-two value

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04498"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04498"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04498  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> or
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> and
  `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR` is not included in
  `pDynamicState->pDynamicStates`,
  [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html)::`fragmentSize.width`
  **must** be less than or equal to `4`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04499"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04499"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04499  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> or
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> and
  `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR` is not included in
  `pDynamicState->pDynamicStates`,
  [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html)::`fragmentSize.height`
  **must** be less than or equal to `4`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04500"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04500"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04500  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> or
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> and
  `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR` is not included in
  `pDynamicState->pDynamicStates`, and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelineFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>pipelineFragmentShadingRate</code></a> feature is
  not enabled,
  [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html)::`fragmentSize.width`
  and
  [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html)::`fragmentSize.height`
  **must** both be equal to `1`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicState-06567"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicState-06567"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicState-06567  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> or
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> and
  `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR` is not included in
  `pDynamicState->pDynamicStates`,
  [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html)::`combinerOps`\[0\]
  **must** be a valid
  [VkFragmentShadingRateCombinerOpKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFragmentShadingRateCombinerOpKHR.html)
  value

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicState-06568"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicState-06568"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicState-06568  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> or
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> and
  `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR` is not included in
  `pDynamicState->pDynamicStates`,
  [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html)::`combinerOps`\[1\]
  **must** be a valid
  [VkFragmentShadingRateCombinerOpKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFragmentShadingRateCombinerOpKHR.html)
  value

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04501"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04501"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04501  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> or
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> and
  `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR` is not included in
  `pDynamicState->pDynamicStates`, and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-primitiveFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>primitiveFragmentShadingRate</code></a> feature
  is not enabled,
  [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html)::`combinerOps`\[0\]
  **must** be `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_KEEP_KHR`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04502"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04502"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04502  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> or
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> and
  `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR` is not included in
  `pDynamicState->pDynamicStates`, and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-attachmentFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>attachmentFragmentShadingRate</code></a> feature
  is not enabled,
  [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html)::`combinerOps`\[1\]
  **must** be `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_KEEP_KHR`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-primitiveFragmentShadingRateWithMultipleViewports-04503"
  id="VUID-VkGraphicsPipelineCreateInfo-primitiveFragmentShadingRateWithMultipleViewports-04503"></a>
  VUID-VkGraphicsPipelineCreateInfo-primitiveFragmentShadingRateWithMultipleViewports-04503  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> and
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-primitiveFragmentShadingRateWithMultipleViewports"
  target="_blank"
  rel="noopener"><code>primitiveFragmentShadingRateWithMultipleViewports</code></a>
  limit is not supported, `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT` is not
  included in `pDynamicState->pDynamicStates`, and
  [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html)::`viewportCount`
  is greater than `1`, entry points specified in `pStages` **must** not
  write to the `PrimitiveShadingRateKHR` built-in

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-primitiveFragmentShadingRateWithMultipleViewports-04504"
  id="VUID-VkGraphicsPipelineCreateInfo-primitiveFragmentShadingRateWithMultipleViewports-04504"></a>
  VUID-VkGraphicsPipelineCreateInfo-primitiveFragmentShadingRateWithMultipleViewports-04504  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> and
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-primitiveFragmentShadingRateWithMultipleViewports"
  target="_blank"
  rel="noopener"><code>primitiveFragmentShadingRateWithMultipleViewports</code></a>
  limit is not supported, and entry points specified in `pStages` write
  to the `ViewportIndex` built-in, they **must** not also write to the
  `PrimitiveShadingRateKHR` built-in

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-primitiveFragmentShadingRateWithMultipleViewports-04505"
  id="VUID-VkGraphicsPipelineCreateInfo-primitiveFragmentShadingRateWithMultipleViewports-04505"></a>
  VUID-VkGraphicsPipelineCreateInfo-primitiveFragmentShadingRateWithMultipleViewports-04505  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> and
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-primitiveFragmentShadingRateWithMultipleViewports"
  target="_blank"
  rel="noopener"><code>primitiveFragmentShadingRateWithMultipleViewports</code></a>
  limit is not supported, and entry points specified in `pStages` write
  to the `ViewportMaskNV` built-in, they **must** not also write to the
  `PrimitiveShadingRateKHR` built-in

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-fragmentShadingRateNonTrivialCombinerOps-04506"
  id="VUID-VkGraphicsPipelineCreateInfo-fragmentShadingRateNonTrivialCombinerOps-04506"></a>
  VUID-VkGraphicsPipelineCreateInfo-fragmentShadingRateNonTrivialCombinerOps-04506  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> or
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-fragmentShadingRateNonTrivialCombinerOps"
  target="_blank"
  rel="noopener"><code>fragmentShadingRateNonTrivialCombinerOps</code></a>
  limit is not supported, and
  `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR` is not included in
  `pDynamicState->pDynamicStates`, elements of
  [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html)::`combinerOps`
  **must** be `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_KEEP_KHR` or
  `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_REPLACE_KHR`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-None-06569"
  id="VUID-VkGraphicsPipelineCreateInfo-None-06569"></a>
  VUID-VkGraphicsPipelineCreateInfo-None-06569  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> and
  `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR` is not included in
  `pDynamicState->pDynamicStates`,
  [VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html)::`shadingRateType`
  **must** be a valid
  [VkFragmentShadingRateTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFragmentShadingRateTypeNV.html) value

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicState-06570"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicState-06570"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicState-06570  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> and
  `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR` is not included in
  `pDynamicState->pDynamicStates`,
  [VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html)::`shadingRate`
  **must** be a valid
  [VkFragmentShadingRateNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFragmentShadingRateNV.html) value

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicState-06571"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicState-06571"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicState-06571  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> and
  `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR` is not included in
  `pDynamicState->pDynamicStates`,
  [VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html)::`combinerOps`\[0\]
  **must** be a valid
  [VkFragmentShadingRateCombinerOpKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFragmentShadingRateCombinerOpKHR.html)
  value

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicState-06572"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicState-06572"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicState-06572  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> and
  `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR` is not included in
  `pDynamicState->pDynamicStates`,
  [VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html)::`combinerOps`\[1\]
  **must** be a valid
  [VkFragmentShadingRateCombinerOpKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFragmentShadingRateCombinerOpKHR.html)
  value

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04569"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04569"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04569  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> and
  `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR` is not included in
  `pDynamicState->pDynamicStates`, and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-fragmentShadingRateEnums"
  target="_blank" rel="noopener"><code>fragmentShadingRateEnums</code></a>
  feature is not enabled,
  [VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html)::`shadingRateType`
  **must** be equal to `VK_FRAGMENT_SHADING_RATE_TYPE_FRAGMENT_SIZE_NV`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04570"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04570"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04570  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> and
  `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR` is not included in
  `pDynamicState->pDynamicStates`, and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelineFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>pipelineFragmentShadingRate</code></a> feature is
  not enabled,
  [VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html)::`shadingRate`
  **must** be equal to
  `VK_FRAGMENT_SHADING_RATE_1_INVOCATION_PER_PIXEL_NV`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04571"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04571"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04571  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> and
  `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR` is not included in
  `pDynamicState->pDynamicStates`, and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-primitiveFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>primitiveFragmentShadingRate</code></a> feature
  is not enabled,
  [VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html)::`combinerOps`\[0\]
  **must** be `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_KEEP_KHR`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04572"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04572"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicState-04572  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> and
  `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR` is not included in
  `pDynamicState->pDynamicStates`, and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-attachmentFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>attachmentFragmentShadingRate</code></a> feature
  is not enabled,
  [VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html)::`combinerOps`\[1\]
  **must** be `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_KEEP_KHR`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-fragmentShadingRateNonTrivialCombinerOps-04573"
  id="VUID-VkGraphicsPipelineCreateInfo-fragmentShadingRateNonTrivialCombinerOps-04573"></a>
  VUID-VkGraphicsPipelineCreateInfo-fragmentShadingRateNonTrivialCombinerOps-04573  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>, and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-fragmentShadingRateNonTrivialCombinerOps"
  target="_blank"
  rel="noopener"><code>fragmentShadingRateNonTrivialCombinerOps</code></a>
  limit is not supported and
  `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR` is not included in
  `pDynamicState->pDynamicStates`, elements of
  [VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html)::`combinerOps`
  **must** be `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_KEEP_KHR` or
  `VK_FRAGMENT_SHADING_RATE_COMBINER_OP_REPLACE_KHR`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-None-04574"
  id="VUID-VkGraphicsPipelineCreateInfo-None-04574"></a>
  VUID-VkGraphicsPipelineCreateInfo-None-04574  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>, and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-supersampleFragmentShadingRates"
  target="_blank"
  rel="noopener"><code>supersampleFragmentShadingRates</code></a>
  feature is not enabled,
  [VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html)::`shadingRate`
  **must** not be equal to
  `VK_FRAGMENT_SHADING_RATE_2_INVOCATIONS_PER_PIXEL_NV`,
  `VK_FRAGMENT_SHADING_RATE_4_INVOCATIONS_PER_PIXEL_NV`,
  `VK_FRAGMENT_SHADING_RATE_8_INVOCATIONS_PER_PIXEL_NV`, or
  `VK_FRAGMENT_SHADING_RATE_16_INVOCATIONS_PER_PIXEL_NV`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-None-04575"
  id="VUID-VkGraphicsPipelineCreateInfo-None-04575"></a>
  VUID-VkGraphicsPipelineCreateInfo-None-04575  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>, and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-noInvocationFragmentShadingRates"
  target="_blank"
  rel="noopener"><code>noInvocationFragmentShadingRates</code></a>
  feature is not enabled,
  [VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html)::`shadingRate`
  **must** not be equal to `VK_FRAGMENT_SHADING_RATE_NO_INVOCATIONS_NV`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-03578"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-03578"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-03578  
  All elements of the `pDynamicStates` member of `pDynamicState`
  **must** not be `VK_DYNAMIC_STATE_RAY_TRACING_PIPELINE_STACK_SIZE_KHR`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04807"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04807"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04807  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> and
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-vertexInputDynamicState"
  target="_blank" rel="noopener"><code>vertexInputDynamicState</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_VERTEX_INPUT_EXT`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07067"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07067"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07067  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  includes a mesh shader, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_VERTEX_INPUT_EXT`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04800"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04800"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-04800  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-colorWriteEnable"
  target="_blank" rel="noopener"><code>colorWriteEnable</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_COLOR_WRITE_ENABLE_EXT`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-rasterizationSamples-04899"
  id="VUID-VkGraphicsPipelineCreateInfo-rasterizationSamples-04899"></a>
  VUID-VkGraphicsPipelineCreateInfo-rasterizationSamples-04899  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>, and the
  [`VK_QCOM_render_pass_shader_resolve`](VK_QCOM_render_pass_shader_resolve.html)
  extension is enabled, `rasterizationSamples` is not dynamic, and if
  subpass has any input attachments, and if the subpass description
  contains `VK_SUBPASS_DESCRIPTION_FRAGMENT_REGION_BIT_QCOM`, then the
  sample count of the input attachments **must** equal
  `rasterizationSamples`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-sampleShadingEnable-04900"
  id="VUID-VkGraphicsPipelineCreateInfo-sampleShadingEnable-04900"></a>
  VUID-VkGraphicsPipelineCreateInfo-sampleShadingEnable-04900  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>, and the
  [`VK_QCOM_render_pass_shader_resolve`](VK_QCOM_render_pass_shader_resolve.html)
  extension is enabled, and if the subpass description contains
  `VK_SUBPASS_DESCRIPTION_FRAGMENT_REGION_BIT_QCOM`, then
  `sampleShadingEnable` **must** be false

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-04901"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-04901"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-04901  
  If `flags` includes `VK_SUBPASS_DESCRIPTION_SHADER_RESOLVE_BIT_QCOM`,
  then the subpass **must** be the last subpass in a subpass dependency
  chain

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-04902"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-04902"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-04902  
  If `flags` includes `VK_SUBPASS_DESCRIPTION_SHADER_RESOLVE_BIT_QCOM`,
  and if `pResolveAttachments` is not `NULL`, then each resolve
  attachment **must** be `VK_ATTACHMENT_UNUSED`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-dynamicRendering-06576"
  id="VUID-VkGraphicsPipelineCreateInfo-dynamicRendering-06576"></a>
  VUID-VkGraphicsPipelineCreateInfo-dynamicRendering-06576  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-dynamicRendering"
  target="_blank" rel="noopener"><code>dynamicRendering</code></a>
  feature is not enabled and the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  `renderPass` **must** not be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkGraphicsPipelineCreateInfo-multiview-06577"
  id="VUID-VkGraphicsPipelineCreateInfo-multiview-06577"></a>
  VUID-VkGraphicsPipelineCreateInfo-multiview-06577  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiview"
  target="_blank" rel="noopener"><code>multiview</code></a> feature is
  not enabled, the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  and `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`viewMask`
  **must** be `0`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-06578"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-06578"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-06578  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  and `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the index
  of the most significant bit in
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`viewMask`
  **must** be less than <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxMultiviewViewCount"
  target="_blank" rel="noopener"><code>maxMultiviewViewCount</code></a>

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-06579"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-06579"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-06579  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  and `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`colorAttachmentCount`
  is not 0,
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`pColorAttachmentFormats`
  **must** be a valid pointer to an array of `colorAttachmentCount`
  valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) values

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-06580"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-06580"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-06580  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  and `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), each
  element of
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`pColorAttachmentFormats`
  **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-06582"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-06582"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-06582  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and any element
  of
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`pColorAttachmentFormats`
  is not `VK_FORMAT_UNDEFINED`, that format **must** be a format with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#potential-format-features"
  target="_blank" rel="noopener">potential format features</a> that
  include `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT` or
  `VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-06583"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-06583"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-06583  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  and `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`depthAttachmentFormat`
  **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-06584"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-06584"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-06584  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  and `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`stencilAttachmentFormat`
  **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-06585"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-06585"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-06585  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`depthAttachmentFormat`
  is not `VK_FORMAT_UNDEFINED`, it **must** be a format with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#potential-format-features"
  target="_blank" rel="noopener">potential format features</a> that
  include `VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-06586"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-06586"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-06586  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`stencilAttachmentFormat`
  is not `VK_FORMAT_UNDEFINED`, it **must** be a format with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#potential-format-features"
  target="_blank" rel="noopener">potential format features</a> that
  include `VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-06587"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-06587"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-06587  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`depthAttachmentFormat`
  is not `VK_FORMAT_UNDEFINED`, it **must** be a format that includes a
  depth component

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-06588"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-06588"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-06588  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`stencilAttachmentFormat`
  is not `VK_FORMAT_UNDEFINED`, it **must** be a format that includes a
  stencil component

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-06589"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-06589"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-06589  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`depthAttachmentFormat`
  is not `VK_FORMAT_UNDEFINED`, and
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`stencilAttachmentFormat`
  is not `VK_FORMAT_UNDEFINED`, `depthAttachmentFormat` **must** equal
  `stencilAttachmentFormat`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-09033"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-09033"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-09033  
  If `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the pipeline
  is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  and either of
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`depthAttachmentFormat`
  or
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`stencilAttachmentFormat`
  are not `VK_FORMAT_UNDEFINED`, and the
  [`VK_EXT_extended_dynamic_state3`](VK_EXT_extended_dynamic_state3.html)
  extension is not enabled or any of the
  `VK_DYNAMIC_STATE_DEPTH_TEST_ENABLE`,
  `VK_DYNAMIC_STATE_DEPTH_WRITE_ENABLE`,
  `VK_DYNAMIC_STATE_DEPTH_COMPARE_OP`,
  `VK_DYNAMIC_STATE_DEPTH_BOUNDS_TEST_ENABLE`,
  `VK_DYNAMIC_STATE_STENCIL_TEST_ENABLE`, `VK_DYNAMIC_STATE_STENCIL_OP`,
  or `VK_DYNAMIC_STATE_DEPTH_BOUNDS` dynamic states are not set,
  `pDepthStencilState` **must** be a valid pointer to a valid
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
  structure

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDepthStencilState-09034"
  id="VUID-VkGraphicsPipelineCreateInfo-pDepthStencilState-09034"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDepthStencilState-09034  
  If `pDepthStencilState` is not `NULL` it **must** be a valid pointer
  to a valid
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
  structure

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-09035"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-09035"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-09035  
  If `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and the
  pipeline is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> but not <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  and the
  [`VK_EXT_extended_dynamic_state3`](VK_EXT_extended_dynamic_state3.html)
  extension is not enabled, or any of the
  `VK_DYNAMIC_STATE_DEPTH_TEST_ENABLE`,
  `VK_DYNAMIC_STATE_DEPTH_WRITE_ENABLE`,
  `VK_DYNAMIC_STATE_DEPTH_COMPARE_OP`,
  `VK_DYNAMIC_STATE_DEPTH_BOUNDS_TEST_ENABLE`,
  `VK_DYNAMIC_STATE_STENCIL_TEST_ENABLE`, `VK_DYNAMIC_STATE_STENCIL_OP`,
  or `VK_DYNAMIC_STATE_DEPTH_BOUNDS` dynamic states are not set,
  `pDepthStencilState` **must** be a valid pointer to a valid
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
  structure

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDepthStencilState-09036"
  id="VUID-VkGraphicsPipelineCreateInfo-pDepthStencilState-09036"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDepthStencilState-09036  
  If `pDepthStencilState` is not `NULL` it **must** be a valid pointer
  to a valid
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
  structure

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-09037"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-09037"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-09037  
  If `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the pipeline
  is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  and any element of
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`pColorAttachmentFormats`
  is not `VK_FORMAT_UNDEFINED`, and the
  [`VK_EXT_extended_dynamic_state3`](VK_EXT_extended_dynamic_state3.html)
  extension is not enabled, or any of the
  `VK_DYNAMIC_STATE_LOGIC_OP_ENABLE_EXT`,
  `VK_DYNAMIC_STATE_LOGIC_OP_EXT`,
  `VK_DYNAMIC_STATE_COLOR_BLEND_ENABLE_EXT`,
  `VK_DYNAMIC_STATE_COLOR_BLEND_EQUATION_EXT`,
  `VK_DYNAMIC_STATE_COLOR_WRITE_MASK_EXT`, or
  `VK_DYNAMIC_STATE_BLEND_CONSTANTS` dynamic states are not set,
  `pColorBlendState` **must** be a valid pointer to a valid
  [VkPipelineColorBlendStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateInfo.html)
  structure

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pColorBlendState-09038"
  id="VUID-VkGraphicsPipelineCreateInfo-pColorBlendState-09038"></a>
  VUID-VkGraphicsPipelineCreateInfo-pColorBlendState-09038  
  If `pColorBlendState` is not `NULL` it **must** be a valid pointer to
  a valid
  [VkPipelineColorBlendStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateInfo.html)
  structure

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-06055"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-06055"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-06055  
  If `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `pColorBlendState` is not dynamic, and the pipeline is being created
  with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  `pColorBlendState->attachmentCount` **must** be equal to
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`colorAttachmentCount`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-06057"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-06057"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-06057  
  If `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the pipeline
  is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>,
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`viewMask`
  is not `0`, and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiview-tess"
  target="_blank"
  rel="noopener"><code>multiviewTessellationShader</code></a> feature is
  not enabled, then `pStages` **must** not include tessellation shaders

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-06058"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-06058"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-06058  
  If `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the pipeline
  is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>,
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`viewMask`
  is not `0`, and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiview-gs"
  target="_blank" rel="noopener"><code>multiviewGeometryShader</code></a>
  feature is not enabled, then `pStages` **must** not include a geometry
  shader

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-06059"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-06059"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-06059  
  If `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the pipeline
  is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`viewMask`
  is not `0`, all of the shaders in the pipeline **must** not include
  variables decorated with the `Layer` built-in decoration in their
  interfaces

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-07720"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-07720"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-07720  
  If `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the pipeline
  is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`viewMask`
  is not `0`, and `multiviewMeshShader` is not enabled, then `pStages`
  **must** not include a mesh shader

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-06061"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-06061"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-06061  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-dynamicRenderingLocalRead"
  target="_blank"
  rel="noopener"><code>dynamicRenderingLocalRead</code></a> feature is
  not enabled, the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>, and
  `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), fragment
  shaders in `pStages` **must** not include the `InputAttachment`
  capability

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-08710"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-08710"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-08710  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> and
  `renderPass` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), fragment
  shaders in `pStages` **must** not include any of the
  `TileImageColorReadAccessEXT`, `TileImageDepthReadAccessEXT`, or
  `TileImageStencilReadAccessEXT` capabilities

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-06062"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-06062"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-06062  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a> and
  `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), for each color
  attachment format defined by the `pColorAttachmentFormats` member of
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html),
  if its <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#potential-format-features"
  target="_blank" rel="noopener">potential format features</a> do not
  contain `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BLEND_BIT`, then the
  `blendEnable` member of the corresponding element of the
  `pAttachments` member of `pColorBlendState` **must** be `VK_FALSE`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-06063"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-06063"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-06063  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a> and
  `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), if the `pNext`
  chain includes
  [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoAMD.html)
  or `VkAttachmentSampleCountInfoNV`, the `colorAttachmentCount` member
  of that structure **must** be equal to the value of
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`colorAttachmentCount`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06591"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06591"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06591  
  If `pStages` includes a fragment shader stage, and the fragment shader
  declares the `EarlyFragmentTests` execution mode, the `flags` member
  of
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
  **must** not include
  `VK_PIPELINE_DEPTH_STENCIL_STATE_CREATE_RASTERIZATION_ORDER_ATTACHMENT_DEPTH_ACCESS_BIT_EXT`
  or
  `VK_PIPELINE_DEPTH_STENCIL_STATE_CREATE_RASTERIZATION_ORDER_ATTACHMENT_STENCIL_ACCESS_BIT_EXT`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06482"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06482"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06482  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-dynamicRenderingLocalRead"
  target="_blank"
  rel="noopener"><code>dynamicRenderingLocalRead</code></a> feature is
  not enabled, the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  and the `flags` member of
  [VkPipelineColorBlendStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateInfo.html)
  includes
  `VK_PIPELINE_COLOR_BLEND_STATE_CREATE_RASTERIZATION_ORDER_ATTACHMENT_ACCESS_BIT_EXT`,
  `renderPass` **must** not be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkGraphicsPipelineCreateInfo-None-09526"
  id="VUID-VkGraphicsPipelineCreateInfo-None-09526"></a>
  VUID-VkGraphicsPipelineCreateInfo-None-09526  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-dynamicRenderingLocalRead"
  target="_blank"
  rel="noopener"><code>dynamicRenderingLocalRead</code></a> feature is
  not enabled, the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  and the `flags` member of
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
  includes
  `VK_PIPELINE_DEPTH_STENCIL_STATE_CREATE_RASTERIZATION_ORDER_ATTACHMENT_DEPTH_ACCESS_BIT_EXT`
  or
  `VK_PIPELINE_DEPTH_STENCIL_STATE_CREATE_RASTERIZATION_ORDER_ATTACHMENT_STENCIL_ACCESS_BIT_EXT`,
  `renderPass` **must** not be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-pColorAttachmentSamples-06592"
  id="VUID-VkGraphicsPipelineCreateInfo-pColorAttachmentSamples-06592"></a>
  VUID-VkGraphicsPipelineCreateInfo-pColorAttachmentSamples-06592  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  elements of the `pColorAttachmentSamples` member of
  [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoAMD.html)
  or [VkAttachmentSampleCountInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoNV.html)
  **must** be valid [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html)
  values

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-depthStencilAttachmentSamples-06593"
  id="VUID-VkGraphicsPipelineCreateInfo-depthStencilAttachmentSamples-06593"></a>
  VUID-VkGraphicsPipelineCreateInfo-depthStencilAttachmentSamples-06593  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a> and
  the `depthStencilAttachmentSamples` member of
  [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoAMD.html)
  or [VkAttachmentSampleCountInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoNV.html)
  is not 0, it **must** be a valid
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) value

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-09527"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-09527"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-09527  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  `renderPass` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and the
  `flags` member of
  [VkPipelineColorBlendStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateInfo.html)
  includes
  `VK_PIPELINE_COLOR_BLEND_STATE_CREATE_RASTERIZATION_ORDER_ATTACHMENT_ACCESS_BIT_EXT`
  `subpass` **must** have been created with
  `VK_SUBPASS_DESCRIPTION_RASTERIZATION_ORDER_ATTACHMENT_COLOR_ACCESS_BIT_EXT`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-09528"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-09528"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-09528  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>, `renderPass`
  is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and the `flags` member
  of
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
  includes
  `VK_PIPELINE_DEPTH_STENCIL_STATE_CREATE_RASTERIZATION_ORDER_ATTACHMENT_DEPTH_ACCESS_BIT_EXT`,
  `subpass` **must** have been created with
  `VK_SUBPASS_DESCRIPTION_RASTERIZATION_ORDER_ATTACHMENT_DEPTH_ACCESS_BIT_EXT`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-09529"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-09529"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-09529  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>, `renderPass`
  is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and the `flags` member
  of
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
  includes
  `VK_PIPELINE_DEPTH_STENCIL_STATE_CREATE_RASTERIZATION_ORDER_ATTACHMENT_STENCIL_ACCESS_BIT_EXT`,
  `subpass` **must** have been created with
  `VK_SUBPASS_DESCRIPTION_RASTERIZATION_ORDER_ATTACHMENT_STENCIL_ACCESS_BIT_EXT`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-pipelineStageCreationFeedbackCount-06594"
  id="VUID-VkGraphicsPipelineCreateInfo-pipelineStageCreationFeedbackCount-06594"></a>
  VUID-VkGraphicsPipelineCreateInfo-pipelineStageCreationFeedbackCount-06594  
  If
  [VkPipelineCreationFeedbackCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedbackCreateInfo.html)::`pipelineStageCreationFeedbackCount`
  is not `0`, it **must** be equal to `stageCount`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-06595"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-06595"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-06595  
  If `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the pipeline
  is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> or
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>, and
  [VkMultiviewPerViewAttributesInfoNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultiviewPerViewAttributesInfoNVX.html)::`perViewAttributesPositionXOnly`
  is `VK_TRUE` then
  [VkMultiviewPerViewAttributesInfoNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultiviewPerViewAttributesInfoNVX.html)::`perViewAttributes`
  **must** also be `VK_TRUE`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06596"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06596"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06596  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes only one of
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT` or
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, and an element
  of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes the other flag, the value of
  [VkMultiviewPerViewAttributesInfoNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultiviewPerViewAttributesInfoNVX.html)::`perViewAttributes`
  specified in both this pipeline and the library **must** be equal

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pLibraries-06597"
  id="VUID-VkGraphicsPipelineCreateInfo-pLibraries-06597"></a>
  VUID-VkGraphicsPipelineCreateInfo-pLibraries-06597  
  If one element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT` and
  another element includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, the value of
  [VkMultiviewPerViewAttributesInfoNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultiviewPerViewAttributesInfoNVX.html)::`perViewAttributes`
  specified in both libraries **must** be equal

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06598"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06598"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06598  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes only one of
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT` or
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, and an element
  of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes the other flag, the value of
  [VkMultiviewPerViewAttributesInfoNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultiviewPerViewAttributesInfoNVX.html)::`perViewAttributesPositionXOnly`
  specified in both this pipeline and the library **must** be equal

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pLibraries-06599"
  id="VUID-VkGraphicsPipelineCreateInfo-pLibraries-06599"></a>
  VUID-VkGraphicsPipelineCreateInfo-pLibraries-06599  
  If one element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT` and
  another element includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, the value of
  [VkMultiviewPerViewAttributesInfoNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultiviewPerViewAttributesInfoNVX.html)::`perViewAttributesPositionXOnly`
  specified in both libraries **must** be equal

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pStages-06600"
  id="VUID-VkGraphicsPipelineCreateInfo-pStages-06600"></a>
  VUID-VkGraphicsPipelineCreateInfo-pStages-06600  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> or
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>, `pStages`
  **must** be a valid pointer to an array of `stageCount` valid
  [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html)
  structures

- <a href="#VUID-VkGraphicsPipelineCreateInfo-stageCount-09587"
  id="VUID-VkGraphicsPipelineCreateInfo-stageCount-09587"></a>
  VUID-VkGraphicsPipelineCreateInfo-stageCount-09587  
  If the pipeline does not require <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> or
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>, `stageCount`
  **must** be zero

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pRasterizationState-06601"
  id="VUID-VkGraphicsPipelineCreateInfo-pRasterizationState-06601"></a>
  VUID-VkGraphicsPipelineCreateInfo-pRasterizationState-06601  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-pRasterizationState-null"
  target="_blank" rel="noopener">related dynamic state is not set</a>,
  `pRasterizationState` **must** be a valid pointer to a valid
  [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)
  structure

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pRasterizationState-09039"
  id="VUID-VkGraphicsPipelineCreateInfo-pRasterizationState-09039"></a>
  VUID-VkGraphicsPipelineCreateInfo-pRasterizationState-09039  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT`, and
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-pMultisampleState-null"
  target="_blank" rel="noopener">related dynamic state is not set</a>,
  then `pMultisampleState` **must** be a valid pointer to a valid
  [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)
  structure

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pRasterizationState-09040"
  id="VUID-VkGraphicsPipelineCreateInfo-pRasterizationState-09040"></a>
  VUID-VkGraphicsPipelineCreateInfo-pRasterizationState-09040  
  If `pRasterizationState` is not `NULL` it **must** be a valid pointer
  to a valid
  [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)
  structure

- <a href="#VUID-VkGraphicsPipelineCreateInfo-layout-06602"
  id="VUID-VkGraphicsPipelineCreateInfo-layout-06602"></a>
  VUID-VkGraphicsPipelineCreateInfo-layout-06602  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>,
  `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html)
  handle

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-06603"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-06603"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-06603  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output state</a>, and
  `renderPass` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `renderPass` **must** be a valid [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html)
  handle

- <a href="#VUID-VkGraphicsPipelineCreateInfo-stageCount-09530"
  id="VUID-VkGraphicsPipelineCreateInfo-stageCount-09530"></a>
  VUID-VkGraphicsPipelineCreateInfo-stageCount-09530  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>,
  `stageCount` **must** be greater than `0`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-graphicsPipelineLibrary-06606"
  id="VUID-VkGraphicsPipelineCreateInfo-graphicsPipelineLibrary-06606"></a>
  VUID-VkGraphicsPipelineCreateInfo-graphicsPipelineLibrary-06606  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-graphicsPipelineLibrary"
  target="_blank" rel="noopener"><code>graphicsPipelineLibrary</code></a>
  feature is not enabled, `flags` **must** not include
  `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06608"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06608"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06608  
  If the pipeline defines, or includes as libraries, all the state
  subsets required for a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-complete"
  target="_blank" rel="noopener">complete graphics pipeline</a>, `flags`
  **must** not include `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06609"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06609"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06609  
  If `flags` includes
  `VK_PIPELINE_CREATE_LINK_TIME_OPTIMIZATION_BIT_EXT`, pipeline
  libraries included via
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)
  **must** have been created with
  `VK_PIPELINE_CREATE_RETAIN_LINK_TIME_OPTIMIZATION_INFO_BIT_EXT`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-09245"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-09245"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-09245  
  If `flags` includes
  `VK_PIPELINE_CREATE_RETAIN_LINK_TIME_OPTIMIZATION_INFO_BIT_EXT`,
  `flags` **must** also include `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06610"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06610"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06610  
  If `flags` includes
  `VK_PIPELINE_CREATE_RETAIN_LINK_TIME_OPTIMIZATION_INFO_BIT_EXT`,
  pipeline libraries included via
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)
  **must** have been created with
  `VK_PIPELINE_CREATE_RETAIN_LINK_TIME_OPTIMIZATION_INFO_BIT_EXT`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pLibraries-06611"
  id="VUID-VkGraphicsPipelineCreateInfo-pLibraries-06611"></a>
  VUID-VkGraphicsPipelineCreateInfo-pLibraries-06611  
  Any pipeline libraries included via
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  **must** not include any <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets"
  target="_blank" rel="noopener">state subset</a> already defined by
  this structure or defined by any other pipeline library in
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06612"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06612"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06612  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes only one of
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT` or
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, and an element
  of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes the other flag, and `layout` was not created with
  `VK_PIPELINE_LAYOUT_CREATE_INDEPENDENT_SETS_BIT_EXT`, then the
  `layout` used by this pipeline and the library **must** be
  *identically defined*

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pLibraries-06613"
  id="VUID-VkGraphicsPipelineCreateInfo-pLibraries-06613"></a>
  VUID-VkGraphicsPipelineCreateInfo-pLibraries-06613  
  If one element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT` and
  another element includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, and the
  `layout` specified by either library was not created with
  `VK_PIPELINE_LAYOUT_CREATE_INDEPENDENT_SETS_BIT_EXT`, then the
  `layout` used by each library **must** be *identically defined*

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06614"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06614"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06614  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes only one of
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT` or
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, an element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes the other subset, and `layout` was created with
  `VK_PIPELINE_LAYOUT_CREATE_INDEPENDENT_SETS_BIT_EXT`, then the
  `layout` used by the library **must** also have been created with
  `VK_PIPELINE_LAYOUT_CREATE_INDEPENDENT_SETS_BIT_EXT`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pLibraries-06615"
  id="VUID-VkGraphicsPipelineCreateInfo-pLibraries-06615"></a>
  VUID-VkGraphicsPipelineCreateInfo-pLibraries-06615  
  If one element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT` and
  another element includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, and the
  `layout` specified by either library was created with
  `VK_PIPELINE_LAYOUT_CREATE_INDEPENDENT_SETS_BIT_EXT`, then the
  `layout` used by both libraries **must** have been created with
  `VK_PIPELINE_LAYOUT_CREATE_INDEPENDENT_SETS_BIT_EXT`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06616"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06616"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06616  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes only one of
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT` or
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, an element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes the other subset, and `layout` was created with
  `VK_PIPELINE_LAYOUT_CREATE_INDEPENDENT_SETS_BIT_EXT`, elements of the
  `pSetLayouts` array which `layout` was created with that are not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) **must** be *identically
  defined* to the element at the same index of `pSetLayouts` used to
  create the library’s `layout`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pLibraries-06617"
  id="VUID-VkGraphicsPipelineCreateInfo-pLibraries-06617"></a>
  VUID-VkGraphicsPipelineCreateInfo-pLibraries-06617  
  If one element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT` and
  another element includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, and the
  `layout` specified by either library was created with
  `VK_PIPELINE_LAYOUT_CREATE_INDEPENDENT_SETS_BIT_EXT`, elements of the
  `pSetLayouts` array which either `layout` was created with that are
  not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) **must** be *identically
  defined* to the element at the same index of `pSetLayouts` used to
  create the other library’s `layout`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06618"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06618"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06618  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes only one of
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT` or
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, and an element
  of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes the other flag, any descriptor set layout *N* specified by
  `layout` in both this pipeline and the library which include bindings
  accessed by shader stages in each **must** be *identically defined*

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pLibraries-06619"
  id="VUID-VkGraphicsPipelineCreateInfo-pLibraries-06619"></a>
  VUID-VkGraphicsPipelineCreateInfo-pLibraries-06619  
  If one element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT` and
  another element includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, any descriptor
  set layout *N* specified by `layout` in both libraries which include
  bindings accessed by shader stages in each **must** be *identically
  defined*

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06620"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06620"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06620  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes only one of
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT` or
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, and an element
  of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes the other flag, push constants specified in `layout` in both
  this pipeline and the library which are available to shader stages in
  each **must** be *identically defined*

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pLibraries-06621"
  id="VUID-VkGraphicsPipelineCreateInfo-pLibraries-06621"></a>
  VUID-VkGraphicsPipelineCreateInfo-pLibraries-06621  
  If one element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT` and
  another element includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, push constants
  specified in `layout` in both this pipeline and the library which are
  available to shader stages in each **must** be *identically defined*

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06679"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06679"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06679  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes only one of
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT` or
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, an element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes the other subset, any element of the `pSetLayouts` array when
  `layout` was created and the corresponding element of the
  `pSetLayouts` array used to create the library’s `layout` **must** not
  both be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pLibraries-06681"
  id="VUID-VkGraphicsPipelineCreateInfo-pLibraries-06681"></a>
  VUID-VkGraphicsPipelineCreateInfo-pLibraries-06681  
  If one element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT` and
  another element includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, and any
  element of the `pSetLayouts` array used to create each library’s
  `layout` was [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), then the
  corresponding element of the `pSetLayouts` array used to create the
  other library’s `layout` **must** not be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06756"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06756"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06756  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes only one of
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT` or
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, an element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes the other subset, and any element of the `pSetLayouts` array
  which `layout` was created with was
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), then the corresponding element
  of the `pSetLayouts` array used to create the library’s `layout`
  **must** not have shader bindings for shaders in the other subset

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06757"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06757"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06757  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes only one of
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT` or
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, an element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes the other subset, and any element of the `pSetLayouts` array
  used to create the library’s `layout` was
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), then the corresponding element
  of the `pSetLayouts` array used to create this pipeline’s `layout`
  **must** not have shader bindings for shaders in the other subset

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pLibraries-06758"
  id="VUID-VkGraphicsPipelineCreateInfo-pLibraries-06758"></a>
  VUID-VkGraphicsPipelineCreateInfo-pLibraries-06758  
  If one element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT` and
  another element includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, and any
  element of the `pSetLayouts` array used to create each library’s
  `layout` was [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), then the
  corresponding element of the `pSetLayouts` array used to create the
  other library’s `layout` **must** not have shader bindings for shaders
  in the other subset

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06682"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06682"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06682  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes both
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT` and
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, `layout`
  **must** have been created with no elements of the `pSetLayouts` array
  set to [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06683"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06683"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06683  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT` and
  `pRasterizationState->rasterizerDiscardEnable` is `VK_TRUE`, `layout`
  **must** have been created with no elements of the `pSetLayouts` array
  set to [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06684"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06684"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06684  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes at least one of and no more than two of
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT`,
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, or
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_OUTPUT_INTERFACE_BIT_EXT`, and
  an element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes one of the other flags, the value of `subpass` **must** be
  equal to that used to create the library

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pLibraries-06623"
  id="VUID-VkGraphicsPipelineCreateInfo-pLibraries-06623"></a>
  VUID-VkGraphicsPipelineCreateInfo-pLibraries-06623  
  If one element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes at least one of and no more than two of
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT`,
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, or
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_OUTPUT_INTERFACE_BIT_EXT`, and
  another element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes one of the other flags, the value of `subpass` used to create
  each library **must** be identical

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderpass-06624"
  id="VUID-VkGraphicsPipelineCreateInfo-renderpass-06624"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderpass-06624  
  If `renderpass` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes at least one of and no more than two of
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT`,
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, or
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_OUTPUT_INTERFACE_BIT_EXT`, and
  an element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes one of the other flags, `renderPass` **must** be compatible
  with that used to create the library

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderpass-06625"
  id="VUID-VkGraphicsPipelineCreateInfo-renderpass-06625"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderpass-06625  
  If `renderpass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes at least one of and no more than two of
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT`,
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, or
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_OUTPUT_INTERFACE_BIT_EXT`, and
  an element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes one of the other flags, the value of `renderPass` used to
  create that library **must** also be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06626"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06626"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06626  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes at least one of and no more than two of
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT`,
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, or
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_OUTPUT_INTERFACE_BIT_EXT`, an
  element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes one of the other flags, and `renderPass` is
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the value of
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`viewMask`
  used by this pipeline and that specified by the library **must** be
  identical

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pLibraries-06627"
  id="VUID-VkGraphicsPipelineCreateInfo-pLibraries-06627"></a>
  VUID-VkGraphicsPipelineCreateInfo-pLibraries-06627  
  If one element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes at least one of and no more than two of
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT`,
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, or
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_OUTPUT_INTERFACE_BIT_EXT`,
  another element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes one of the other flags, and `renderPass` was
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) for both libraries, the value of
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`viewMask`
  set by each library **must** be identical

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pLibraries-06628"
  id="VUID-VkGraphicsPipelineCreateInfo-pLibraries-06628"></a>
  VUID-VkGraphicsPipelineCreateInfo-pLibraries-06628  
  If one element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes at least one of and no more than two of
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT`,
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, or
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_OUTPUT_INTERFACE_BIT_EXT`, and
  another element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes one of the other flags, the `renderPass` objects used to
  create each library **must** be compatible or all equal to
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderpass-06631"
  id="VUID-VkGraphicsPipelineCreateInfo-renderpass-06631"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderpass-06631  
  If `renderPass` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the
  pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>, and the
  [`VK_EXT_extended_dynamic_state3`](VK_EXT_extended_dynamic_state3.html)
  extension is not enabled or any of the
  `VK_DYNAMIC_STATE_RASTERIZATION_SAMPLES_EXT`,
  `VK_DYNAMIC_STATE_SAMPLE_MASK_EXT`, or
  `VK_DYNAMIC_STATE_ALPHA_TO_COVERAGE_ENABLE_EXT` dynamic states is not
  set, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-alphaToOne"
  target="_blank" rel="noopener">alphaToOne</a> is enabled on the device
  and `VK_DYNAMIC_STATE_ALPHA_TO_ONE_ENABLE_EXT` is not set, then
  `pMultisampleState` **must** be a valid pointer to a valid
  [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)
  structure

- <a href="#VUID-VkGraphicsPipelineCreateInfo-Input-06632"
  id="VUID-VkGraphicsPipelineCreateInfo-Input-06632"></a>
  VUID-VkGraphicsPipelineCreateInfo-Input-06632  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> with a
  fragment shader that either enables <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-sampleshading"
  target="_blank" rel="noopener">sample shading</a> or decorates any
  variable in the `Input` storage class with `Sample`, and the
  [`VK_EXT_extended_dynamic_state3`](VK_EXT_extended_dynamic_state3.html)
  extension is not enabled or any of the
  `VK_DYNAMIC_STATE_RASTERIZATION_SAMPLES_EXT`,
  `VK_DYNAMIC_STATE_SAMPLE_MASK_EXT`, or
  `VK_DYNAMIC_STATE_ALPHA_TO_COVERAGE_ENABLE_EXT` dynamic states is not
  set, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-alphaToOne"
  target="_blank" rel="noopener">alphaToOne</a> is enabled on the device
  and `VK_DYNAMIC_STATE_ALPHA_TO_ONE_ENABLE_EXT` is not set, then
  `pMultisampleState` **must** be a valid pointer to a valid
  [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)
  structure

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06633"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06633"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06633  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT` with a
  `pMultisampleState` that was not `NULL`, and an element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  was created with
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_OUTPUT_INTERFACE_BIT_EXT`,
  `pMultisampleState` **must** be *identically defined* to that used to
  create the library

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pLibraries-06634"
  id="VUID-VkGraphicsPipelineCreateInfo-pLibraries-06634"></a>
  VUID-VkGraphicsPipelineCreateInfo-pLibraries-06634  
  If an element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  was created with
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT` with a
  `pMultisampleState` that was not `NULL`, and if
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_OUTPUT_INTERFACE_BIT_EXT`,
  `pMultisampleState` **must** be *identically defined* to that used to
  create the library

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pLibraries-06635"
  id="VUID-VkGraphicsPipelineCreateInfo-pLibraries-06635"></a>
  VUID-VkGraphicsPipelineCreateInfo-pLibraries-06635  
  If one element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  was created with
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT` with a
  `pMultisampleState` that was not `NULL`, and if a different element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  was created with
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_OUTPUT_INTERFACE_BIT_EXT`, the
  `pMultisampleState` used to create each library **must** be
  *identically defined*

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pLibraries-06636"
  id="VUID-VkGraphicsPipelineCreateInfo-pLibraries-06636"></a>
  VUID-VkGraphicsPipelineCreateInfo-pLibraries-06636  
  If one element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  was created with
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_OUTPUT_INTERFACE_BIT_EXT` and a
  value of `pMultisampleState->sampleShadingEnable` equal `VK_TRUE`, and
  if a different element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  was created with
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, the
  `pMultisampleState` used to create each library **must** be
  *identically defined*

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06637"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06637"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06637  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_OUTPUT_INTERFACE_BIT_EXT`,
  `pMultisampleState->sampleShadingEnable` is `VK_TRUE`, and an element
  of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  was created with
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`,
  `pMultisampleState` **must** be *identically defined* to that used to
  create the library

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pLibraries-09567"
  id="VUID-VkGraphicsPipelineCreateInfo-pLibraries-09567"></a>
  VUID-VkGraphicsPipelineCreateInfo-pLibraries-09567  
  If one element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  was created with
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_OUTPUT_INTERFACE_BIT_EXT` and a
  value of `pMultisampleState->sampleShadingEnable` equal `VK_TRUE`, and
  if
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`,
  `pMultisampleState` **must** be *identically defined* to that used to
  create the library

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06638"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06638"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06638  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes only one of
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT` or
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, and an element
  of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes the other flag, values specified in
  [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html)
  for both this pipeline and that library **must** be identical

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pLibraries-06639"
  id="VUID-VkGraphicsPipelineCreateInfo-pLibraries-06639"></a>
  VUID-VkGraphicsPipelineCreateInfo-pLibraries-06639  
  If one element of
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT` and
  another element includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, values
  specified in
  [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html)
  for both this pipeline and that library **must** be identical

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06640"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06640"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06640  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT` or
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, `pStages`
  **must** be a valid pointer to an array of `stageCount` valid
  [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html)
  structures

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06642"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06642"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06642  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT` or
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, `layout`
  **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) handle

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06643"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06643"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06643  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT`, or
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`,
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_OUTPUT_INTERFACE_BIT_EXT`, and
  `renderPass` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `renderPass` **must** be a valid [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html)
  handle

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06644"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06644"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06644  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT` or
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, `stageCount`
  **must** be greater than `0`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06645"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06645"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06645  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  is non-zero, if `flags` includes
  `VK_PIPELINE_CREATE_CAPTURE_INTERNAL_REPRESENTATIONS_BIT_KHR`, any
  libraries **must** have also been created with
  `VK_PIPELINE_CREATE_CAPTURE_INTERNAL_REPRESENTATIONS_BIT_KHR`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pLibraries-06646"
  id="VUID-VkGraphicsPipelineCreateInfo-pLibraries-06646"></a>
  VUID-VkGraphicsPipelineCreateInfo-pLibraries-06646  
  If
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes more than one library, and any library was created with
  `VK_PIPELINE_CREATE_CAPTURE_INTERNAL_REPRESENTATIONS_BIT_KHR`, all
  libraries **must** have also been created with
  `VK_PIPELINE_CREATE_CAPTURE_INTERNAL_REPRESENTATIONS_BIT_KHR`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pLibraries-06647"
  id="VUID-VkGraphicsPipelineCreateInfo-pLibraries-06647"></a>
  VUID-VkGraphicsPipelineCreateInfo-pLibraries-06647  
  If
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`
  includes at least one library,
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  is non-zero, and any library was created with
  `VK_PIPELINE_CREATE_CAPTURE_INTERNAL_REPRESENTATIONS_BIT_KHR`, `flags`
  **must** include
  `VK_PIPELINE_CREATE_CAPTURE_INTERNAL_REPRESENTATIONS_BIT_KHR`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-None-07826"
  id="VUID-VkGraphicsPipelineCreateInfo-None-07826"></a>
  VUID-VkGraphicsPipelineCreateInfo-None-07826  
  If the pipeline includes a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-complete"
  target="_blank" rel="noopener">complete set of state</a>, and there
  are no libraries included in
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)::`pLibraries`,
  then [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) **must** be a valid
  pipeline layout

- <a href="#VUID-VkGraphicsPipelineCreateInfo-layout-07827"
  id="VUID-VkGraphicsPipelineCreateInfo-layout-07827"></a>
  VUID-VkGraphicsPipelineCreateInfo-layout-07827  
  If the pipeline includes a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-complete"
  target="_blank" rel="noopener">complete set of state</a> specified
  entirely by libraries, and each library was created with a
  [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) created without
  `VK_PIPELINE_LAYOUT_CREATE_INDEPENDENT_SETS_BIT_EXT`, then `layout`
  **must** be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-compatibility"
  target="_blank" rel="noopener">compatible</a> with the layouts in
  those libraries

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06729"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06729"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06729  
  If `flags` includes
  `VK_PIPELINE_CREATE_LINK_TIME_OPTIMIZATION_BIT_EXT`, the pipeline
  includes a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-complete"
  target="_blank" rel="noopener">complete set of state</a> specified
  entirely by libraries, and each library was created with a
  [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) created with
  `VK_PIPELINE_LAYOUT_CREATE_INDEPENDENT_SETS_BIT_EXT`, then `layout`
  **must** be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-compatibility"
  target="_blank" rel="noopener">compatible</a> with the union of the
  libraries' pipeline layouts other than the inclusion/exclusion of
  `VK_PIPELINE_LAYOUT_CREATE_INDEPENDENT_SETS_BIT_EXT`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-06730"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-06730"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-06730  
  If `flags` does not include
  `VK_PIPELINE_CREATE_LINK_TIME_OPTIMIZATION_BIT_EXT`, the pipeline
  includes a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-complete"
  target="_blank" rel="noopener">complete set of state</a> specified
  entirely by libraries, and each library was created with a
  [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) created with
  `VK_PIPELINE_LAYOUT_CREATE_INDEPENDENT_SETS_BIT_EXT`, then `layout`
  **must** be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-compatibility"
  target="_blank" rel="noopener">compatible</a> with the union of the
  libraries' pipeline layouts

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-conservativePointAndLineRasterization-08892"
  id="VUID-VkGraphicsPipelineCreateInfo-conservativePointAndLineRasterization-08892"></a>
  VUID-VkGraphicsPipelineCreateInfo-conservativePointAndLineRasterization-08892  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-conservativePointAndLineRasterization"
  target="_blank"
  rel="noopener"><code>conservativePointAndLineRasterization</code></a>
  is not supported; the pipeline is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-vertex-input"
  target="_blank" rel="noopener">vertex input state</a> and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>; the
  pipeline does not include a geometry shader; and the value of
  [VkPipelineInputAssemblyStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineInputAssemblyStateCreateInfo.html)::`topology`
  is `VK_PRIMITIVE_TOPOLOGY_POINT_LIST`,
  `VK_PRIMITIVE_TOPOLOGY_LINE_LIST`, or
  `VK_PRIMITIVE_TOPOLOGY_LINE_STRIP`, and either
  `VK_DYNAMIC_STATE_PRIMITIVE_TOPOLOGY` dynamic state is not enabled or
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-dynamicPrimitiveTopologyUnrestricted"
  target="_blank"
  rel="noopener"><code>dynamicPrimitiveTopologyUnrestricted</code></a>
  is `VK_FALSE`, then
  [VkPipelineRasterizationConservativeStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationConservativeStateCreateInfoEXT.html)::`conservativeRasterizationMode`
  **must** be `VK_CONSERVATIVE_RASTERIZATION_MODE_DISABLED_EXT`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-conservativePointAndLineRasterization-06760"
  id="VUID-VkGraphicsPipelineCreateInfo-conservativePointAndLineRasterization-06760"></a>
  VUID-VkGraphicsPipelineCreateInfo-conservativePointAndLineRasterization-06760  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-conservativePointAndLineRasterization"
  target="_blank"
  rel="noopener"><code>conservativePointAndLineRasterization</code></a>
  is not supported, the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  the pipeline includes a geometry shader with either the `OutputPoints`
  or `OutputLineStrip` execution modes,
  [VkPipelineRasterizationConservativeStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationConservativeStateCreateInfoEXT.html)::`conservativeRasterizationMode`
  **must** be `VK_CONSERVATIVE_RASTERIZATION_MODE_DISABLED_EXT`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-conservativePointAndLineRasterization-06761"
  id="VUID-VkGraphicsPipelineCreateInfo-conservativePointAndLineRasterization-06761"></a>
  VUID-VkGraphicsPipelineCreateInfo-conservativePointAndLineRasterization-06761  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-conservativePointAndLineRasterization"
  target="_blank"
  rel="noopener"><code>conservativePointAndLineRasterization</code></a>
  is not supported, the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  the pipeline includes a mesh shader with either the `OutputPoints` or
  `OutputLinesNV` execution modes,
  [VkPipelineRasterizationConservativeStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationConservativeStateCreateInfoEXT.html)::`conservativeRasterizationMode`
  **must** be `VK_CONSERVATIVE_RASTERIZATION_MODE_DISABLED_EXT`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pStages-06894"
  id="VUID-VkGraphicsPipelineCreateInfo-pStages-06894"></a>
  VUID-VkGraphicsPipelineCreateInfo-pStages-06894  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> but
  not <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>, elements of
  `pStages` **must** not have `stage` set to
  `VK_SHADER_STAGE_FRAGMENT_BIT`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pStages-06895"
  id="VUID-VkGraphicsPipelineCreateInfo-pStages-06895"></a>
  VUID-VkGraphicsPipelineCreateInfo-pStages-06895  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> but not <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>,
  elements of `pStages` **must** not have `stage` set to a shader stage
  which participates in pre-rasterization

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pStages-06896"
  id="VUID-VkGraphicsPipelineCreateInfo-pStages-06896"></a>
  VUID-VkGraphicsPipelineCreateInfo-pStages-06896  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, all
  elements of `pStages` **must** have a `stage` set to a shader stage
  which participates in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>

- <a href="#VUID-VkGraphicsPipelineCreateInfo-stage-06897"
  id="VUID-VkGraphicsPipelineCreateInfo-stage-06897"></a>
  VUID-VkGraphicsPipelineCreateInfo-stage-06897  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> and/or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, any
  value of `stage` **must** not be set in more than one element of
  `pStages`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3TessellationDomainOrigin-07370"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3TessellationDomainOrigin-07370"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3TessellationDomainOrigin-07370  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3TessellationDomainOrigin"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3TessellationDomainOrigin</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_TESSELLATION_DOMAIN_ORIGIN_EXT`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3DepthClampEnable-07371"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3DepthClampEnable-07371"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3DepthClampEnable-07371  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3DepthClampEnable"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3DepthClampEnable</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_DEPTH_CLAMP_ENABLE_EXT`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3PolygonMode-07372"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3PolygonMode-07372"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3PolygonMode-07372  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3PolygonMode"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3PolygonMode</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_POLYGON_MODE_EXT`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3RasterizationSamples-07373"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3RasterizationSamples-07373"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3RasterizationSamples-07373  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3RasterizationSamples"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3RasterizationSamples</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_RASTERIZATION_SAMPLES_EXT`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3SampleMask-07374"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3SampleMask-07374"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3SampleMask-07374  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3SampleMask"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3SampleMask</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_SAMPLE_MASK_EXT`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3AlphaToCoverageEnable-07375"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3AlphaToCoverageEnable-07375"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3AlphaToCoverageEnable-07375  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3AlphaToCoverageEnable"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3AlphaToCoverageEnable</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_ALPHA_TO_COVERAGE_ENABLE_EXT`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3AlphaToOneEnable-07376"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3AlphaToOneEnable-07376"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3AlphaToOneEnable-07376  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3AlphaToOneEnable"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3AlphaToOneEnable</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_ALPHA_TO_ONE_ENABLE_EXT`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3LogicOpEnable-07377"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3LogicOpEnable-07377"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3LogicOpEnable-07377  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3LogicOpEnable"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3LogicOpEnable</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_LOGIC_OP_ENABLE_EXT`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ColorBlendEnable-07378"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ColorBlendEnable-07378"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ColorBlendEnable-07378  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3ColorBlendEnable"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3ColorBlendEnable</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_COLOR_BLEND_ENABLE_EXT`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ColorBlendEquation-07379"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ColorBlendEquation-07379"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ColorBlendEquation-07379  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3ColorBlendEquation"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3ColorBlendEquation</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_COLOR_BLEND_EQUATION_EXT`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ColorWriteMask-07380"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ColorWriteMask-07380"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ColorWriteMask-07380  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3ColorWriteMask"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3ColorWriteMask</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_COLOR_WRITE_MASK_EXT`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3RasterizationStream-07381"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3RasterizationStream-07381"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3RasterizationStream-07381  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3RasterizationStream"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3RasterizationStream</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_RASTERIZATION_STREAM_EXT`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ConservativeRasterizationMode-07382"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ConservativeRasterizationMode-07382"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ConservativeRasterizationMode-07382  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3ConservativeRasterizationMode"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3ConservativeRasterizationMode</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_CONSERVATIVE_RASTERIZATION_MODE_EXT`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ExtraPrimitiveOverestimationSize-07383"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ExtraPrimitiveOverestimationSize-07383"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ExtraPrimitiveOverestimationSize-07383  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3ExtraPrimitiveOverestimationSize"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3ExtraPrimitiveOverestimationSize</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_EXTRA_PRIMITIVE_OVERESTIMATION_SIZE_EXT`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicState-09639"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicState-09639"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicState-09639  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>,
  `pDynamicState` includes
  `VK_DYNAMIC_STATE_CONSERVATIVE_RASTERIZATION_MODE_EXT`, and
  `pDynamicState` does not include
  `VK_DYNAMIC_STATE_EXTRA_PRIMITIVE_OVERESTIMATION_SIZE_EXT`,
  `pRasterizationState` **must** include a
  [VkPipelineRasterizationConservativeStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationConservativeStateCreateInfoEXT.html)
  in its `pNext` chain

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3DepthClipEnable-07384"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3DepthClipEnable-07384"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3DepthClipEnable-07384  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3DepthClipEnable"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3DepthClipEnable</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_DEPTH_CLIP_ENABLE_EXT`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3SampleLocationsEnable-07385"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3SampleLocationsEnable-07385"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3SampleLocationsEnable-07385  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3SampleLocationsEnable"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3SampleLocationsEnable</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_ENABLE_EXT`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ColorBlendAdvanced-07386"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ColorBlendAdvanced-07386"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ColorBlendAdvanced-07386  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3ColorBlendAdvanced"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3ColorBlendAdvanced</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_COLOR_BLEND_ADVANCED_EXT`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ProvokingVertexMode-07387"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ProvokingVertexMode-07387"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ProvokingVertexMode-07387  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3ProvokingVertexMode"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3ProvokingVertexMode</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_PROVOKING_VERTEX_MODE_EXT`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3LineRasterizationMode-07388"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3LineRasterizationMode-07388"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3LineRasterizationMode-07388  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3LineRasterizationMode"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3LineRasterizationMode</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_LINE_RASTERIZATION_MODE_EXT`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3LineStippleEnable-07389"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3LineStippleEnable-07389"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3LineStippleEnable-07389  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3LineStippleEnable"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3LineStippleEnable</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_LINE_STIPPLE_ENABLE_EXT`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3DepthClipNegativeOneToOne-07390"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3DepthClipNegativeOneToOne-07390"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3DepthClipNegativeOneToOne-07390  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3DepthClipNegativeOneToOne"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3DepthClipNegativeOneToOne</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_DEPTH_CLIP_NEGATIVE_ONE_TO_ONE_EXT`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ViewportWScalingEnable-07391"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ViewportWScalingEnable-07391"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ViewportWScalingEnable-07391  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3ViewportWScalingEnable"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3ViewportWScalingEnable</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_VIEWPORT_W_SCALING_ENABLE_NV`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ViewportSwizzle-07392"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ViewportSwizzle-07392"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ViewportSwizzle-07392  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3ViewportSwizzle"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3ViewportSwizzle</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_VIEWPORT_SWIZZLE_NV`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3CoverageToColorEnable-07393"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3CoverageToColorEnable-07393"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3CoverageToColorEnable-07393  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3CoverageToColorEnable"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3CoverageToColorEnable</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_COVERAGE_TO_COLOR_ENABLE_NV`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3CoverageToColorLocation-07394"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3CoverageToColorLocation-07394"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3CoverageToColorLocation-07394  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3CoverageToColorLocation"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3CoverageToColorLocation</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_COVERAGE_TO_COLOR_LOCATION_NV`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3CoverageModulationMode-07395"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3CoverageModulationMode-07395"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3CoverageModulationMode-07395  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3CoverageModulationMode"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3CoverageModulationMode</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_COVERAGE_MODULATION_MODE_NV`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3CoverageModulationTableEnable-07396"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3CoverageModulationTableEnable-07396"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3CoverageModulationTableEnable-07396  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3CoverageModulationTableEnable"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3CoverageModulationTableEnable</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_COVERAGE_MODULATION_TABLE_ENABLE_NV`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3CoverageModulationTable-07397"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3CoverageModulationTable-07397"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3CoverageModulationTable-07397  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3CoverageModulationTable"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3CoverageModulationTable</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_COVERAGE_MODULATION_TABLE_NV`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3CoverageReductionMode-07398"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3CoverageReductionMode-07398"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3CoverageReductionMode-07398  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3CoverageReductionMode"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3CoverageReductionMode</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_COVERAGE_REDUCTION_MODE_NV`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3RepresentativeFragmentTestEnable-07399"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3RepresentativeFragmentTestEnable-07399"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3RepresentativeFragmentTestEnable-07399  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3RepresentativeFragmentTestEnable"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3RepresentativeFragmentTestEnable</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_REPRESENTATIVE_FRAGMENT_TEST_ENABLE_NV`

- <a
  href="#VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ShadingRateImageEnable-07400"
  id="VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ShadingRateImageEnable-07400"></a>
  VUID-VkGraphicsPipelineCreateInfo-extendedDynamicState3ShadingRateImageEnable-07400  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedDynamicState3ShadingRateImageEnable"
  target="_blank"
  rel="noopener"><code>extendedDynamicState3ShadingRateImageEnable</code></a>
  feature is not enabled, there **must** be no element of the
  `pDynamicStates` member of `pDynamicState` set to
  `VK_DYNAMIC_STATE_SHADING_RATE_IMAGE_ENABLE_NV`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-07401"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-07401"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-07401  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_OPACITY_MICROMAP_BIT_EXT`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-07997"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-07997"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-07997  
  `flags` **must** not include
  `VK_PIPELINE_CREATE_RAY_TRACING_DISPLACEMENT_MICROMAP_BIT_NV`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07730"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07730"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07730  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  no element of the `pDynamicStates` member of `pDynamicState` is
  `VK_DYNAMIC_STATE_VIEWPORT` or `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT`,
  and if <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiview-per-view-viewports"
  target="_blank"
  rel="noopener"><code>multiviewPerViewViewports</code></a> is enabled,
  then the index of the most significant bit in each element of
  [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassMultiviewCreateInfo.html)::`pViewMasks`
  **must** be less than `pViewportState->viewportCount`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07731"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07731"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicStates-07731  
  If the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>, and
  no element of the `pDynamicStates` member of `pDynamicState` is
  `VK_DYNAMIC_STATE_SCISSOR` or `VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT`,
  and if <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiview-per-view-viewports"
  target="_blank"
  rel="noopener"><code>multiviewPerViewViewports</code></a> is enabled,
  then the index of the most significant bit in each element of
  [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassMultiviewCreateInfo.html)::`pViewMasks`
  **must** be less than `pViewportState->scissorCount`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pStages-08711"
  id="VUID-VkGraphicsPipelineCreateInfo-pStages-08711"></a>
  VUID-VkGraphicsPipelineCreateInfo-pStages-08711  
  If `pStages` includes a fragment shader stage,
  `VK_DYNAMIC_STATE_DEPTH_WRITE_ENABLE` is not set in
  [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`,
  and the fragment shader declares the `EarlyFragmentTests` execution
  mode and uses `OpDepthAttachmentReadEXT`, the `depthWriteEnable`
  member of
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
  **must** be `VK_FALSE`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pStages-08712"
  id="VUID-VkGraphicsPipelineCreateInfo-pStages-08712"></a>
  VUID-VkGraphicsPipelineCreateInfo-pStages-08712  
  If `pStages` includes a fragment shader stage,
  `VK_DYNAMIC_STATE_STENCIL_WRITE_MASK` is not set in
  [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`,
  and the fragment shader declares the `EarlyFragmentTests` execution
  mode and uses `OpStencilAttachmentReadEXT`, the value of
  [VkStencilOpState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOpState.html)::`writeMask` for both
  `front` and `back` in
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
  **must** be `0`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-08744"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-08744"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-08744  
  If `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the pipeline
  requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output state</a> or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>, the pipeline
  enables <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-sampleshading"
  target="_blank" rel="noopener">sample shading</a>,
  `rasterizationSamples` is not dynamic, and the `pNext` chain includes
  a [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)
  structure, `rasterizationSamples` **must** be a valid
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) value that is set
  in `imageCreateSampleCounts` (as defined in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-creation-limits"
  target="_blank" rel="noopener">Image Creation Limits</a>) for every
  element of `depthAttachmentFormat`, `stencilAttachmentFormat` and the
  `pColorAttachmentFormats` array which is not `VK_FORMAT_UNDEFINED`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-08897"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-08897"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-08897  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_VERTEX_INPUT_INTERFACE_BIT_EXT`, <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> is
  specified either in a library or by the inclusion of
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT`, and
  that state includes a vertex shader stage in `pStages`, the pipeline
  **must** define <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-vertex-input"
  target="_blank" rel="noopener">vertex input state</a>

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-08898"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-08898"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-08898  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_VERTEX_INPUT_INTERFACE_BIT_EXT`, and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> is
  not specified, the pipeline **must** define <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-vertex-input"
  target="_blank" rel="noopener">vertex input state</a>

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-08899"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-08899"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-08899  
  If `flags` does not include `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`, <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> is
  specified either in a library or by the inclusion of
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT`, and
  that state includes a vertex shader stage in `pStages`, the pipeline
  **must** either define <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-vertex-input"
  target="_blank" rel="noopener">vertex input state</a> or include that
  state in a linked pipeline library

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-08900"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-08900"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-08900  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT` the
  pipeline **must** define <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a>

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-08901"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-08901"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-08901  
  If `flags` does not include `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`, the
  pipeline **must** either define <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> or
  include that state in a linked pipeline library

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-08903"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-08903"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-08903  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> is
  specified either in a library or by the inclusion of
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT`, and
  that state either includes
  `VK_DYNAMIC_STATE_RASTERIZER_DISCARD_ENABLE` or has
  `pRasterizationState->rasterizerDiscardEnable` set to `VK_FALSE`, the
  pipeline **must** define <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-08904"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-08904"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-08904  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, and
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> is
  not specified, the pipeline **must** define <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a>

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-08906"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-08906"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-08906  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> is
  specified either in a library or by the inclusion of
  `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_OUTPUT_INTERFACE_BIT_EXT`, and
  that state either includes
  `VK_DYNAMIC_STATE_RASTERIZER_DISCARD_ENABLE` or has
  `pRasterizationState->rasterizerDiscardEnable` set to `VK_FALSE`, the
  pipeline **must** define <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-08907"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-08907"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-08907  
  If
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html)::`flags`
  includes `VK_GRAPHICS_PIPELINE_LIBRARY_FRAGMENT_SHADER_BIT_EXT`, and
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> is
  not specified, the pipeline **must** define <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>

- <a href="#VUID-VkGraphicsPipelineCreateInfo-flags-08909"
  id="VUID-VkGraphicsPipelineCreateInfo-flags-08909"></a>
  VUID-VkGraphicsPipelineCreateInfo-flags-08909  
  If `flags` does not include `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`, <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader state</a> is
  specified either in a library or by the inclusion of
  `VK_GRAPHICS_PIPELINE_LIBRARY_PRE_RASTERIZATION_SHADERS_BIT_EXT`, and
  that state either includes
  `VK_DYNAMIC_STATE_RASTERIZER_DISCARD_ENABLE` or has
  `pRasterizationState->rasterizerDiscardEnable` set to `VK_FALSE`, the
  pipeline **must** define <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a> and
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> or include
  those states in linked pipeline libraries

- <a href="#VUID-VkGraphicsPipelineCreateInfo-None-09043"
  id="VUID-VkGraphicsPipelineCreateInfo-None-09043"></a>
  VUID-VkGraphicsPipelineCreateInfo-None-09043  
  If `pDynamicState->pDynamicStates` does not include
  `VK_DYNAMIC_STATE_COLOR_WRITE_MASK_EXT`, and the format of any color
  attachment is `VK_FORMAT_E5B9G9R9_UFLOAT_PACK32`, the `colorWriteMask`
  member of the corresponding element of
  `pColorBlendState->pAttachments` **must** either include all of
  `VK_COLOR_COMPONENT_R_BIT`, `VK_COLOR_COMPONENT_G_BIT`, and
  `VK_COLOR_COMPONENT_B_BIT`, or none of them

- <a href="#VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09301"
  id="VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09301"></a>
  VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09301  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-externalFormatResolve"
  target="_blank" rel="noopener"><code>externalFormatResolve</code></a>
  feature is enabled, the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)::`externalFormat`
  is not `0`,
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`viewMask`
  **must** be `0`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09304"
  id="VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09304"></a>
  VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09304  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-externalFormatResolve"
  target="_blank" rel="noopener"><code>externalFormatResolve</code></a>
  feature is enabled, the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)::`externalFormat`
  is not `0`, and `rasterizationSamples` is not dynamic,
  [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)::`rasterizationSamples`
  **must** be `1`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09305"
  id="VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09305"></a>
  VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09305  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-externalFormatResolve"
  target="_blank" rel="noopener"><code>externalFormatResolve</code></a>
  feature is enabled, the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)::`externalFormat`
  is not `0`, and `blendEnable` is not dynamic, the `blendEnable` member
  of each element of `pColorBlendState->pAttachments` **must** be
  `VK_FALSE`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09306"
  id="VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09306"></a>
  VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09306  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-externalFormatResolve"
  target="_blank" rel="noopener"><code>externalFormatResolve</code></a>
  feature is enabled, the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)::`externalFormat`
  is not `0`, and `pDynamicState->pDynamicStates` does not include
  `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR`,
  [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html)::`width`
  **must** be `1`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09307"
  id="VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09307"></a>
  VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09307  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-externalFormatResolve"
  target="_blank" rel="noopener"><code>externalFormatResolve</code></a>
  feature is enabled, the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)::`externalFormat`
  is not `0`, and `pDynamicState->pDynamicStates` does not include
  `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR`,
  [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html)::`height`
  **must** be `1`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09308"
  id="VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09308"></a>
  VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09308  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-externalFormatResolve"
  target="_blank" rel="noopener"><code>externalFormatResolve</code></a>
  feature is enabled, the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">pre-rasterization shader state</a> and
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)::`externalFormat`
  is not `0`, the last <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader stage</a>
  **must** not statically use a variable with the
  `PrimitiveShadingRateKHR` built-in

- <a href="#VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09309"
  id="VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09309"></a>
  VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09309  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-externalFormatResolve"
  target="_blank" rel="noopener"><code>externalFormatResolve</code></a>
  feature is enabled, the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)::`externalFormat`
  is not `0`,
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`colorAttachmentCount`
  **must** be `1`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09310"
  id="VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09310"></a>
  VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09310  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-externalFormatResolve"
  target="_blank" rel="noopener"><code>externalFormatResolve</code></a>
  feature is enabled, the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)::`externalFormat`
  is not `0`, the fragment shader **must** not declare the
  `DepthReplacing` or `StencilRefReplacingEXT` execution modes

- <a href="#VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09313"
  id="VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09313"></a>
  VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09313  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-externalFormatResolve"
  target="_blank" rel="noopener"><code>externalFormatResolve</code></a>
  feature is enabled, the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  `renderPass` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `subpass`
  includes an external format resolve attachment, and
  `rasterizationSamples` is not dynamic,
  [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)::`rasterizationSamples`
  **must** be `VK_SAMPLE_COUNT_1_BIT`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09314"
  id="VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09314"></a>
  VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09314  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-externalFormatResolve"
  target="_blank" rel="noopener"><code>externalFormatResolve</code></a>
  feature is enabled, the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  `renderPass` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `subpass`
  includes an external format resolve attachment, and `blendEnable` is
  not dynamic, the `blendEnable` member of each element of
  `pColorBlendState->pAttachments` **must** be `VK_FALSE`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09315"
  id="VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09315"></a>
  VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09315  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-externalFormatResolve"
  target="_blank" rel="noopener"><code>externalFormatResolve</code></a>
  feature is enabled, the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  `renderPass` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `subpass`
  includes an external format resolve attachment, and
  `pDynamicState->pDynamicStates` does not include
  `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR`,
  [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html)::`width`
  **must** be `1`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09316"
  id="VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09316"></a>
  VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09316  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-externalFormatResolve"
  target="_blank" rel="noopener"><code>externalFormatResolve</code></a>
  feature is enabled, the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  `renderPass` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `subpass`
  includes an external format resolve attachment, and
  `pDynamicState->pDynamicStates` does not include
  `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR`,
  [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html)::`height`
  **must** be `1`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09317"
  id="VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09317"></a>
  VUID-VkGraphicsPipelineCreateInfo-externalFormatResolve-09317  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-externalFormatResolve"
  target="_blank" rel="noopener"><code>externalFormatResolve</code></a>
  feature is enabled, the pipeline requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">pre-rasterization shader state</a> and
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output interface state</a>,
  `renderPass` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and
  `subpass` includes an external format resolve attachment, the last <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
  target="_blank" rel="noopener">pre-rasterization shader stage</a>
  **must** not statically use a variable with the
  `PrimitiveShadingRateKHR` built-in

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-09531"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-09531"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-09531  
  If the pipeline is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output state</a>, the value of
  `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and
  [VkRenderingInputAttachmentIndexInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInputAttachmentIndexInfoKHR.html)
  is included,
  [VkRenderingInputAttachmentIndexInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInputAttachmentIndexInfoKHR.html)::`colorAttachmentCount`
  **must** be equal to
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`colorAttachmentCount`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-09652"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-09652"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-09652  
  If the pipeline is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-shader"
  target="_blank" rel="noopener">fragment shader state</a> and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output state</a>, the value of
  `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and
  [VkRenderingInputAttachmentIndexInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInputAttachmentIndexInfoKHR.html)
  is not included, the fragment shader **must** not contain any input
  attachments with a `InputAttachmentIndex` greater than or equal to
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`colorAttachmentCount`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-renderPass-09532"
  id="VUID-VkGraphicsPipelineCreateInfo-renderPass-09532"></a>
  VUID-VkGraphicsPipelineCreateInfo-renderPass-09532  
  If the pipeline is being created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-fragment-output"
  target="_blank" rel="noopener">fragment output state</a>, and the
  value of `renderPass` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  [VkRenderingAttachmentLocationInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingAttachmentLocationInfoKHR.html)::`colorAttachmentCount`
  **must** be equal to
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`colorAttachmentCount`

Valid Usage (Implicit)

- <a href="#VUID-VkGraphicsPipelineCreateInfo-sType-sType"
  id="VUID-VkGraphicsPipelineCreateInfo-sType-sType"></a>
  VUID-VkGraphicsPipelineCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_GRAPHICS_PIPELINE_CREATE_INFO`

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pNext-pNext"
  id="VUID-VkGraphicsPipelineCreateInfo-pNext-pNext"></a>
  VUID-VkGraphicsPipelineCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoAMD.html),
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html),
  [VkGraphicsPipelineLibraryCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryCreateInfoEXT.html),
  [VkGraphicsPipelineShaderGroupsCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineShaderGroupsCreateInfoNV.html),
  [VkMultiviewPerViewAttributesInfoNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultiviewPerViewAttributesInfoNVX.html),
  [VkPipelineCompilerControlCreateInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCompilerControlCreateInfoAMD.html),
  [VkPipelineCreateFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlags2CreateInfoKHR.html),
  [VkPipelineCreationFeedbackCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedbackCreateInfo.html),
  [VkPipelineDiscardRectangleStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDiscardRectangleStateCreateInfoEXT.html),
  [VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html),
  [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html),
  [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html),
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html),
  [VkPipelineRepresentativeFragmentTestStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRepresentativeFragmentTestStateCreateInfoNV.html),
  [VkPipelineRobustnessCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRobustnessCreateInfoEXT.html),
  [VkRenderingAttachmentLocationInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingAttachmentLocationInfoKHR.html),
  or
  [VkRenderingInputAttachmentIndexInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInputAttachmentIndexInfoKHR.html)

- <a href="#VUID-VkGraphicsPipelineCreateInfo-sType-unique"
  id="VUID-VkGraphicsPipelineCreateInfo-sType-unique"></a>
  VUID-VkGraphicsPipelineCreateInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkGraphicsPipelineCreateInfo-pDynamicState-parameter"
  id="VUID-VkGraphicsPipelineCreateInfo-pDynamicState-parameter"></a>
  VUID-VkGraphicsPipelineCreateInfo-pDynamicState-parameter  
  If `pDynamicState` is not `NULL`, `pDynamicState` **must** be a valid
  pointer to a valid
  [VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)
  structure

- <a href="#VUID-VkGraphicsPipelineCreateInfo-commonparent"
  id="VUID-VkGraphicsPipelineCreateInfo-commonparent"></a>
  VUID-VkGraphicsPipelineCreateInfo-commonparent  
  Each of `basePipelineHandle`, `layout`, and `renderPass` that are
  valid handles of non-ignored parameters **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html),
[VkPipelineColorBlendStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateInfo.html),
[VkPipelineCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlags.html),
[VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html),
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html),
[VkPipelineInputAssemblyStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineInputAssemblyStateCreateInfo.html),
[VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html),
[VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html),
[VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html),
[VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html),
[VkPipelineTessellationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineTessellationStateCreateInfo.html),
[VkPipelineVertexInputStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineVertexInputStateCreateInfo.html),
[VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html),
[VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateGraphicsPipelines](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateGraphicsPipelines.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkGraphicsPipelineCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
