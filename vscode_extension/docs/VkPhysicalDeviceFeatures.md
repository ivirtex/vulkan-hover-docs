# VkPhysicalDeviceFeatures(3) Manual Page

## Name

VkPhysicalDeviceFeatures - Structure describing the fine-grained
features that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceFeatures` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkPhysicalDeviceFeatures {
    VkBool32    robustBufferAccess;
    VkBool32    fullDrawIndexUint32;
    VkBool32    imageCubeArray;
    VkBool32    independentBlend;
    VkBool32    geometryShader;
    VkBool32    tessellationShader;
    VkBool32    sampleRateShading;
    VkBool32    dualSrcBlend;
    VkBool32    logicOp;
    VkBool32    multiDrawIndirect;
    VkBool32    drawIndirectFirstInstance;
    VkBool32    depthClamp;
    VkBool32    depthBiasClamp;
    VkBool32    fillModeNonSolid;
    VkBool32    depthBounds;
    VkBool32    wideLines;
    VkBool32    largePoints;
    VkBool32    alphaToOne;
    VkBool32    multiViewport;
    VkBool32    samplerAnisotropy;
    VkBool32    textureCompressionETC2;
    VkBool32    textureCompressionASTC_LDR;
    VkBool32    textureCompressionBC;
    VkBool32    occlusionQueryPrecise;
    VkBool32    pipelineStatisticsQuery;
    VkBool32    vertexPipelineStoresAndAtomics;
    VkBool32    fragmentStoresAndAtomics;
    VkBool32    shaderTessellationAndGeometryPointSize;
    VkBool32    shaderImageGatherExtended;
    VkBool32    shaderStorageImageExtendedFormats;
    VkBool32    shaderStorageImageMultisample;
    VkBool32    shaderStorageImageReadWithoutFormat;
    VkBool32    shaderStorageImageWriteWithoutFormat;
    VkBool32    shaderUniformBufferArrayDynamicIndexing;
    VkBool32    shaderSampledImageArrayDynamicIndexing;
    VkBool32    shaderStorageBufferArrayDynamicIndexing;
    VkBool32    shaderStorageImageArrayDynamicIndexing;
    VkBool32    shaderClipDistance;
    VkBool32    shaderCullDistance;
    VkBool32    shaderFloat64;
    VkBool32    shaderInt64;
    VkBool32    shaderInt16;
    VkBool32    shaderResourceResidency;
    VkBool32    shaderResourceMinLod;
    VkBool32    sparseBinding;
    VkBool32    sparseResidencyBuffer;
    VkBool32    sparseResidencyImage2D;
    VkBool32    sparseResidencyImage3D;
    VkBool32    sparseResidency2Samples;
    VkBool32    sparseResidency4Samples;
    VkBool32    sparseResidency8Samples;
    VkBool32    sparseResidency16Samples;
    VkBool32    sparseResidencyAliased;
    VkBool32    variableMultisampleRate;
    VkBool32    inheritedQueries;
} VkPhysicalDeviceFeatures;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- <span id="features-robustBufferAccess"></span> `robustBufferAccess`
  specifies that accesses to buffers are bounds-checked against the
  range of the buffer descriptor (as determined by
  `VkDescriptorBufferInfo`::`range`,
  [VkBufferViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferViewCreateInfo.html)::`range`, or the
  size of the buffer). Out of bounds accesses **must** not cause
  application termination, and the effects of shader loads, stores, and
  atomics **must** conform to an implementation-dependent behavior as
  described below.

  - A buffer access is considered to be out of bounds if any of the
    following are true:

    - The pointer was formed by `OpImageTexelPointer` and the coordinate
      is less than zero or greater than or equal to the number of whole
      elements in the bound range.

    - The pointer was not formed by `OpImageTexelPointer` and the object
      pointed to is not wholly contained within the bound range. This
      includes accesses performed via *variable pointers* where the
      buffer descriptor being accessed cannot be statically determined.
      Uninitialized pointers and pointers equal to `OpConstantNull` are
      treated as pointing to a zero-sized object, so all accesses
      through such pointers are considered to be out of bounds. Buffer
      accesses through buffer device addresses are not bounds-checked.

    - If the <a
      href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-cooperativeMatrixRobustBufferAccess-NV"
      target="_blank"
      rel="noopener"><code>VkPhysicalDeviceCooperativeMatrixFeaturesNV</code>::<code>cooperativeMatrixRobustBufferAccess</code></a>
      feature is not enabled, then accesses using
      `OpCooperativeMatrixLoadNV` and `OpCooperativeMatrixStoreNV`
      **may** not be bounds-checked.

    - If the <a
      href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-cooperativeMatrixRobustBufferAccess"
      target="_blank"
      rel="noopener"><code>VkPhysicalDeviceCooperativeMatrixFeaturesKHR</code>::<code>cooperativeMatrixRobustBufferAccess</code></a>
      feature is not enabled, then accesses using
      `OpCooperativeMatrixLoadKHR` and `OpCooperativeMatrixStoreKHR`
      **may** not be bounds-checked.

      <table>
      <colgroup>
      <col style="width: 50%" />
      <col style="width: 50%" />
      </colgroup>
      <tbody>
      <tr>
      <td class="icon"><em></em></td>
      <td class="content">Note
      <p>If a SPIR-V <code>OpLoad</code> instruction loads a structure and the
      tail end of the structure is out of bounds, then all members of the
      structure are considered out of bounds even if the members at the end
      are not statically used.</p></td>
      </tr>
      </tbody>
      </table>

## <a href="#_description" class="anchor"></a>Description

- If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustBufferAccess2"
  target="_blank" rel="noopener"><code>robustBufferAccess2</code></a> is
  not enabled and any buffer access is determined to be out of bounds,
  then any other access of the same type (load, store, or atomic) to the
  same buffer that accesses an address less than 16 bytes away from the
  out of bounds address **may** also be considered out of bounds.

- If the access is a load that reads from the same memory locations as a
  prior store in the same shader invocation, with no other intervening
  accesses to the same memory locations in that shader invocation, then
  the result of the load **may** be the value stored by the store
  instruction, even if the access is out of bounds. If the load is
  `Volatile`, then an out of bounds load **must** return the appropriate
  out of bounds value.

  - Accesses to descriptors written with a
    [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) resource or view are not
    considered to be out of bounds. Instead, each type of descriptor
    access defines a specific behavior for accesses to a null
    descriptor.

  - Out-of-bounds buffer loads will return any of the following values:

- If the access is to a uniform buffer and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustBufferAccess2"
  target="_blank" rel="noopener"><code>robustBufferAccess2</code></a> is
  enabled, loads of offsets between the end of the descriptor range and
  the end of the descriptor range rounded up to a multiple of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-robustUniformBufferAccessSizeAlignment"
  target="_blank"
  rel="noopener"><code>robustUniformBufferAccessSizeAlignment</code></a>
  bytes **must** return either zero values or the contents of the memory
  at the offset being loaded. Loads of offsets past the descriptor range
  rounded up to a multiple of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-robustUniformBufferAccessSizeAlignment"
  target="_blank"
  rel="noopener"><code>robustUniformBufferAccessSizeAlignment</code></a>
  bytes **must** return zero values.

- If the access is to a storage buffer and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustBufferAccess2"
  target="_blank" rel="noopener"><code>robustBufferAccess2</code></a> is
  enabled, loads of offsets between the end of the descriptor range and
  the end of the descriptor range rounded up to a multiple of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-robustStorageBufferAccessSizeAlignment"
  target="_blank"
  rel="noopener"><code>robustStorageBufferAccessSizeAlignment</code></a>
  bytes **must** return either zero values or the contents of the memory
  at the offset being loaded. Loads of offsets past the descriptor range
  rounded up to a multiple of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-robustStorageBufferAccessSizeAlignment"
  target="_blank"
  rel="noopener"><code>robustStorageBufferAccessSizeAlignment</code></a>
  bytes **must** return zero values. Similarly, stores to addresses
  between the end of the descriptor range and the end of the descriptor
  range rounded up to a multiple of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-robustStorageBufferAccessSizeAlignment"
  target="_blank"
  rel="noopener"><code>robustStorageBufferAccessSizeAlignment</code></a>
  bytes **may** be discarded.

- Non-atomic accesses to storage buffers that are a multiple of 32 bits
  **may** be decomposed into 32-bit accesses that are individually
  bounds-checked.

- If the access is to an index buffer and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustBufferAccess2"
  target="_blank" rel="noopener"><code>robustBufferAccess2</code></a> is
  enabled, zero values **must** be returned.

- If the access is to a uniform texel buffer or storage texel buffer and
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustBufferAccess2"
  target="_blank" rel="noopener"><code>robustBufferAccess2</code></a> is
  enabled, zero values **must** be returned, and then <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-conversion-to-rgba"
  target="_blank" rel="noopener">Conversion to RGBA</a> is applied based
  on the buffer view’s format.

- Values from anywhere within the memory range(s) bound to the buffer
  (possibly including bytes of memory past the end of the buffer, up to
  the end of the bound range).

- Zero values, or (0,0,0,x) vectors for vector reads where x is a valid
  value represented in the type of the vector components and **may** be
  any of:

  - 0, 1, or the maximum representable positive integer value, for
    signed or unsigned integer components

  - 0.0 or 1.0, for floating-point components

    - Out-of-bounds writes **may** modify values within the memory
      range(s) bound to the buffer, but **must** not modify any other
      memory.

- If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustBufferAccess2"
  target="_blank" rel="noopener"><code>robustBufferAccess2</code></a> is
  enabled, out of bounds writes **must** not modify any memory.

  - Out-of-bounds atomics **may** modify values within the memory
    range(s) bound to the buffer, but **must** not modify any other
    memory, and return an undefined value.

- If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustBufferAccess2"
  target="_blank" rel="noopener"><code>robustBufferAccess2</code></a> is
  enabled, out of bounds atomics **must** not modify any memory, and
  return an undefined value.

  - If <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustBufferAccess2"
    target="_blank" rel="noopener"><code>robustBufferAccess2</code></a>
    is disabled, vertex input attributes are considered out of bounds if
    the offset of the attribute in the bound vertex buffer range plus
    the size of the attribute is greater than either:

- `vertexBufferRangeSize`, if `bindingStride` == 0; or

- (`vertexBufferRangeSize` - (`vertexBufferRangeSize` %
  `bindingStride`))

  where `vertexBufferRangeSize` is the byte size of the memory range
  bound to the vertex buffer binding and `bindingStride` is the byte
  stride of the corresponding vertex input binding. Further, if any
  vertex input attribute using a specific vertex input binding is out of
  bounds, then all vertex input attributes using that vertex input
  binding for that vertex shader invocation are considered out of
  bounds.

- If a vertex input attribute is out of bounds, it will be assigned one
  of the following values:

  - Values from anywhere within the memory range(s) bound to the buffer,
    converted according to the format of the attribute.

  - Zero values, format converted according to the format of the
    attribute.

  - Zero values, or (0,0,0,x) vectors, as described above.

    - If <a
      href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustBufferAccess2"
      target="_blank" rel="noopener"><code>robustBufferAccess2</code></a>
      is enabled, vertex input attributes are considered out of bounds
      if the offset of the attribute in the bound vertex buffer range
      plus the size of the attribute is greater than the byte size of
      the memory range bound to the vertex buffer binding.

- If a vertex input attribute is out of bounds, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fxvertex-input-extraction"
  target="_blank" rel="noopener">raw data</a> extracted are zero values,
  and missing G, B, or A components are [filled with
  (0,0,1)](#fxvertex-input-extraction).

  - If `robustBufferAccess` is not enabled, applications **must** not
    perform out of bounds accesses except under the conditions enabled
    by the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelineRobustness"
    target="_blank" rel="noopener"><code>pipelineRobustness</code></a>
    feature .

    - <span id="features-fullDrawIndexUint32"></span>
      `fullDrawIndexUint32` specifies the full 32-bit range of indices
      is supported for indexed draw calls when using a
      [VkIndexType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndexType.html) of `VK_INDEX_TYPE_UINT32`.
      `maxDrawIndexedIndexValue` is the maximum index value that **may**
      be used (aside from the primitive restart index, which is always
      2<sup>32</sup>-1 when the [VkIndexType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndexType.html) is
      `VK_INDEX_TYPE_UINT32`). If this feature is supported,
      `maxDrawIndexedIndexValue` **must** be 2<sup>32</sup>-1; otherwise
      it **must** be no smaller than 2<sup>24</sup>-1. See <a
      href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxDrawIndexedIndexValue"
      target="_blank" rel="noopener"><code>maxDrawIndexedIndexValue</code></a>.

    - <span id="features-imageCubeArray"></span> `imageCubeArray`
      specifies whether image views with a
      [VkImageViewType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewType.html) of
      `VK_IMAGE_VIEW_TYPE_CUBE_ARRAY` **can** be created, and that the
      corresponding `SampledCubeArray` and `ImageCubeArray` SPIR-V
      capabilities **can** be used in shader code.

    - <span id="features-independentBlend"></span> `independentBlend`
      specifies whether the `VkPipelineColorBlendAttachmentState`
      settings are controlled independently per-attachment. If this
      feature is not enabled, the `VkPipelineColorBlendAttachmentState`
      settings for all color attachments **must** be identical.
      Otherwise, a different `VkPipelineColorBlendAttachmentState`
      **can** be provided for each bound color attachment.

    - <span id="features-geometryShader"></span> `geometryShader`
      specifies whether geometry shaders are supported. If this feature
      is not enabled, the `VK_SHADER_STAGE_GEOMETRY_BIT` and
      `VK_PIPELINE_STAGE_GEOMETRY_SHADER_BIT` enum values **must** not
      be used. This also specifies whether shader modules **can**
      declare the `Geometry` capability.

    - <span id="features-tessellationShader"></span>
      `tessellationShader` specifies whether tessellation control and
      evaluation shaders are supported. If this feature is not enabled,
      the `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT`,
      `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`,
      `VK_PIPELINE_STAGE_TESSELLATION_CONTROL_SHADER_BIT`,
      `VK_PIPELINE_STAGE_TESSELLATION_EVALUATION_SHADER_BIT`, and
      `VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_STATE_CREATE_INFO` enum
      values **must** not be used. This also specifies whether shader
      modules **can** declare the `Tessellation` capability.

    - <span id="features-sampleRateShading"></span> `sampleRateShading`
      specifies whether <a
      href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-sampleshading"
      target="_blank" rel="noopener">Sample Shading</a> and multisample
      interpolation are supported. If this feature is not enabled, the
      `sampleShadingEnable` member of the
      [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)
      structure **must** be set to `VK_FALSE` and the `minSampleShading`
      member is ignored. This also specifies whether shader modules
      **can** declare the `SampleRateShading` capability.

    - <span id="features-dualSrcBlend"></span> `dualSrcBlend` specifies
      whether blend operations which take two sources are supported. If
      this feature is not enabled, the `VK_BLEND_FACTOR_SRC1_COLOR`,
      `VK_BLEND_FACTOR_ONE_MINUS_SRC1_COLOR`,
      `VK_BLEND_FACTOR_SRC1_ALPHA`, and
      `VK_BLEND_FACTOR_ONE_MINUS_SRC1_ALPHA` enum values **must** not be
      used as source or destination blending factors. See <a
      href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-dsb"
      class="bare" target="_blank"
      rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-dsb</a>.

    - <span id="features-logicOp"></span> `logicOp` specifies whether
      logic operations are supported. If this feature is not enabled,
      the `logicOpEnable` member of the
      [VkPipelineColorBlendStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateInfo.html)
      structure **must** be set to `VK_FALSE`, and the `logicOp` member
      is ignored.

    - <span id="features-multiDrawIndirect"></span> `multiDrawIndirect`
      specifies whether multiple draw indirect is supported. If this
      feature is not enabled, the `drawCount` parameter to the
      `vkCmdDrawIndirect` and `vkCmdDrawIndexedIndirect` commands
      **must** be 0 or 1. The `maxDrawIndirectCount` member of the
      `VkPhysicalDeviceLimits` structure **must** also be 1 if this
      feature is not supported. See <a
      href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxDrawIndirectCount"
      target="_blank" rel="noopener"><code>maxDrawIndirectCount</code></a>.

    - <span id="features-drawIndirectFirstInstance"></span>
      `drawIndirectFirstInstance` specifies whether indirect drawing
      calls support the `firstInstance` parameter. If this feature is
      not enabled, the `firstInstance` member of all
      `VkDrawIndirectCommand` and `VkDrawIndexedIndirectCommand`
      structures that are provided to the `vkCmdDrawIndirect` and
      `vkCmdDrawIndexedIndirect` commands **must** be 0.

    - <span id="features-depthClamp"></span> `depthClamp` specifies
      whether depth clamping is supported. If this feature is not
      enabled, the `depthClampEnable` member of the
      [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)
      structure **must** be set to `VK_FALSE`. Otherwise, setting
      `depthClampEnable` to `VK_TRUE` will enable depth clamping.

    - <span id="features-depthBiasClamp"></span> `depthBiasClamp`
      specifies whether depth bias clamping is supported. If this
      feature is not enabled, the `depthBiasClamp` member of the
      [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)
      structure **must** be set to 0.0 unless the
      `VK_DYNAMIC_STATE_DEPTH_BIAS` dynamic state is enabled, and the
      `depthBiasClamp` parameter to `vkCmdSetDepthBias` **must** be set
      to 0.0.

    - <span id="features-fillModeNonSolid"></span> `fillModeNonSolid`
      specifies whether point and wireframe fill modes are supported. If
      this feature is not enabled, the `VK_POLYGON_MODE_POINT` and
      `VK_POLYGON_MODE_LINE` enum values **must** not be used.

    - <span id="features-depthBounds"></span> `depthBounds` specifies
      whether depth bounds tests are supported. If this feature is not
      enabled, the `depthBoundsTestEnable` member of the
      [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
      structure **must** be set to `VK_FALSE`. When
      `depthBoundsTestEnable` is set to `VK_FALSE`, the `minDepthBounds`
      and `maxDepthBounds` members of the
      [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)
      structure are ignored.

    - <span id="features-wideLines"></span> `wideLines` specifies
      whether lines with width other than 1.0 are supported. If this
      feature is not enabled, the `lineWidth` member of the
      [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)
      structure **must** be set to 1.0 unless the
      `VK_DYNAMIC_STATE_LINE_WIDTH` dynamic state is enabled, and the
      `lineWidth` parameter to `vkCmdSetLineWidth` **must** be set to
      1.0. When this feature is supported, the range and granularity of
      supported line widths are indicated by the `lineWidthRange` and
      `lineWidthGranularity` members of the `VkPhysicalDeviceLimits`
      structure, respectively.

    - <span id="features-largePoints"></span> `largePoints` specifies
      whether points with size greater than 1.0 are supported. If this
      feature is not enabled, only a point size of 1.0 written by a
      shader is supported. The range and granularity of supported point
      sizes are indicated by the `pointSizeRange` and
      `pointSizeGranularity` members of the `VkPhysicalDeviceLimits`
      structure, respectively.

    - <span id="features-alphaToOne"></span> `alphaToOne` specifies
      whether the implementation is able to replace the alpha value of
      the fragment shader color output in the <a
      href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-covg"
      target="_blank" rel="noopener">Multisample Coverage</a> fragment
      operation. If this feature is not enabled, then the
      `alphaToOneEnable` member of the
      [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)
      structure **must** be set to `VK_FALSE`. Otherwise setting
      `alphaToOneEnable` to `VK_TRUE` will enable alpha-to-one behavior.

    - <span id="features-multiViewport"></span> `multiViewport`
      specifies whether more than one viewport is supported. If this
      feature is not enabled:

  - The `viewportCount` and `scissorCount` members of the
    [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html)
    structure **must** be set to 1.

  - The `firstViewport` and `viewportCount` parameters to the
    `vkCmdSetViewport` command **must** be set to 0 and 1, respectively.

  - The `firstScissor` and `scissorCount` parameters to the
    `vkCmdSetScissor` command **must** be set to 0 and 1, respectively.

  - The `exclusiveScissorCount` member of the
    [VkPipelineViewportExclusiveScissorStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportExclusiveScissorStateCreateInfoNV.html)
    structure **must** be set to 0 or 1.

  - The `firstExclusiveScissor` and `exclusiveScissorCount` parameters
    to the `vkCmdSetExclusiveScissorNV` command **must** be set to 0 and
    1, respectively.

    - <span id="features-samplerAnisotropy"></span> `samplerAnisotropy`
      specifies whether anisotropic filtering is supported. If this
      feature is not enabled, the `anisotropyEnable` member of the
      [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html) structure **must**
      be `VK_FALSE`.

    - <span id="features-textureCompressionETC2"></span>
      `textureCompressionETC2` specifies whether all of the ETC2 and EAC
      compressed texture formats are supported. If this feature is
      enabled, then the `VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT`,
      `VK_FORMAT_FEATURE_BLIT_SRC_BIT` and
      `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT` features
      **must** be supported in `optimalTilingFeatures` for the following
      formats:

  - `VK_FORMAT_ETC2_R8G8B8_UNORM_BLOCK`

  - `VK_FORMAT_ETC2_R8G8B8_SRGB_BLOCK`

  - `VK_FORMAT_ETC2_R8G8B8A1_UNORM_BLOCK`

  - `VK_FORMAT_ETC2_R8G8B8A1_SRGB_BLOCK`

  - `VK_FORMAT_ETC2_R8G8B8A8_UNORM_BLOCK`

  - `VK_FORMAT_ETC2_R8G8B8A8_SRGB_BLOCK`

  - `VK_FORMAT_EAC_R11_UNORM_BLOCK`

  - `VK_FORMAT_EAC_R11_SNORM_BLOCK`

  - `VK_FORMAT_EAC_R11G11_UNORM_BLOCK`

  - `VK_FORMAT_EAC_R11G11_SNORM_BLOCK`

    To query for additional properties, or if the feature is not
    enabled,
    [vkGetPhysicalDeviceFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties.html)
    and
    [vkGetPhysicalDeviceImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties.html)
    **can** be used to check for supported properties of individual
    formats as normal.

    - <span id="features-textureCompressionASTC_LDR"></span>
      `textureCompressionASTC_LDR` specifies whether all of the ASTC LDR
      compressed texture formats are supported. If this feature is
      enabled, then the `VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT`,
      `VK_FORMAT_FEATURE_BLIT_SRC_BIT` and
      `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT` features
      **must** be supported in `optimalTilingFeatures` for the following
      formats:

  - `VK_FORMAT_ASTC_4x4_UNORM_BLOCK`

  - `VK_FORMAT_ASTC_4x4_SRGB_BLOCK`

  - `VK_FORMAT_ASTC_5x4_UNORM_BLOCK`

  - `VK_FORMAT_ASTC_5x4_SRGB_BLOCK`

  - `VK_FORMAT_ASTC_5x5_UNORM_BLOCK`

  - `VK_FORMAT_ASTC_5x5_SRGB_BLOCK`

  - `VK_FORMAT_ASTC_6x5_UNORM_BLOCK`

  - `VK_FORMAT_ASTC_6x5_SRGB_BLOCK`

  - `VK_FORMAT_ASTC_6x6_UNORM_BLOCK`

  - `VK_FORMAT_ASTC_6x6_SRGB_BLOCK`

  - `VK_FORMAT_ASTC_8x5_UNORM_BLOCK`

  - `VK_FORMAT_ASTC_8x5_SRGB_BLOCK`

  - `VK_FORMAT_ASTC_8x6_UNORM_BLOCK`

  - `VK_FORMAT_ASTC_8x6_SRGB_BLOCK`

  - `VK_FORMAT_ASTC_8x8_UNORM_BLOCK`

  - `VK_FORMAT_ASTC_8x8_SRGB_BLOCK`

  - `VK_FORMAT_ASTC_10x5_UNORM_BLOCK`

  - `VK_FORMAT_ASTC_10x5_SRGB_BLOCK`

  - `VK_FORMAT_ASTC_10x6_UNORM_BLOCK`

  - `VK_FORMAT_ASTC_10x6_SRGB_BLOCK`

  - `VK_FORMAT_ASTC_10x8_UNORM_BLOCK`

  - `VK_FORMAT_ASTC_10x8_SRGB_BLOCK`

  - `VK_FORMAT_ASTC_10x10_UNORM_BLOCK`

  - `VK_FORMAT_ASTC_10x10_SRGB_BLOCK`

  - `VK_FORMAT_ASTC_12x10_UNORM_BLOCK`

  - `VK_FORMAT_ASTC_12x10_SRGB_BLOCK`

  - `VK_FORMAT_ASTC_12x12_UNORM_BLOCK`

  - `VK_FORMAT_ASTC_12x12_SRGB_BLOCK`

    To query for additional properties, or if the feature is not
    enabled,
    [vkGetPhysicalDeviceFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties.html)
    and
    [vkGetPhysicalDeviceImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties.html)
    **can** be used to check for supported properties of individual
    formats as normal.

    - <span id="features-textureCompressionBC"></span>
      `textureCompressionBC` specifies whether all of the BC compressed
      texture formats are supported. If this feature is enabled, then
      the `VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT`,
      `VK_FORMAT_FEATURE_BLIT_SRC_BIT` and
      `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT` features
      **must** be supported in `optimalTilingFeatures` for the following
      formats:

  - `VK_FORMAT_BC1_RGB_UNORM_BLOCK`

  - `VK_FORMAT_BC1_RGB_SRGB_BLOCK`

  - `VK_FORMAT_BC1_RGBA_UNORM_BLOCK`

  - `VK_FORMAT_BC1_RGBA_SRGB_BLOCK`

  - `VK_FORMAT_BC2_UNORM_BLOCK`

  - `VK_FORMAT_BC2_SRGB_BLOCK`

  - `VK_FORMAT_BC3_UNORM_BLOCK`

  - `VK_FORMAT_BC3_SRGB_BLOCK`

  - `VK_FORMAT_BC4_UNORM_BLOCK`

  - `VK_FORMAT_BC4_SNORM_BLOCK`

  - `VK_FORMAT_BC5_UNORM_BLOCK`

  - `VK_FORMAT_BC5_SNORM_BLOCK`

  - `VK_FORMAT_BC6H_UFLOAT_BLOCK`

  - `VK_FORMAT_BC6H_SFLOAT_BLOCK`

  - `VK_FORMAT_BC7_UNORM_BLOCK`

  - `VK_FORMAT_BC7_SRGB_BLOCK`

    To query for additional properties, or if the feature is not
    enabled,
    [vkGetPhysicalDeviceFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties.html)
    and
    [vkGetPhysicalDeviceImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties.html)
    **can** be used to check for supported properties of individual
    formats as normal.

    - <span id="features-occlusionQueryPrecise"></span>
      `occlusionQueryPrecise` specifies whether occlusion queries
      returning actual sample counts are supported. Occlusion queries
      are created in a `VkQueryPool` by specifying the `queryType` of
      `VK_QUERY_TYPE_OCCLUSION` in the
      [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfo.html) structure
      which is passed to `vkCreateQueryPool`. If this feature is
      enabled, queries of this type **can** enable
      `VK_QUERY_CONTROL_PRECISE_BIT` in the `flags` parameter to
      `vkCmdBeginQuery`. If this feature is not supported, the
      implementation supports only boolean occlusion queries. When any
      samples are passed, boolean queries will return a non-zero result
      value, otherwise a result value of zero is returned. When this
      feature is enabled and `VK_QUERY_CONTROL_PRECISE_BIT` is set,
      occlusion queries will report the actual number of samples passed.

    - <span id="features-pipelineStatisticsQuery"></span>
      `pipelineStatisticsQuery` specifies whether the pipeline
      statistics queries are supported. If this feature is not enabled,
      queries of type `VK_QUERY_TYPE_PIPELINE_STATISTICS` **cannot** be
      created, and none of the
      [VkQueryPipelineStatisticFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPipelineStatisticFlagBits.html)
      bits **can** be set in the `pipelineStatistics` member of the
      [VkQueryPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateInfo.html) structure.

    - <span id="features-vertexPipelineStoresAndAtomics"></span>
      `vertexPipelineStoresAndAtomics` specifies whether storage buffers
      and images support stores and atomic operations in the vertex,
      tessellation, and geometry shader stages. If this feature is not
      enabled, all storage image, storage texel buffer, and storage
      buffer variables used by these stages in shader modules **must**
      be decorated with the `NonWritable` decoration (or the `readonly`
      memory qualifier in GLSL).

    - <span id="features-fragmentStoresAndAtomics"></span>
      `fragmentStoresAndAtomics` specifies whether storage buffers and
      images support stores and atomic operations in the fragment shader
      stage. If this feature is not enabled, all storage image, storage
      texel buffer, and storage buffer variables used by the fragment
      stage in shader modules **must** be decorated with the
      `NonWritable` decoration (or the `readonly` memory qualifier in
      GLSL).

    - <span id="features-shaderTessellationAndGeometryPointSize"></span>
      `shaderTessellationAndGeometryPointSize` specifies whether the
      `PointSize` built-in decoration is available in the tessellation
      control, tessellation evaluation, and geometry shader stages. If
      this feature is not enabled, members decorated with the
      `PointSize` built-in decoration **must** not be read from or
      written to and all points written from a tessellation or geometry
      shader will have a size of 1.0. This also specifies whether shader
      modules **can** declare the `TessellationPointSize` capability for
      tessellation control and evaluation shaders, or if the shader
      modules **can** declare the `GeometryPointSize` capability for
      geometry shaders. An implementation supporting this feature
      **must** also support one or both of the <a
      href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-tessellationShader"
      target="_blank" rel="noopener"><code>tessellationShader</code></a>
      or <a
      href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-geometryShader"
      target="_blank" rel="noopener"><code>geometryShader</code></a>
      features.

    - <span id="features-shaderImageGatherExtended"></span>
      `shaderImageGatherExtended` specifies whether the extended set of
      image gather instructions are available in shader code. If this
      feature is not enabled, the `OpImage*Gather` instructions do not
      support the `Offset` and `ConstOffsets` operands. This also
      specifies whether shader modules **can** declare the
      `ImageGatherExtended` capability.

    - <span id="features-shaderStorageImageExtendedFormats"></span>
      `shaderStorageImageExtendedFormats` specifies whether all the
      “storage image extended formats” below are supported; if this
      feature is supported, then the
      `VK_FORMAT_FEATURE_STORAGE_IMAGE_BIT` **must** be supported in
      `optimalTilingFeatures` for the following formats:

  - `VK_FORMAT_R16G16_SFLOAT`

  - `VK_FORMAT_B10G11R11_UFLOAT_PACK32`

  - `VK_FORMAT_R16_SFLOAT`

  - `VK_FORMAT_R16G16B16A16_UNORM`

  - `VK_FORMAT_A2B10G10R10_UNORM_PACK32`

  - `VK_FORMAT_R16G16_UNORM`

  - `VK_FORMAT_R8G8_UNORM`

  - `VK_FORMAT_R16_UNORM`

  - `VK_FORMAT_R8_UNORM`

  - `VK_FORMAT_R16G16B16A16_SNORM`

  - `VK_FORMAT_R16G16_SNORM`

  - `VK_FORMAT_R8G8_SNORM`

  - `VK_FORMAT_R16_SNORM`

  - `VK_FORMAT_R8_SNORM`

  - `VK_FORMAT_R16G16_SINT`

  - `VK_FORMAT_R8G8_SINT`

  - `VK_FORMAT_R16_SINT`

  - `VK_FORMAT_R8_SINT`

  - `VK_FORMAT_A2B10G10R10_UINT_PACK32`

  - `VK_FORMAT_R16G16_UINT`

  - `VK_FORMAT_R8G8_UINT`

  - `VK_FORMAT_R16_UINT`

  - `VK_FORMAT_R8_UINT`

    <table>
    <colgroup>
    <col style="width: 50%" />
    <col style="width: 50%" />
    </colgroup>
    <tbody>
    <tr>
    <td class="icon"><em></em></td>
    <td class="content">Note
    <p><code>shaderStorageImageExtendedFormats</code> feature only adds a
    guarantee of format support, which is specified for the whole physical
    device. Therefore enabling or disabling the feature via <a
    href="vkCreateDevice.html">vkCreateDevice</a> has no practical
    effect.</p>
    <p>To query for additional properties, or if the feature is not
    supported, <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties.html">vkGetPhysicalDeviceFormatProperties</a>
    and <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties.html">vkGetPhysicalDeviceImageFormatProperties</a>
    <strong>can</strong> be used to check for supported properties of
    individual formats, as usual rules allow.</p>
    <p><code>VK_FORMAT_R32G32_UINT</code>,
    <code>VK_FORMAT_R32G32_SINT</code>, and
    <code>VK_FORMAT_R32G32_SFLOAT</code> from
    <code>StorageImageExtendedFormats</code> SPIR-V capability, are already
    covered by core Vulkan <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-mandatory-features-32bit"
    target="_blank" rel="noopener">mandatory format support</a>.</p></td>
    </tr>
    </tbody>
    </table>

    - <span id="features-shaderStorageImageMultisample"></span>
      `shaderStorageImageMultisample` specifies whether multisampled
      storage images are supported. If this feature is not enabled,
      images that are created with a `usage` that includes
      `VK_IMAGE_USAGE_STORAGE_BIT` **must** be created with `samples`
      equal to `VK_SAMPLE_COUNT_1_BIT`. This also specifies whether
      shader modules **can** declare the `StorageImageMultisample` and
      `ImageMSArray` capabilities.

    - <span id="features-shaderStorageImageReadWithoutFormat"></span>
      `shaderStorageImageReadWithoutFormat` specifies whether storage
      images and storage texel buffers require a format qualifier to be
      specified when reading. `shaderStorageImageReadWithoutFormat`
      applies only to formats listed in the <a
      href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-without-shader-storage-format"
      target="_blank" rel="noopener">storage without format</a> list.

    - <span id="features-shaderStorageImageWriteWithoutFormat"></span>
      `shaderStorageImageWriteWithoutFormat` specifies whether storage
      images and storage texel buffers require a format qualifier to be
      specified when writing. `shaderStorageImageWriteWithoutFormat`
      applies only to formats listed in the <a
      href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-without-shader-storage-format"
      target="_blank" rel="noopener">storage without format</a> list.

    - <span id="features-shaderUniformBufferArrayDynamicIndexing"></span>
      `shaderUniformBufferArrayDynamicIndexing` specifies whether arrays
      of uniform buffers **can** be indexed by *dynamically uniform*
      integer expressions in shader code. If this feature is not
      enabled, resources with a descriptor type of
      `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER` or
      `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` **must** be indexed
      only by constant integral expressions when aggregated into arrays
      in shader code. This also specifies whether shader modules **can**
      declare the `UniformBufferArrayDynamicIndexing` capability.

    - <span id="features-shaderSampledImageArrayDynamicIndexing"></span>
      `shaderSampledImageArrayDynamicIndexing` specifies whether arrays
      of samplers or sampled images **can** be indexed by dynamically
      uniform integer expressions in shader code. If this feature is not
      enabled, resources with a descriptor type of
      `VK_DESCRIPTOR_TYPE_SAMPLER`,
      `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, or
      `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE` **must** be indexed only by
      constant integral expressions when aggregated into arrays in
      shader code. This also specifies whether shader modules **can**
      declare the `SampledImageArrayDynamicIndexing` capability.

    - <span id="features-shaderStorageBufferArrayDynamicIndexing"></span>
      `shaderStorageBufferArrayDynamicIndexing` specifies whether arrays
      of storage buffers **can** be indexed by dynamically uniform
      integer expressions in shader code. If this feature is not
      enabled, resources with a descriptor type of
      `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER` or
      `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` **must** be indexed
      only by constant integral expressions when aggregated into arrays
      in shader code. This also specifies whether shader modules **can**
      declare the `StorageBufferArrayDynamicIndexing` capability.

    - <span id="features-shaderStorageImageArrayDynamicIndexing"></span>
      `shaderStorageImageArrayDynamicIndexing` specifies whether arrays
      of storage images **can** be indexed by dynamically uniform
      integer expressions in shader code. If this feature is not
      enabled, resources with a descriptor type of
      `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE` **must** be indexed only by
      constant integral expressions when aggregated into arrays in
      shader code. This also specifies whether shader modules **can**
      declare the `StorageImageArrayDynamicIndexing` capability.

    - <span id="features-shaderClipDistance"></span>
      `shaderClipDistance` specifies whether clip distances are
      supported in shader code. If this feature is not enabled, any
      members decorated with the `ClipDistance` built-in decoration
      **must** not be read from or written to in shader modules. This
      also specifies whether shader modules **can** declare the
      `ClipDistance` capability.

    - <span id="features-shaderCullDistance"></span>
      `shaderCullDistance` specifies whether cull distances are
      supported in shader code. If this feature is not enabled, any
      members decorated with the `CullDistance` built-in decoration
      **must** not be read from or written to in shader modules. This
      also specifies whether shader modules **can** declare the
      `CullDistance` capability.

    - <span id="features-shaderFloat64"></span> `shaderFloat64`
      specifies whether 64-bit floats (doubles) are supported in shader
      code. If this feature is not enabled, 64-bit floating-point types
      **must** not be used in shader code. This also specifies whether
      shader modules **can** declare the `Float64` capability. Declaring
      and using 64-bit floats is enabled for all storage classes that
      SPIR-V allows with the `Float64` capability.

    - <span id="features-shaderInt64"></span> `shaderInt64` specifies
      whether 64-bit integers (signed and unsigned) are supported in
      shader code. If this feature is not enabled, 64-bit integer types
      **must** not be used in shader code. This also specifies whether
      shader modules **can** declare the `Int64` capability. Declaring
      and using 64-bit integers is enabled for all storage classes that
      SPIR-V allows with the `Int64` capability.

    - <span id="features-shaderInt16"></span> `shaderInt16` specifies
      whether 16-bit integers (signed and unsigned) are supported in
      shader code. If this feature is not enabled, 16-bit integer types
      **must** not be used in shader code. This also specifies whether
      shader modules **can** declare the `Int16` capability. However,
      this only enables a subset of the storage classes that SPIR-V
      allows for the `Int16` SPIR-V capability: Declaring and using
      16-bit integers in the `Private`, `Workgroup` (for non-Block
      variables), and `Function` storage classes is enabled, while
      declaring them in the interface storage classes (e.g.,
      `UniformConstant`, `Uniform`, `StorageBuffer`, `Input`, `Output`,
      and `PushConstant`) is not enabled.

    - <span id="features-shaderResourceResidency"></span>
      `shaderResourceResidency` specifies whether image operations that
      return resource residency information are supported in shader
      code. If this feature is not enabled, the `OpImageSparse*`
      instructions **must** not be used in shader code. This also
      specifies whether shader modules **can** declare the
      `SparseResidency` capability. The feature requires at least one of
      the `sparseResidency*` features to be supported.

    - <span id="features-shaderResourceMinLod"></span>
      `shaderResourceMinLod` specifies whether image operations
      specifying the minimum resource LOD are supported in shader code.
      If this feature is not enabled, the `MinLod` image operand
      **must** not be used in shader code. This also specifies whether
      shader modules **can** declare the `MinLod` capability.

    - <span id="features-sparseBinding"></span> `sparseBinding`
      specifies whether resource memory **can** be managed at opaque
      sparse block level instead of at the object level. If this feature
      is not enabled, resource memory **must** be bound only on a
      per-object basis using the `vkBindBufferMemory` and
      `vkBindImageMemory` commands. In this case, buffers and images
      **must** not be created with `VK_BUFFER_CREATE_SPARSE_BINDING_BIT`
      and `VK_IMAGE_CREATE_SPARSE_BINDING_BIT` set in the `flags` member
      of the [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html) and
      [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) structures,
      respectively. Otherwise resource memory **can** be managed as
      described in <a
      href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#sparsememory-sparseresourcefeatures"
      target="_blank" rel="noopener">Sparse Resource Features</a>.

    - <span id="features-sparseResidencyBuffer"></span>
      `sparseResidencyBuffer` specifies whether the device **can**
      access partially resident buffers. If this feature is not enabled,
      buffers **must** not be created with
      `VK_BUFFER_CREATE_SPARSE_RESIDENCY_BIT` set in the `flags` member
      of the [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html) structure.

    - <span id="features-sparseResidencyImage2D"></span>
      `sparseResidencyImage2D` specifies whether the device **can**
      access partially resident 2D images with 1 sample per pixel. If
      this feature is not enabled, images with an `imageType` of
      `VK_IMAGE_TYPE_2D` and `samples` set to `VK_SAMPLE_COUNT_1_BIT`
      **must** not be created with
      `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT` set in the `flags` member
      of the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) structure.

    - <span id="features-sparseResidencyImage3D"></span>
      `sparseResidencyImage3D` specifies whether the device **can**
      access partially resident 3D images. If this feature is not
      enabled, images with an `imageType` of `VK_IMAGE_TYPE_3D` **must**
      not be created with `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT` set in
      the `flags` member of the
      [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) structure.

    - <span id="features-sparseResidency2Samples"></span>
      `sparseResidency2Samples` specifies whether the physical device
      **can** access partially resident 2D images with 2 samples per
      pixel. If this feature is not enabled, images with an `imageType`
      of `VK_IMAGE_TYPE_2D` and `samples` set to `VK_SAMPLE_COUNT_2_BIT`
      **must** not be created with
      `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT` set in the `flags` member
      of the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) structure.

    - <span id="features-sparseResidency4Samples"></span>
      `sparseResidency4Samples` specifies whether the physical device
      **can** access partially resident 2D images with 4 samples per
      pixel. If this feature is not enabled, images with an `imageType`
      of `VK_IMAGE_TYPE_2D` and `samples` set to `VK_SAMPLE_COUNT_4_BIT`
      **must** not be created with
      `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT` set in the `flags` member
      of the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) structure.

    - <span id="features-sparseResidency8Samples"></span>
      `sparseResidency8Samples` specifies whether the physical device
      **can** access partially resident 2D images with 8 samples per
      pixel. If this feature is not enabled, images with an `imageType`
      of `VK_IMAGE_TYPE_2D` and `samples` set to `VK_SAMPLE_COUNT_8_BIT`
      **must** not be created with
      `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT` set in the `flags` member
      of the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) structure.

    - <span id="features-sparseResidency16Samples"></span>
      `sparseResidency16Samples` specifies whether the physical device
      **can** access partially resident 2D images with 16 samples per
      pixel. If this feature is not enabled, images with an `imageType`
      of `VK_IMAGE_TYPE_2D` and `samples` set to
      `VK_SAMPLE_COUNT_16_BIT` **must** not be created with
      `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT` set in the `flags` member
      of the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) structure.

    - <span id="features-sparseResidencyAliased"></span>
      `sparseResidencyAliased` specifies whether the physical device
      **can** correctly access data aliased into multiple locations. If
      this feature is not enabled, the
      `VK_BUFFER_CREATE_SPARSE_ALIASED_BIT` and
      `VK_IMAGE_CREATE_SPARSE_ALIASED_BIT` enum values **must** not be
      used in `flags` members of the
      [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html) and
      [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) structures,
      respectively.

    - <span id="features-variableMultisampleRate"></span>
      `variableMultisampleRate` specifies whether all pipelines that
      will be bound to a command buffer during a <a
      href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-noattachments"
      target="_blank" rel="noopener">subpass which uses no attachments</a>
      **must** have the same value for
      [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)::`rasterizationSamples`.
      If set to `VK_TRUE`, the implementation supports variable
      multisample rates in a subpass which uses no attachments. If set
      to `VK_FALSE`, then all pipelines bound in such a subpass **must**
      have the same multisample rate. This has no effect in situations
      where a subpass uses any attachments.

    - <span id="features-inheritedQueries"></span> `inheritedQueries`
      specifies whether a secondary command buffer **may** be executed
      while a query is active.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html),
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
[vkGetPhysicalDeviceFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceFeatures"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
