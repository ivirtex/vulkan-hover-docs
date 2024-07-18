# VkDynamicState(3) Manual Page

## Name

VkDynamicState - Indicate which dynamic state is taken from dynamic
state commands



## <a href="#_c_specification" class="anchor"></a>C Specification

The source of different pieces of dynamic state is specified by the
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`
property of the currently active pipeline, each of whose elements
**must** be one of the values:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkDynamicState {
    VK_DYNAMIC_STATE_VIEWPORT = 0,
    VK_DYNAMIC_STATE_SCISSOR = 1,
    VK_DYNAMIC_STATE_LINE_WIDTH = 2,
    VK_DYNAMIC_STATE_DEPTH_BIAS = 3,
    VK_DYNAMIC_STATE_BLEND_CONSTANTS = 4,
    VK_DYNAMIC_STATE_DEPTH_BOUNDS = 5,
    VK_DYNAMIC_STATE_STENCIL_COMPARE_MASK = 6,
    VK_DYNAMIC_STATE_STENCIL_WRITE_MASK = 7,
    VK_DYNAMIC_STATE_STENCIL_REFERENCE = 8,
  // Provided by VK_VERSION_1_3
    VK_DYNAMIC_STATE_CULL_MODE = 1000267000,
  // Provided by VK_VERSION_1_3
    VK_DYNAMIC_STATE_FRONT_FACE = 1000267001,
  // Provided by VK_VERSION_1_3
    VK_DYNAMIC_STATE_PRIMITIVE_TOPOLOGY = 1000267002,
  // Provided by VK_VERSION_1_3
    VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT = 1000267003,
  // Provided by VK_VERSION_1_3
    VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT = 1000267004,
  // Provided by VK_VERSION_1_3
    VK_DYNAMIC_STATE_VERTEX_INPUT_BINDING_STRIDE = 1000267005,
  // Provided by VK_VERSION_1_3
    VK_DYNAMIC_STATE_DEPTH_TEST_ENABLE = 1000267006,
  // Provided by VK_VERSION_1_3
    VK_DYNAMIC_STATE_DEPTH_WRITE_ENABLE = 1000267007,
  // Provided by VK_VERSION_1_3
    VK_DYNAMIC_STATE_DEPTH_COMPARE_OP = 1000267008,
  // Provided by VK_VERSION_1_3
    VK_DYNAMIC_STATE_DEPTH_BOUNDS_TEST_ENABLE = 1000267009,
  // Provided by VK_VERSION_1_3
    VK_DYNAMIC_STATE_STENCIL_TEST_ENABLE = 1000267010,
  // Provided by VK_VERSION_1_3
    VK_DYNAMIC_STATE_STENCIL_OP = 1000267011,
  // Provided by VK_VERSION_1_3
    VK_DYNAMIC_STATE_RASTERIZER_DISCARD_ENABLE = 1000377001,
  // Provided by VK_VERSION_1_3
    VK_DYNAMIC_STATE_DEPTH_BIAS_ENABLE = 1000377002,
  // Provided by VK_VERSION_1_3
    VK_DYNAMIC_STATE_PRIMITIVE_RESTART_ENABLE = 1000377004,
  // Provided by VK_NV_clip_space_w_scaling
    VK_DYNAMIC_STATE_VIEWPORT_W_SCALING_NV = 1000087000,
  // Provided by VK_EXT_discard_rectangles
    VK_DYNAMIC_STATE_DISCARD_RECTANGLE_EXT = 1000099000,
  // Provided by VK_EXT_discard_rectangles
    VK_DYNAMIC_STATE_DISCARD_RECTANGLE_ENABLE_EXT = 1000099001,
  // Provided by VK_EXT_discard_rectangles
    VK_DYNAMIC_STATE_DISCARD_RECTANGLE_MODE_EXT = 1000099002,
  // Provided by VK_EXT_sample_locations
    VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_EXT = 1000143000,
  // Provided by VK_KHR_ray_tracing_pipeline
    VK_DYNAMIC_STATE_RAY_TRACING_PIPELINE_STACK_SIZE_KHR = 1000347000,
  // Provided by VK_NV_shading_rate_image
    VK_DYNAMIC_STATE_VIEWPORT_SHADING_RATE_PALETTE_NV = 1000164004,
  // Provided by VK_NV_shading_rate_image
    VK_DYNAMIC_STATE_VIEWPORT_COARSE_SAMPLE_ORDER_NV = 1000164006,
  // Provided by VK_NV_scissor_exclusive
    VK_DYNAMIC_STATE_EXCLUSIVE_SCISSOR_ENABLE_NV = 1000205000,
  // Provided by VK_NV_scissor_exclusive
    VK_DYNAMIC_STATE_EXCLUSIVE_SCISSOR_NV = 1000205001,
  // Provided by VK_KHR_fragment_shading_rate
    VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR = 1000226000,
  // Provided by VK_EXT_vertex_input_dynamic_state
    VK_DYNAMIC_STATE_VERTEX_INPUT_EXT = 1000352000,
  // Provided by VK_EXT_extended_dynamic_state2
    VK_DYNAMIC_STATE_PATCH_CONTROL_POINTS_EXT = 1000377000,
  // Provided by VK_EXT_extended_dynamic_state2
    VK_DYNAMIC_STATE_LOGIC_OP_EXT = 1000377003,
  // Provided by VK_EXT_color_write_enable
    VK_DYNAMIC_STATE_COLOR_WRITE_ENABLE_EXT = 1000381000,
  // Provided by VK_EXT_extended_dynamic_state3
    VK_DYNAMIC_STATE_DEPTH_CLAMP_ENABLE_EXT = 1000455003,
  // Provided by VK_EXT_extended_dynamic_state3
    VK_DYNAMIC_STATE_POLYGON_MODE_EXT = 1000455004,
  // Provided by VK_EXT_extended_dynamic_state3
    VK_DYNAMIC_STATE_RASTERIZATION_SAMPLES_EXT = 1000455005,
  // Provided by VK_EXT_extended_dynamic_state3
    VK_DYNAMIC_STATE_SAMPLE_MASK_EXT = 1000455006,
  // Provided by VK_EXT_extended_dynamic_state3
    VK_DYNAMIC_STATE_ALPHA_TO_COVERAGE_ENABLE_EXT = 1000455007,
  // Provided by VK_EXT_extended_dynamic_state3
    VK_DYNAMIC_STATE_ALPHA_TO_ONE_ENABLE_EXT = 1000455008,
  // Provided by VK_EXT_extended_dynamic_state3
    VK_DYNAMIC_STATE_LOGIC_OP_ENABLE_EXT = 1000455009,
  // Provided by VK_EXT_extended_dynamic_state3
    VK_DYNAMIC_STATE_COLOR_BLEND_ENABLE_EXT = 1000455010,
  // Provided by VK_EXT_extended_dynamic_state3
    VK_DYNAMIC_STATE_COLOR_BLEND_EQUATION_EXT = 1000455011,
  // Provided by VK_EXT_extended_dynamic_state3
    VK_DYNAMIC_STATE_COLOR_WRITE_MASK_EXT = 1000455012,
  // Provided by VK_EXT_extended_dynamic_state3 with VK_KHR_maintenance2 or VK_VERSION_1_1
    VK_DYNAMIC_STATE_TESSELLATION_DOMAIN_ORIGIN_EXT = 1000455002,
  // Provided by VK_EXT_extended_dynamic_state3 with VK_EXT_transform_feedback
    VK_DYNAMIC_STATE_RASTERIZATION_STREAM_EXT = 1000455013,
  // Provided by VK_EXT_conservative_rasterization with VK_EXT_extended_dynamic_state3
    VK_DYNAMIC_STATE_CONSERVATIVE_RASTERIZATION_MODE_EXT = 1000455014,
  // Provided by VK_EXT_conservative_rasterization with VK_EXT_extended_dynamic_state3
    VK_DYNAMIC_STATE_EXTRA_PRIMITIVE_OVERESTIMATION_SIZE_EXT = 1000455015,
  // Provided by VK_EXT_depth_clip_enable with VK_EXT_extended_dynamic_state3
    VK_DYNAMIC_STATE_DEPTH_CLIP_ENABLE_EXT = 1000455016,
  // Provided by VK_EXT_extended_dynamic_state3 with VK_EXT_sample_locations
    VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_ENABLE_EXT = 1000455017,
  // Provided by VK_EXT_blend_operation_advanced with VK_EXT_extended_dynamic_state3
    VK_DYNAMIC_STATE_COLOR_BLEND_ADVANCED_EXT = 1000455018,
  // Provided by VK_EXT_extended_dynamic_state3 with VK_EXT_provoking_vertex
    VK_DYNAMIC_STATE_PROVOKING_VERTEX_MODE_EXT = 1000455019,
  // Provided by VK_EXT_extended_dynamic_state3 with VK_EXT_line_rasterization
    VK_DYNAMIC_STATE_LINE_RASTERIZATION_MODE_EXT = 1000455020,
  // Provided by VK_EXT_extended_dynamic_state3 with VK_EXT_line_rasterization
    VK_DYNAMIC_STATE_LINE_STIPPLE_ENABLE_EXT = 1000455021,
  // Provided by VK_EXT_depth_clip_control with VK_EXT_extended_dynamic_state3
    VK_DYNAMIC_STATE_DEPTH_CLIP_NEGATIVE_ONE_TO_ONE_EXT = 1000455022,
  // Provided by VK_EXT_extended_dynamic_state3 with VK_NV_clip_space_w_scaling
    VK_DYNAMIC_STATE_VIEWPORT_W_SCALING_ENABLE_NV = 1000455023,
  // Provided by VK_EXT_extended_dynamic_state3 with VK_NV_viewport_swizzle
    VK_DYNAMIC_STATE_VIEWPORT_SWIZZLE_NV = 1000455024,
  // Provided by VK_EXT_extended_dynamic_state3 with VK_NV_fragment_coverage_to_color
    VK_DYNAMIC_STATE_COVERAGE_TO_COLOR_ENABLE_NV = 1000455025,
  // Provided by VK_EXT_extended_dynamic_state3 with VK_NV_fragment_coverage_to_color
    VK_DYNAMIC_STATE_COVERAGE_TO_COLOR_LOCATION_NV = 1000455026,
  // Provided by VK_EXT_extended_dynamic_state3 with VK_NV_framebuffer_mixed_samples
    VK_DYNAMIC_STATE_COVERAGE_MODULATION_MODE_NV = 1000455027,
  // Provided by VK_EXT_extended_dynamic_state3 with VK_NV_framebuffer_mixed_samples
    VK_DYNAMIC_STATE_COVERAGE_MODULATION_TABLE_ENABLE_NV = 1000455028,
  // Provided by VK_EXT_extended_dynamic_state3 with VK_NV_framebuffer_mixed_samples
    VK_DYNAMIC_STATE_COVERAGE_MODULATION_TABLE_NV = 1000455029,
  // Provided by VK_EXT_extended_dynamic_state3 with VK_NV_shading_rate_image
    VK_DYNAMIC_STATE_SHADING_RATE_IMAGE_ENABLE_NV = 1000455030,
  // Provided by VK_EXT_extended_dynamic_state3 with VK_NV_representative_fragment_test
    VK_DYNAMIC_STATE_REPRESENTATIVE_FRAGMENT_TEST_ENABLE_NV = 1000455031,
  // Provided by VK_EXT_extended_dynamic_state3 with VK_NV_coverage_reduction_mode
    VK_DYNAMIC_STATE_COVERAGE_REDUCTION_MODE_NV = 1000455032,
  // Provided by VK_EXT_attachment_feedback_loop_dynamic_state
    VK_DYNAMIC_STATE_ATTACHMENT_FEEDBACK_LOOP_ENABLE_EXT = 1000524000,
  // Provided by VK_KHR_line_rasterization
    VK_DYNAMIC_STATE_LINE_STIPPLE_KHR = 1000259000,
  // Provided by VK_EXT_line_rasterization
    VK_DYNAMIC_STATE_LINE_STIPPLE_EXT = VK_DYNAMIC_STATE_LINE_STIPPLE_KHR,
  // Provided by VK_EXT_extended_dynamic_state
    VK_DYNAMIC_STATE_CULL_MODE_EXT = VK_DYNAMIC_STATE_CULL_MODE,
  // Provided by VK_EXT_extended_dynamic_state
    VK_DYNAMIC_STATE_FRONT_FACE_EXT = VK_DYNAMIC_STATE_FRONT_FACE,
  // Provided by VK_EXT_extended_dynamic_state
    VK_DYNAMIC_STATE_PRIMITIVE_TOPOLOGY_EXT = VK_DYNAMIC_STATE_PRIMITIVE_TOPOLOGY,
  // Provided by VK_EXT_extended_dynamic_state
    VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT_EXT = VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT,
  // Provided by VK_EXT_extended_dynamic_state
    VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT_EXT = VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT,
  // Provided by VK_EXT_extended_dynamic_state
    VK_DYNAMIC_STATE_VERTEX_INPUT_BINDING_STRIDE_EXT = VK_DYNAMIC_STATE_VERTEX_INPUT_BINDING_STRIDE,
  // Provided by VK_EXT_extended_dynamic_state
    VK_DYNAMIC_STATE_DEPTH_TEST_ENABLE_EXT = VK_DYNAMIC_STATE_DEPTH_TEST_ENABLE,
  // Provided by VK_EXT_extended_dynamic_state
    VK_DYNAMIC_STATE_DEPTH_WRITE_ENABLE_EXT = VK_DYNAMIC_STATE_DEPTH_WRITE_ENABLE,
  // Provided by VK_EXT_extended_dynamic_state
    VK_DYNAMIC_STATE_DEPTH_COMPARE_OP_EXT = VK_DYNAMIC_STATE_DEPTH_COMPARE_OP,
  // Provided by VK_EXT_extended_dynamic_state
    VK_DYNAMIC_STATE_DEPTH_BOUNDS_TEST_ENABLE_EXT = VK_DYNAMIC_STATE_DEPTH_BOUNDS_TEST_ENABLE,
  // Provided by VK_EXT_extended_dynamic_state
    VK_DYNAMIC_STATE_STENCIL_TEST_ENABLE_EXT = VK_DYNAMIC_STATE_STENCIL_TEST_ENABLE,
  // Provided by VK_EXT_extended_dynamic_state
    VK_DYNAMIC_STATE_STENCIL_OP_EXT = VK_DYNAMIC_STATE_STENCIL_OP,
  // Provided by VK_EXT_extended_dynamic_state2
    VK_DYNAMIC_STATE_RASTERIZER_DISCARD_ENABLE_EXT = VK_DYNAMIC_STATE_RASTERIZER_DISCARD_ENABLE,
  // Provided by VK_EXT_extended_dynamic_state2
    VK_DYNAMIC_STATE_DEPTH_BIAS_ENABLE_EXT = VK_DYNAMIC_STATE_DEPTH_BIAS_ENABLE,
  // Provided by VK_EXT_extended_dynamic_state2
    VK_DYNAMIC_STATE_PRIMITIVE_RESTART_ENABLE_EXT = VK_DYNAMIC_STATE_PRIMITIVE_RESTART_ENABLE,
} VkDynamicState;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_DYNAMIC_STATE_VIEWPORT` specifies that the `pViewports` state in
  [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetViewport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewport.html) before any drawing commands.
  The number of viewports used by a pipeline is still specified by the
  `viewportCount` member of
  [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html).

- `VK_DYNAMIC_STATE_SCISSOR` specifies that the `pScissors` state in
  [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetScissor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetScissor.html) before any drawing commands.
  The number of scissor rectangles used by a pipeline is still specified
  by the `scissorCount` member of
  [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html).

- `VK_DYNAMIC_STATE_LINE_WIDTH` specifies that the `lineWidth` state in
  [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetLineWidth](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineWidth.html) before any drawing
  commands that generate line primitives for the rasterizer.

- `VK_DYNAMIC_STATE_DEPTH_BIAS` specifies that any instance of
  [VkDepthBiasRepresentationInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDepthBiasRepresentationInfoEXT.html)
  included in the `pNext` chain of
  [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)
  as well as the `depthBiasConstantFactor`, `depthBiasClamp` and
  `depthBiasSlopeFactor` states in
  [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetDepthBias](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthBias.html) or
  [vkCmdSetDepthBias2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthBias2EXT.html) before any draws
  are performed with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-depthbias-enable"
  target="_blank" rel="noopener">depth bias enabled</a>.

- `VK_DYNAMIC_STATE_BLEND_CONSTANTS` specifies that the `blendConstants`
  state in
  [VkPipelineColorBlendStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetBlendConstants](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetBlendConstants.html) before any draws
  are performed with a pipeline state with
  `VkPipelineColorBlendAttachmentState` member `blendEnable` set to
  `VK_TRUE` and any of the blend functions using a constant blend color.

- `VK_DYNAMIC_STATE_DEPTH_BOUNDS` specifies that the `minDepthBounds`
  and `maxDepthBounds` states of
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetDepthBounds](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthBounds.html) before any draws are
  performed with a pipeline state with
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
  member `depthBoundsTestEnable` set to `VK_TRUE`.

- `VK_DYNAMIC_STATE_STENCIL_COMPARE_MASK` specifies that the
  `compareMask` state in
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
  for both `front` and `back` will be ignored and **must** be set
  dynamically with
  [vkCmdSetStencilCompareMask](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilCompareMask.html) before
  any draws are performed with a pipeline state with
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
  member `stencilTestEnable` set to `VK_TRUE`

- `VK_DYNAMIC_STATE_STENCIL_WRITE_MASK` specifies that the `writeMask`
  state in
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
  for both `front` and `back` will be ignored and **must** be set
  dynamically with
  [vkCmdSetStencilWriteMask](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilWriteMask.html) before any
  draws are performed with a pipeline state with
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
  member `stencilTestEnable` set to `VK_TRUE`

- `VK_DYNAMIC_STATE_STENCIL_REFERENCE` specifies that the `reference`
  state in
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
  for both `front` and `back` will be ignored and **must** be set
  dynamically with
  [vkCmdSetStencilReference](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilReference.html) before any
  draws are performed with a pipeline state with
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
  member `stencilTestEnable` set to `VK_TRUE`

- `VK_DYNAMIC_STATE_VIEWPORT_W_SCALING_NV` specifies that the
  `pViewportWScalings` state in
  [VkPipelineViewportWScalingStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportWScalingStateCreateInfoNV.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetViewportWScalingNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWScalingNV.html) before
  any draws are performed with a pipeline state with
  [VkPipelineViewportWScalingStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportWScalingStateCreateInfoNV.html)
  member `viewportScalingEnable` set to `VK_TRUE`

- `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_EXT` specifies that the
  `pDiscardRectangles` state in
  [VkPipelineDiscardRectangleStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDiscardRectangleStateCreateInfoEXT.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetDiscardRectangleEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDiscardRectangleEXT.html) before
  any draw or clear commands.

- `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_ENABLE_EXT` specifies that the
  presence of the
  [VkPipelineDiscardRectangleStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDiscardRectangleStateCreateInfoEXT.html)
  structure in the
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html)
  chain with a `discardRectangleCount` greater than zero does not
  implicitly enable discard rectangles and they **must** be enabled
  dynamically with
  [vkCmdSetDiscardRectangleEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDiscardRectangleEnableEXT.html)
  before any draw commands. This is available on implementations that
  support at least `specVersion` `2` of the
  [`VK_EXT_discard_rectangles`](VK_EXT_discard_rectangles.html)
  extension.

- `VK_DYNAMIC_STATE_DISCARD_RECTANGLE_MODE_EXT` specifies that the
  `discardRectangleMode` state in
  [VkPipelineDiscardRectangleStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDiscardRectangleStateCreateInfoEXT.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetDiscardRectangleModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDiscardRectangleModeEXT.html)
  before any draw commands. This is available on implementations that
  support at least `specVersion` `2` of the
  [`VK_EXT_discard_rectangles`](VK_EXT_discard_rectangles.html)
  extension.

- `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_EXT` specifies that the
  `sampleLocationsInfo` state in
  [VkPipelineSampleLocationsStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineSampleLocationsStateCreateInfoEXT.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetSampleLocationsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetSampleLocationsEXT.html) before
  any draw or clear commands. Enabling custom sample locations is still
  indicated by the `sampleLocationsEnable` member of
  [VkPipelineSampleLocationsStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineSampleLocationsStateCreateInfoEXT.html).

- `VK_DYNAMIC_STATE_EXCLUSIVE_SCISSOR_NV` specifies that the
  `pExclusiveScissors` state in
  [VkPipelineViewportExclusiveScissorStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportExclusiveScissorStateCreateInfoNV.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetExclusiveScissorNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetExclusiveScissorNV.html) before
  any drawing commands.

- `VK_DYNAMIC_STATE_EXCLUSIVE_SCISSOR_ENABLE_NV` specifies that the
  exclusive scissors **must** be explicitly enabled with
  [vkCmdSetExclusiveScissorEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetExclusiveScissorEnableNV.html)
  and the `exclusiveScissorCount` value in
  [VkPipelineViewportExclusiveScissorStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportExclusiveScissorStateCreateInfoNV.html)
  will not implicitly enable them. This is available on implementations
  that support at least `specVersion` `2` of the
  [`VK_NV_scissor_exclusive`](VK_NV_scissor_exclusive.html) extension.

- `VK_DYNAMIC_STATE_VIEWPORT_SHADING_RATE_PALETTE_NV` specifies that the
  `pShadingRatePalettes` state in
  [VkPipelineViewportShadingRateImageStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportShadingRateImageStateCreateInfoNV.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetViewportShadingRatePaletteNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportShadingRatePaletteNV.html)
  before any drawing commands.

- `VK_DYNAMIC_STATE_VIEWPORT_COARSE_SAMPLE_ORDER_NV` specifies that the
  coarse sample order state in
  [VkPipelineViewportCoarseSampleOrderStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportCoarseSampleOrderStateCreateInfoNV.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetCoarseSampleOrderNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoarseSampleOrderNV.html) before
  any drawing commands.

- `VK_DYNAMIC_STATE_LINE_STIPPLE_EXT` specifies that the
  `lineStippleFactor` and `lineStipplePattern` state in
  [VkPipelineRasterizationLineStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationLineStateCreateInfoKHR.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetLineStippleKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineStippleKHR.html) before any draws
  are performed with a pipeline state with
  [VkPipelineRasterizationLineStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationLineStateCreateInfoKHR.html)
  member `stippledLineEnable` set to `VK_TRUE`.

- `VK_DYNAMIC_STATE_CULL_MODE` specifies that the `cullMode` state in
  [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetCullMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCullMode.html) before any drawing commands.

- `VK_DYNAMIC_STATE_FRONT_FACE` specifies that the `frontFace` state in
  [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetFrontFace](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetFrontFace.html) before any drawing
  commands.

- `VK_DYNAMIC_STATE_PRIMITIVE_TOPOLOGY` specifies that the `topology`
  state in
  [VkPipelineInputAssemblyStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineInputAssemblyStateCreateInfo.html)
  only specifies the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#drawing-primitive-topology-class"
  target="_blank" rel="noopener">topology class</a>, and the specific
  topology order and adjacency **must** be set dynamically with
  [vkCmdSetPrimitiveTopology](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetPrimitiveTopology.html) before any
  drawing commands.

- `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT` specifies that the
  `viewportCount` and `pViewports` state in
  [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetViewportWithCount](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWithCount.html) before any
  draw call.

- `VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT` specifies that the
  `scissorCount` and `pScissors` state in
  [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetScissorWithCount](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetScissorWithCount.html) before any
  draw call.

- `VK_DYNAMIC_STATE_VERTEX_INPUT_BINDING_STRIDE` specifies that the
  `stride` state in
  [VkVertexInputBindingDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputBindingDescription.html)
  will be ignored and **must** be set dynamically with
  [vkCmdBindVertexBuffers2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindVertexBuffers2.html) before any
  draw call.

- `VK_DYNAMIC_STATE_DEPTH_TEST_ENABLE` specifies that the
  `depthTestEnable` state in
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetDepthTestEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthTestEnable.html) before any
  draw call.

- `VK_DYNAMIC_STATE_DEPTH_WRITE_ENABLE` specifies that the
  `depthWriteEnable` state in
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetDepthWriteEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthWriteEnable.html) before any
  draw call.

- `VK_DYNAMIC_STATE_DEPTH_COMPARE_OP` specifies that the
  `depthCompareOp` state in
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetDepthCompareOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthCompareOp.html) before any draw
  call.

- `VK_DYNAMIC_STATE_DEPTH_BOUNDS_TEST_ENABLE` specifies that the
  `depthBoundsTestEnable` state in
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetDepthBoundsTestEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthBoundsTestEnable.html)
  before any draw call.

- `VK_DYNAMIC_STATE_STENCIL_TEST_ENABLE` specifies that the
  `stencilTestEnable` state in
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetStencilTestEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilTestEnable.html) before any
  draw call.

- `VK_DYNAMIC_STATE_STENCIL_OP` specifies that the `failOp`, `passOp`,
  `depthFailOp`, and `compareOp` states in
  `VkPipelineDepthStencilStateCreateInfo` for both `front` and `back`
  will be ignored and **must** be set dynamically with
  [vkCmdSetStencilOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilOp.html) before any draws are
  performed with a pipeline state with
  `VkPipelineDepthStencilStateCreateInfo` member `stencilTestEnable` set
  to `VK_TRUE`

- `VK_DYNAMIC_STATE_PATCH_CONTROL_POINTS_EXT` specifies that the
  `patchControlPoints` state in
  [VkPipelineTessellationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineTessellationStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetPatchControlPointsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetPatchControlPointsEXT.html)
  before any drawing commands.

- `VK_DYNAMIC_STATE_RASTERIZER_DISCARD_ENABLE` specifies that the
  `rasterizerDiscardEnable` state in
  [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetRasterizerDiscardEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizerDiscardEnable.html)
  before any drawing commands.

- `VK_DYNAMIC_STATE_DEPTH_BIAS_ENABLE` specifies that the
  `depthBiasEnable` state in
  [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetDepthBiasEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthBiasEnable.html) before any
  drawing commands.

- `VK_DYNAMIC_STATE_LOGIC_OP_EXT` specifies that the `logicOp` state in
  [VkPipelineColorBlendStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetLogicOpEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLogicOpEXT.html) before any drawing
  commands.

- `VK_DYNAMIC_STATE_PRIMITIVE_RESTART_ENABLE` specifies that the
  `primitiveRestartEnable` state in
  [VkPipelineInputAssemblyStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineInputAssemblyStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetPrimitiveRestartEnable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetPrimitiveRestartEnable.html)
  before any drawing commands.

- `VK_DYNAMIC_STATE_FRAGMENT_SHADING_RATE_KHR` specifies that state in
  [VkPipelineFragmentShadingRateStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateStateCreateInfoKHR.html)
  and
  [VkPipelineFragmentShadingRateEnumStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineFragmentShadingRateEnumStateCreateInfoNV.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetFragmentShadingRateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetFragmentShadingRateKHR.html)
  or
  [vkCmdSetFragmentShadingRateEnumNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetFragmentShadingRateEnumNV.html)
  before any drawing commands.

- `VK_DYNAMIC_STATE_RAY_TRACING_PIPELINE_STACK_SIZE_KHR` specifies that
  the default stack size computation for the pipeline will be ignored
  and **must** be set dynamically with
  [vkCmdSetRayTracingPipelineStackSizeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRayTracingPipelineStackSizeKHR.html)
  before any ray tracing calls are performed.

- `VK_DYNAMIC_STATE_VERTEX_INPUT_EXT` specifies that the
  `pVertexInputState` state will be ignored and **must** be set
  dynamically with [vkCmdSetVertexInputEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetVertexInputEXT.html)
  before any drawing commands

- `VK_DYNAMIC_STATE_COLOR_WRITE_ENABLE_EXT` specifies that the
  `pColorWriteEnables` state in
  [VkPipelineColorWriteCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorWriteCreateInfoEXT.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetColorWriteEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorWriteEnableEXT.html) before
  any draw call.

- `VK_DYNAMIC_STATE_TESSELLATION_DOMAIN_ORIGIN_EXT` specifies that the
  `domainOrigin` state in
  [VkPipelineTessellationDomainOriginStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineTessellationDomainOriginStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetTessellationDomainOriginEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetTessellationDomainOriginEXT.html)
  before any draw call.

- `VK_DYNAMIC_STATE_DEPTH_CLAMP_ENABLE_EXT` specifies that the
  `depthClampEnable` state in
  [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetDepthClampEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthClampEnableEXT.html) before
  any draw call.

- `VK_DYNAMIC_STATE_POLYGON_MODE_EXT` specifies that the `polygonMode`
  state in
  [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetPolygonModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetPolygonModeEXT.html) before any draw
  call.

- `VK_DYNAMIC_STATE_RASTERIZATION_SAMPLES_EXT` specifies that the
  `rasterizationSamples` state in
  [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetRasterizationSamplesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizationSamplesEXT.html)
  before any draw call.

- `VK_DYNAMIC_STATE_SAMPLE_MASK_EXT` specifies that the `pSampleMask`
  state in
  [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetSampleMaskEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetSampleMaskEXT.html) before any draw
  call.

- `VK_DYNAMIC_STATE_ALPHA_TO_COVERAGE_ENABLE_EXT` specifies that the
  `alphaToCoverageEnable` state in
  [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetAlphaToCoverageEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetAlphaToCoverageEnableEXT.html)
  before any draw call.

- `VK_DYNAMIC_STATE_ALPHA_TO_ONE_ENABLE_EXT` specifies that the
  `alphaToOneEnable` state in
  [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetAlphaToOneEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetAlphaToOneEnableEXT.html) before
  any draw call.

- `VK_DYNAMIC_STATE_LOGIC_OP_ENABLE_EXT` specifies that the
  `logicOpEnable` state in
  [VkPipelineColorBlendStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateInfo.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetLogicOpEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLogicOpEnableEXT.html) before any
  draw call.

- `VK_DYNAMIC_STATE_COLOR_BLEND_ENABLE_EXT` specifies that the
  `blendEnable` state in
  [VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAttachmentState.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetColorBlendEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendEnableEXT.html) before
  any draw call.

- `VK_DYNAMIC_STATE_COLOR_BLEND_EQUATION_EXT` specifies that the
  `srcColorBlendFactor`, `dstColorBlendFactor`, `colorBlendOp`,
  `srcAlphaBlendFactor`, `dstAlphaBlendFactor`, and `alphaBlendOp`
  states in
  [VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAttachmentState.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetColorBlendEquationEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendEquationEXT.html)
  before any draw call.

- `VK_DYNAMIC_STATE_COLOR_WRITE_MASK_EXT` specifies that the
  `colorWriteMask` state in
  [VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAttachmentState.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetColorWriteMaskEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorWriteMaskEXT.html) before any
  draw call.

- `VK_DYNAMIC_STATE_RASTERIZATION_STREAM_EXT` specifies that the
  `rasterizationStream` state in
  [VkPipelineRasterizationStateStreamCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateStreamCreateInfoEXT.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetRasterizationStreamEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRasterizationStreamEXT.html)
  before any draw call.

- `VK_DYNAMIC_STATE_CONSERVATIVE_RASTERIZATION_MODE_EXT` specifies that
  the `conservativeRasterizationMode` state in
  [VkPipelineRasterizationConservativeStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationConservativeStateCreateInfoEXT.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetConservativeRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetConservativeRasterizationModeEXT.html)
  before any draw call.

- `VK_DYNAMIC_STATE_EXTRA_PRIMITIVE_OVERESTIMATION_SIZE_EXT` specifies
  that the `extraPrimitiveOverestimationSize` state in
  [VkPipelineRasterizationConservativeStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationConservativeStateCreateInfoEXT.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetExtraPrimitiveOverestimationSizeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetExtraPrimitiveOverestimationSizeEXT.html)
  before any draw call.

- `VK_DYNAMIC_STATE_DEPTH_CLIP_ENABLE_EXT` specifies that the
  `depthClipEnable` state in
  [VkPipelineRasterizationDepthClipStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationDepthClipStateCreateInfoEXT.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetDepthClipEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthClipEnableEXT.html) before
  any draw call.

- `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_ENABLE_EXT` specifies that the
  `sampleLocationsEnable` state in
  [VkPipelineSampleLocationsStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineSampleLocationsStateCreateInfoEXT.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetSampleLocationsEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetSampleLocationsEnableEXT.html)
  before any draw call.

- `VK_DYNAMIC_STATE_COLOR_BLEND_ADVANCED_EXT` specifies that the
  `colorBlendOp` state in
  [VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAttachmentState.html),
  and `srcPremultiplied`, `dstPremultiplied`, and `blendOverlap` states
  in
  [VkPipelineColorBlendAdvancedStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAdvancedStateCreateInfoEXT.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetColorBlendAdvancedEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorBlendAdvancedEXT.html)
  before any draw call.

- `VK_DYNAMIC_STATE_PROVOKING_VERTEX_MODE_EXT` specifies that the
  `provokingVertexMode` state in
  [VkPipelineRasterizationProvokingVertexStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationProvokingVertexStateCreateInfoEXT.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetProvokingVertexModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetProvokingVertexModeEXT.html)
  before any draw call.

- `VK_DYNAMIC_STATE_LINE_RASTERIZATION_MODE_EXT` specifies that the
  `lineRasterizationMode` state in
  [VkPipelineRasterizationLineStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationLineStateCreateInfoKHR.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetLineRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineRasterizationModeEXT.html)
  before any draw call.

- `VK_DYNAMIC_STATE_LINE_STIPPLE_ENABLE_EXT` specifies that the
  `stippledLineEnable` state in
  [VkPipelineRasterizationLineStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationLineStateCreateInfoKHR.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetLineStippleEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineStippleEnableEXT.html)
  before any draw call.

- `VK_DYNAMIC_STATE_DEPTH_CLIP_NEGATIVE_ONE_TO_ONE_EXT` specifies that
  the `negativeOneToOne` state in
  [VkPipelineViewportDepthClipControlCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportDepthClipControlCreateInfoEXT.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetDepthClipNegativeOneToOneEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthClipNegativeOneToOneEXT.html)
  before any draw call.

- `VK_DYNAMIC_STATE_VIEWPORT_W_SCALING_ENABLE_NV` specifies that the
  `viewportWScalingEnable` state in
  [VkPipelineViewportWScalingStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportWScalingStateCreateInfoNV.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetViewportWScalingEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWScalingEnableNV.html)
  before any draw call.

- `VK_DYNAMIC_STATE_VIEWPORT_SWIZZLE_NV` specifies that the
  `viewportCount`, and `pViewportSwizzles` states in
  [VkPipelineViewportSwizzleStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportSwizzleStateCreateInfoNV.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetViewportSwizzleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportSwizzleNV.html) before any
  draw call.

- `VK_DYNAMIC_STATE_COVERAGE_TO_COLOR_ENABLE_NV` specifies that the
  `coverageToColorEnable` state in
  [VkPipelineCoverageToColorStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageToColorStateCreateInfoNV.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetCoverageToColorEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageToColorEnableNV.html)
  before any draw call.

- `VK_DYNAMIC_STATE_COVERAGE_TO_COLOR_LOCATION_NV` specifies that the
  `coverageToColorLocation` state in
  [VkPipelineCoverageToColorStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageToColorStateCreateInfoNV.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetCoverageToColorLocationNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageToColorLocationNV.html)
  before any draw call.

- `VK_DYNAMIC_STATE_COVERAGE_MODULATION_MODE_NV` specifies that the
  `coverageModulationMode` state in
  [VkPipelineCoverageModulationStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageModulationStateCreateInfoNV.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetCoverageModulationModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageModulationModeNV.html)
  before any draw call.

- `VK_DYNAMIC_STATE_COVERAGE_MODULATION_TABLE_ENABLE_NV` specifies that
  the `coverageModulationTableEnable` state in
  [VkPipelineCoverageModulationStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageModulationStateCreateInfoNV.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetCoverageModulationTableEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageModulationTableEnableNV.html)
  before any draw call.

- `VK_DYNAMIC_STATE_COVERAGE_MODULATION_TABLE_NV` specifies that the
  `coverageModulationTableCount`, and `pCoverageModulationTable` states
  in
  [VkPipelineCoverageModulationStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageModulationStateCreateInfoNV.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetCoverageModulationTableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageModulationTableNV.html)
  before any draw call.

- `VK_DYNAMIC_STATE_SHADING_RATE_IMAGE_ENABLE_NV` specifies that the
  `shadingRateImageEnable` state in
  [VkPipelineViewportShadingRateImageStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportShadingRateImageStateCreateInfoNV.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetShadingRateImageEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetShadingRateImageEnableNV.html)
  before any draw call.

- `VK_DYNAMIC_STATE_REPRESENTATIVE_FRAGMENT_TEST_ENABLE_NV` specifies
  that the `representativeFragmentTestEnable` state in
  [VkPipelineRepresentativeFragmentTestStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRepresentativeFragmentTestStateCreateInfoNV.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetRepresentativeFragmentTestEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRepresentativeFragmentTestEnableNV.html)
  before any draw call.

- `VK_DYNAMIC_STATE_COVERAGE_REDUCTION_MODE_NV` specifies that the
  `coverageReductionMode` state in
  [VkPipelineCoverageReductionStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageReductionStateCreateInfoNV.html)
  will be ignored and **must** be set dynamically with
  [vkCmdSetCoverageReductionModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageReductionModeNV.html)
  before any draw call.

- `VK_DYNAMIC_STATE_ATTACHMENT_FEEDBACK_LOOP_ENABLE_EXT` specifies that
  the `VK_PIPELINE_CREATE_COLOR_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT` and
  `VK_PIPELINE_CREATE_DEPTH_STENCIL_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT`
  flags will be ignored and **must** be set dynamically with
  [vkCmdSetAttachmentFeedbackLoopEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetAttachmentFeedbackLoopEnableEXT.html)
  before any draw call.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDynamicState"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
