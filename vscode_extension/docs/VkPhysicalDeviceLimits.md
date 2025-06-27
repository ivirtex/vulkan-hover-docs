# VkPhysicalDeviceLimits(3) Manual Page

## Name

VkPhysicalDeviceLimits - Structure reporting implementation-dependent
physical device limits



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceLimits` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkPhysicalDeviceLimits {
    uint32_t              maxImageDimension1D;
    uint32_t              maxImageDimension2D;
    uint32_t              maxImageDimension3D;
    uint32_t              maxImageDimensionCube;
    uint32_t              maxImageArrayLayers;
    uint32_t              maxTexelBufferElements;
    uint32_t              maxUniformBufferRange;
    uint32_t              maxStorageBufferRange;
    uint32_t              maxPushConstantsSize;
    uint32_t              maxMemoryAllocationCount;
    uint32_t              maxSamplerAllocationCount;
    VkDeviceSize          bufferImageGranularity;
    VkDeviceSize          sparseAddressSpaceSize;
    uint32_t              maxBoundDescriptorSets;
    uint32_t              maxPerStageDescriptorSamplers;
    uint32_t              maxPerStageDescriptorUniformBuffers;
    uint32_t              maxPerStageDescriptorStorageBuffers;
    uint32_t              maxPerStageDescriptorSampledImages;
    uint32_t              maxPerStageDescriptorStorageImages;
    uint32_t              maxPerStageDescriptorInputAttachments;
    uint32_t              maxPerStageResources;
    uint32_t              maxDescriptorSetSamplers;
    uint32_t              maxDescriptorSetUniformBuffers;
    uint32_t              maxDescriptorSetUniformBuffersDynamic;
    uint32_t              maxDescriptorSetStorageBuffers;
    uint32_t              maxDescriptorSetStorageBuffersDynamic;
    uint32_t              maxDescriptorSetSampledImages;
    uint32_t              maxDescriptorSetStorageImages;
    uint32_t              maxDescriptorSetInputAttachments;
    uint32_t              maxVertexInputAttributes;
    uint32_t              maxVertexInputBindings;
    uint32_t              maxVertexInputAttributeOffset;
    uint32_t              maxVertexInputBindingStride;
    uint32_t              maxVertexOutputComponents;
    uint32_t              maxTessellationGenerationLevel;
    uint32_t              maxTessellationPatchSize;
    uint32_t              maxTessellationControlPerVertexInputComponents;
    uint32_t              maxTessellationControlPerVertexOutputComponents;
    uint32_t              maxTessellationControlPerPatchOutputComponents;
    uint32_t              maxTessellationControlTotalOutputComponents;
    uint32_t              maxTessellationEvaluationInputComponents;
    uint32_t              maxTessellationEvaluationOutputComponents;
    uint32_t              maxGeometryShaderInvocations;
    uint32_t              maxGeometryInputComponents;
    uint32_t              maxGeometryOutputComponents;
    uint32_t              maxGeometryOutputVertices;
    uint32_t              maxGeometryTotalOutputComponents;
    uint32_t              maxFragmentInputComponents;
    uint32_t              maxFragmentOutputAttachments;
    uint32_t              maxFragmentDualSrcAttachments;
    uint32_t              maxFragmentCombinedOutputResources;
    uint32_t              maxComputeSharedMemorySize;
    uint32_t              maxComputeWorkGroupCount[3];
    uint32_t              maxComputeWorkGroupInvocations;
    uint32_t              maxComputeWorkGroupSize[3];
    uint32_t              subPixelPrecisionBits;
    uint32_t              subTexelPrecisionBits;
    uint32_t              mipmapPrecisionBits;
    uint32_t              maxDrawIndexedIndexValue;
    uint32_t              maxDrawIndirectCount;
    float                 maxSamplerLodBias;
    float                 maxSamplerAnisotropy;
    uint32_t              maxViewports;
    uint32_t              maxViewportDimensions[2];
    float                 viewportBoundsRange[2];
    uint32_t              viewportSubPixelBits;
    size_t                minMemoryMapAlignment;
    VkDeviceSize          minTexelBufferOffsetAlignment;
    VkDeviceSize          minUniformBufferOffsetAlignment;
    VkDeviceSize          minStorageBufferOffsetAlignment;
    int32_t               minTexelOffset;
    uint32_t              maxTexelOffset;
    int32_t               minTexelGatherOffset;
    uint32_t              maxTexelGatherOffset;
    float                 minInterpolationOffset;
    float                 maxInterpolationOffset;
    uint32_t              subPixelInterpolationOffsetBits;
    uint32_t              maxFramebufferWidth;
    uint32_t              maxFramebufferHeight;
    uint32_t              maxFramebufferLayers;
    VkSampleCountFlags    framebufferColorSampleCounts;
    VkSampleCountFlags    framebufferDepthSampleCounts;
    VkSampleCountFlags    framebufferStencilSampleCounts;
    VkSampleCountFlags    framebufferNoAttachmentsSampleCounts;
    uint32_t              maxColorAttachments;
    VkSampleCountFlags    sampledImageColorSampleCounts;
    VkSampleCountFlags    sampledImageIntegerSampleCounts;
    VkSampleCountFlags    sampledImageDepthSampleCounts;
    VkSampleCountFlags    sampledImageStencilSampleCounts;
    VkSampleCountFlags    storageImageSampleCounts;
    uint32_t              maxSampleMaskWords;
    VkBool32              timestampComputeAndGraphics;
    float                 timestampPeriod;
    uint32_t              maxClipDistances;
    uint32_t              maxCullDistances;
    uint32_t              maxCombinedClipAndCullDistances;
    uint32_t              discreteQueuePriorities;
    float                 pointSizeRange[2];
    float                 lineWidthRange[2];
    float                 pointSizeGranularity;
    float                 lineWidthGranularity;
    VkBool32              strictLines;
    VkBool32              standardSampleLocations;
    VkDeviceSize          optimalBufferCopyOffsetAlignment;
    VkDeviceSize          optimalBufferCopyRowPitchAlignment;
    VkDeviceSize          nonCoherentAtomSize;
} VkPhysicalDeviceLimits;
```

## <a href="#_members" class="anchor"></a>Members

The `VkPhysicalDeviceLimits` are properties of the physical device.
These are available in the `limits` member of the
[VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html) structure
which is returned from
[vkGetPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties.html).

## <a href="#_description" class="anchor"></a>Description

- <span id="limits-maxImageDimension1D"></span> `maxImageDimension1D` is
  the largest dimension (`width`) that is guaranteed to be supported for
  all images created with an `imageType` of `VK_IMAGE_TYPE_1D`. Some
  combinations of image parameters (format, usage, etc.) **may** allow
  support for larger dimensions, which **can** be queried using
  [vkGetPhysicalDeviceImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties.html).

- <span id="limits-maxImageDimension2D"></span> `maxImageDimension2D` is
  the largest dimension (`width` or `height`) that is guaranteed to be
  supported for all images created with an `imageType` of
  `VK_IMAGE_TYPE_2D` and without `VK_IMAGE_CREATE_CUBE_COMPATIBLE_BIT`
  set in `flags`. Some combinations of image parameters (format, usage,
  etc.) **may** allow support for larger dimensions, which **can** be
  queried using
  [vkGetPhysicalDeviceImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties.html).

- <span id="limits-maxImageDimension3D"></span> `maxImageDimension3D` is
  the largest dimension (`width`, `height`, or `depth`) that is
  guaranteed to be supported for all images created with an `imageType`
  of `VK_IMAGE_TYPE_3D`. Some combinations of image parameters (format,
  usage, etc.) **may** allow support for larger dimensions, which
  **can** be queried using
  [vkGetPhysicalDeviceImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties.html).

- <span id="limits-maxImageDimensionCube"></span>
  `maxImageDimensionCube` is the largest dimension (`width` or `height`)
  that is guaranteed to be supported for all images created with an
  `imageType` of `VK_IMAGE_TYPE_2D` and with
  `VK_IMAGE_CREATE_CUBE_COMPATIBLE_BIT` set in `flags`. Some
  combinations of image parameters (format, usage, etc.) **may** allow
  support for larger dimensions, which **can** be queried using
  [vkGetPhysicalDeviceImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties.html).

- <span id="limits-maxImageArrayLayers"></span> `maxImageArrayLayers` is
  the maximum number of layers (`arrayLayers`) for an image.

- <span id="limits-maxTexelBufferElements"></span>
  `maxTexelBufferElements` is the maximum number of addressable texels
  for a buffer view created on a buffer which was created with the
  `VK_BUFFER_USAGE_UNIFORM_TEXEL_BUFFER_BIT` or
  `VK_BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT` set in the `usage` member
  of the [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html) structure.

- <span id="limits-maxUniformBufferRange"></span>
  `maxUniformBufferRange` is the maximum value that **can** be specified
  in the `range` member of a
  [VkDescriptorBufferInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBufferInfo.html) structure passed
  to [vkUpdateDescriptorSets](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSets.html) for
  descriptors of type `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER` or
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`.

- <span id="limits-maxStorageBufferRange"></span>
  `maxStorageBufferRange` is the maximum value that **can** be specified
  in the `range` member of a
  [VkDescriptorBufferInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBufferInfo.html) structure passed
  to [vkUpdateDescriptorSets](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSets.html) for
  descriptors of type `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER` or
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`.

- <span id="limits-maxPushConstantsSize"></span> `maxPushConstantsSize`
  is the maximum size, in bytes, of the pool of push constant memory.
  For each of the push constant ranges indicated by the
  `pPushConstantRanges` member of the
  [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html)
  structure, (`offset` + `size`) **must** be less than or equal to this
  limit.

- <span id="limits-maxMemoryAllocationCount"></span>
  `maxMemoryAllocationCount` is the maximum number of device memory
  allocations, as created by [vkAllocateMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAllocateMemory.html),
  which **can** simultaneously exist.

- <span id="limits-maxSamplerAllocationCount"></span>
  `maxSamplerAllocationCount` is the maximum number of sampler objects,
  as created by [vkCreateSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSampler.html), which **can**
  simultaneously exist on a device.

- <span id="limits-bufferImageGranularity"></span>
  `bufferImageGranularity` is the granularity, in bytes, at which buffer
  or linear image resources, and optimal image resources **can** be
  bound to adjacent offsets in the same `VkDeviceMemory` object without
  aliasing. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-bufferimagegranularity"
  target="_blank" rel="noopener">Buffer-Image Granularity</a> for more
  details.

- <span id="limits-sparseAddressSpaceSize"></span>
  `sparseAddressSpaceSize` is the total amount of address space
  available, in bytes, for sparse memory resources. This is an upper
  bound on the sum of the sizes of all sparse resources, regardless of
  whether any memory is bound to them. If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedSparseAddressSpace"
  target="_blank"
  rel="noopener"><code>extendedSparseAddressSpace</code></a> feature is
  enabled, then the difference between <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-extendedSparseAddressSpaceSize"
  target="_blank"
  rel="noopener"><code>extendedSparseAddressSpaceSize</code></a> and
  `sparseAddressSpaceSize` can also be used, by `VkImage` created with
  the `usage` member of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) only
  containing bits in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-extendedSparseImageUsageFlags"
  target="_blank"
  rel="noopener"><code>extendedSparseImageUsageFlags</code></a> and
  `VkBuffer` created with the `usage` member of
  [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html) only containing bits in
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-extendedSparseBufferUsageFlags"
  target="_blank"
  rel="noopener"><code>extendedSparseBufferUsageFlags</code></a>.

- <span id="limits-maxBoundDescriptorSets"></span>
  `maxBoundDescriptorSets` is the maximum number of descriptor sets that
  **can** be simultaneously used by a pipeline. All `DescriptorSet`
  decorations in shader modules **must** have a value less than
  `maxBoundDescriptorSets`. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-sets"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-sets</a>.

- <span id="limits-maxPerStageDescriptorSamplers"></span>
  `maxPerStageDescriptorSamplers` is the maximum number of samplers that
  **can** be accessible to a single shader stage in a pipeline layout.
  Descriptors with a type of `VK_DESCRIPTOR_TYPE_SAMPLER` or
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` count against this limit.
  Only descriptors in descriptor set layouts created without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  count against this limit. A descriptor is accessible to a shader stage
  when the `stageFlags` member of the `VkDescriptorSetLayoutBinding`
  structure has the bit for that shader stage set. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-sampler"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-sampler</a>
  and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-combinedimagesampler"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-combinedimagesampler</a>.

- <span id="limits-maxPerStageDescriptorUniformBuffers"></span>
  `maxPerStageDescriptorUniformBuffers` is the maximum number of uniform
  buffers that **can** be accessible to a single shader stage in a
  pipeline layout. Descriptors with a type of
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER` or
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` count against this limit.
  Only descriptors in descriptor set layouts created without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  count against this limit. A descriptor is accessible to a shader stage
  when the `stageFlags` member of the `VkDescriptorSetLayoutBinding`
  structure has the bit for that shader stage set. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-uniformbuffer"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-uniformbuffer</a>
  and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-uniformbufferdynamic"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-uniformbufferdynamic</a>.

- <span id="limits-maxPerStageDescriptorStorageBuffers"></span>
  `maxPerStageDescriptorStorageBuffers` is the maximum number of storage
  buffers that **can** be accessible to a single shader stage in a
  pipeline layout. Descriptors with a type of
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER` or
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` count against this limit.
  Only descriptors in descriptor set layouts created without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  count against this limit. A descriptor is accessible to a pipeline
  shader stage when the `stageFlags` member of the
  `VkDescriptorSetLayoutBinding` structure has the bit for that shader
  stage set. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storagebuffer"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storagebuffer</a>
  and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storagebufferdynamic"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storagebufferdynamic</a>.

- <span id="limits-maxPerStageDescriptorSampledImages"></span>
  `maxPerStageDescriptorSampledImages` is the maximum number of sampled
  images that **can** be accessible to a single shader stage in a
  pipeline layout. Descriptors with a type of
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`,
  `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, or
  `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` count against this limit.
  Only descriptors in descriptor set layouts created without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  count against this limit. A descriptor is accessible to a pipeline
  shader stage when the `stageFlags` member of the
  `VkDescriptorSetLayoutBinding` structure has the bit for that shader
  stage set. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-combinedimagesampler"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-combinedimagesampler</a>,
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-sampledimage"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-sampledimage</a>,
  and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-uniformtexelbuffer"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-uniformtexelbuffer</a>.

- <span id="limits-maxPerStageDescriptorStorageImages"></span>
  `maxPerStageDescriptorStorageImages` is the maximum number of storage
  images that **can** be accessible to a single shader stage in a
  pipeline layout. Descriptors with a type of
  `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, or
  `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` count against this limit.
  Only descriptors in descriptor set layouts created without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  count against this limit. A descriptor is accessible to a pipeline
  shader stage when the `stageFlags` member of the
  `VkDescriptorSetLayoutBinding` structure has the bit for that shader
  stage set. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storageimage"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storageimage</a>,
  and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storagetexelbuffer"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storagetexelbuffer</a>.

- <span id="limits-maxPerStageDescriptorInputAttachments"></span>
  `maxPerStageDescriptorInputAttachments` is the maximum number of input
  attachments that **can** be accessible to a single shader stage in a
  pipeline layout. Descriptors with a type of
  `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT` count against this limit. Only
  descriptors in descriptor set layouts created without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  count against this limit. A descriptor is accessible to a pipeline
  shader stage when the `stageFlags` member of the
  `VkDescriptorSetLayoutBinding` structure has the bit for that shader
  stage set. These are only supported for the fragment stage. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-inputattachment"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-inputattachment</a>.

- <span id="limits-maxPerStageResources"></span> `maxPerStageResources`
  is the maximum number of resources that **can** be accessible to a
  single shader stage in a pipeline layout. Descriptors with a type of
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`,
  `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`,
  `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`,
  `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER`,
  `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER`,
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER`,
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER`,
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`,
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`, or
  `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT` count against this limit. Only
  descriptors in descriptor set layouts created without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  count against this limit. For the fragment shader stage the
  framebuffer color attachments also count against this limit.

- <span id="limits-maxDescriptorSetSamplers"></span>
  `maxDescriptorSetSamplers` is the maximum number of samplers that
  **can** be included in a pipeline layout. Descriptors with a type of
  `VK_DESCRIPTOR_TYPE_SAMPLER` or
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` count against this limit.
  Only descriptors in descriptor set layouts created without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  count against this limit. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-sampler"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-sampler</a>
  and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-combinedimagesampler"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-combinedimagesampler</a>.

- <span id="limits-maxDescriptorSetUniformBuffers"></span>
  `maxDescriptorSetUniformBuffers` is the maximum number of uniform
  buffers that **can** be included in a pipeline layout. Descriptors
  with a type of `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER` or
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` count against this limit.
  Only descriptors in descriptor set layouts created without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  count against this limit. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-uniformbuffer"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-uniformbuffer</a>
  and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-uniformbufferdynamic"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-uniformbufferdynamic</a>.

- <span id="limits-maxDescriptorSetUniformBuffersDynamic"></span>
  `maxDescriptorSetUniformBuffersDynamic` is the maximum number of
  dynamic uniform buffers that **can** be included in a pipeline layout.
  Descriptors with a type of `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`
  count against this limit. Only descriptors in descriptor set layouts
  created without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  count against this limit. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-uniformbufferdynamic"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-uniformbufferdynamic</a>.

- <span id="limits-maxDescriptorSetStorageBuffers"></span>
  `maxDescriptorSetStorageBuffers` is the maximum number of storage
  buffers that **can** be included in a pipeline layout. Descriptors
  with a type of `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER` or
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` count against this limit.
  Only descriptors in descriptor set layouts created without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  count against this limit. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storagebuffer"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storagebuffer</a>
  and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storagebufferdynamic"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storagebufferdynamic</a>.

- <span id="limits-maxDescriptorSetStorageBuffersDynamic"></span>
  `maxDescriptorSetStorageBuffersDynamic` is the maximum number of
  dynamic storage buffers that **can** be included in a pipeline layout.
  Descriptors with a type of `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`
  count against this limit. Only descriptors in descriptor set layouts
  created without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  count against this limit. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storagebufferdynamic"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storagebufferdynamic</a>.

- <span id="limits-maxDescriptorSetSampledImages"></span>
  `maxDescriptorSetSampledImages` is the maximum number of sampled
  images that **can** be included in a pipeline layout. Descriptors with
  a type of `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`,
  `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, or
  `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` count against this limit.
  Only descriptors in descriptor set layouts created without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  count against this limit. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-combinedimagesampler"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-combinedimagesampler</a>,
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-sampledimage"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-sampledimage</a>,
  and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-uniformtexelbuffer"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-uniformtexelbuffer</a>.

- <span id="limits-maxDescriptorSetStorageImages"></span>
  `maxDescriptorSetStorageImages` is the maximum number of storage
  images that **can** be included in a pipeline layout. Descriptors with
  a type of `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, or
  `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` count against this limit.
  Only descriptors in descriptor set layouts created without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set
  count against this limit. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storageimage"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storageimage</a>,
  and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storagetexelbuffer"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storagetexelbuffer</a>.

- <span id="limits-maxDescriptorSetInputAttachments"></span>
  `maxDescriptorSetInputAttachments` is the maximum number of input
  attachments that **can** be included in a pipeline layout. Descriptors
  with a type of `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT` count against
  this limit. Only descriptors in descriptor set layouts created without
  the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit
  set count against this limit. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-inputattachment"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-inputattachment</a>.

- <span id="limits-maxVertexInputAttributes"></span>
  `maxVertexInputAttributes` is the maximum number of vertex input
  attributes that **can** be specified for a graphics pipeline. These
  are described in the array of `VkVertexInputAttributeDescription`
  structures that are provided at graphics pipeline creation time via
  the `pVertexAttributeDescriptions` member of the
  [VkPipelineVertexInputStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineVertexInputStateCreateInfo.html)
  structure. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fxvertex-attrib"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fxvertex-attrib</a>
  and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fxvertex-input"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fxvertex-input</a>.

- <span id="limits-maxVertexInputBindings"></span>
  `maxVertexInputBindings` is the maximum number of vertex buffers that
  **can** be specified for providing vertex attributes to a graphics
  pipeline. These are described in the array of
  `VkVertexInputBindingDescription` structures that are provided at
  graphics pipeline creation time via the `pVertexBindingDescriptions`
  member of the
  [VkPipelineVertexInputStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineVertexInputStateCreateInfo.html)
  structure. The `binding` member of `VkVertexInputBindingDescription`
  **must** be less than this limit. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fxvertex-input"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fxvertex-input</a>.

- <span id="limits-maxVertexInputAttributeOffset"></span>
  `maxVertexInputAttributeOffset` is the maximum vertex input attribute
  offset that **can** be added to the vertex input binding stride. The
  `offset` member of the `VkVertexInputAttributeDescription` structure
  **must** be less than or equal to this limit. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fxvertex-input"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fxvertex-input</a>.

- <span id="limits-maxVertexInputBindingStride"></span>
  `maxVertexInputBindingStride` is the maximum vertex input binding
  stride that **can** be specified in a vertex input binding. The
  `stride` member of the `VkVertexInputBindingDescription` structure
  **must** be less than or equal to this limit. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fxvertex-input"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fxvertex-input</a>.

- <span id="limits-maxVertexOutputComponents"></span>
  `maxVertexOutputComponents` is the maximum number of components of
  output variables which **can** be output by a vertex shader. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-vertex"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-vertex</a>.

- <span id="limits-maxTessellationGenerationLevel"></span>
  `maxTessellationGenerationLevel` is the maximum tessellation
  generation level supported by the fixed-function tessellation
  primitive generator. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#tessellation"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#tessellation</a>.

- <span id="limits-maxTessellationPatchSize"></span>
  `maxTessellationPatchSize` is the maximum patch size, in vertices, of
  patches that **can** be processed by the tessellation control shader
  and tessellation primitive generator. The `patchControlPoints` member
  of the
  [VkPipelineTessellationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineTessellationStateCreateInfo.html)
  structure specified at pipeline creation time and the value provided
  in the `OutputVertices` execution mode of shader modules **must** be
  less than or equal to this limit. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#tessellation"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#tessellation</a>.

- <span id="limits-maxTessellationControlPerVertexInputComponents"></span>
  `maxTessellationControlPerVertexInputComponents` is the maximum number
  of components of input variables which **can** be provided as
  per-vertex inputs to the tessellation control shader stage.

- <span id="limits-maxTessellationControlPerVertexOutputComponents"></span>
  `maxTessellationControlPerVertexOutputComponents` is the maximum
  number of components of per-vertex output variables which **can** be
  output from the tessellation control shader stage.

- <span id="limits-maxTessellationControlPerPatchOutputComponents"></span>
  `maxTessellationControlPerPatchOutputComponents` is the maximum number
  of components of per-patch output variables which **can** be output
  from the tessellation control shader stage.

- <span id="limits-maxTessellationControlTotalOutputComponents"></span>
  `maxTessellationControlTotalOutputComponents` is the maximum total
  number of components of per-vertex and per-patch output variables
  which **can** be output from the tessellation control shader stage.

- <span id="limits-maxTessellationEvaluationInputComponents"></span>
  `maxTessellationEvaluationInputComponents` is the maximum number of
  components of input variables which **can** be provided as per-vertex
  inputs to the tessellation evaluation shader stage.

- <span id="limits-maxTessellationEvaluationOutputComponents"></span>
  `maxTessellationEvaluationOutputComponents` is the maximum number of
  components of per-vertex output variables which **can** be output from
  the tessellation evaluation shader stage.

- <span id="limits-maxGeometryShaderInvocations"></span>
  `maxGeometryShaderInvocations` is the maximum invocation count
  supported for instanced geometry shaders. The value provided in the
  `Invocations` execution mode of shader modules **must** be less than
  or equal to this limit. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#geometry"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#geometry</a>.

- <span id="limits-maxGeometryInputComponents"></span>
  `maxGeometryInputComponents` is the maximum number of components of
  input variables which **can** be provided as inputs to the geometry
  shader stage.

- <span id="limits-maxGeometryOutputComponents"></span>
  `maxGeometryOutputComponents` is the maximum number of components of
  output variables which **can** be output from the geometry shader
  stage.

- <span id="limits-maxGeometryOutputVertices"></span>
  `maxGeometryOutputVertices` is the maximum number of vertices which
  **can** be emitted by any geometry shader.

- <span id="limits-maxGeometryTotalOutputComponents"></span>
  `maxGeometryTotalOutputComponents` is the maximum total number of
  components of output variables, across all emitted vertices, which
  **can** be output from the geometry shader stage.

- <span id="limits-maxFragmentInputComponents"></span>
  `maxFragmentInputComponents` is the maximum number of components of
  input variables which **can** be provided as inputs to the fragment
  shader stage.

- <span id="limits-maxFragmentOutputAttachments"></span>
  `maxFragmentOutputAttachments` is the maximum number of output
  attachments which **can** be written to by the fragment shader stage.

- <span id="limits-maxFragmentDualSrcAttachments"></span>
  `maxFragmentDualSrcAttachments` is the maximum number of output
  attachments which **can** be written to by the fragment shader stage
  when blending is enabled and one of the dual source blend modes is in
  use. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-dsb"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-dsb</a>
  and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-dualSrcBlend"
  target="_blank" rel="noopener"><code>dualSrcBlend</code></a>.

- <span id="limits-maxFragmentCombinedOutputResources"></span>
  `maxFragmentCombinedOutputResources` is the total number of storage
  buffers, storage images, and output `Location` decorated color
  attachments (described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-fragmentoutput"
  target="_blank" rel="noopener">Fragment Output Interface</a>) which
  **can** be used in the fragment shader stage.

- <span id="limits-maxComputeSharedMemorySize"></span>
  `maxComputeSharedMemorySize` is the maximum total storage size, in
  bytes, available for variables declared with the `Workgroup` storage
  class in shader modules (or with the `shared` storage qualifier in
  GLSL) in the compute shader stage.

- <span id="limits-maxComputeWorkGroupCount"></span>
  `maxComputeWorkGroupCount`\[3\] is the maximum number of local
  workgroups that **can** be dispatched by a single dispatching command.
  These three values represent the maximum number of local workgroups
  for the X, Y, and Z dimensions, respectively. The workgroup count
  parameters to the dispatching commands **must** be less than or equal
  to the corresponding limit. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dispatch"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dispatch</a>.

- <span id="limits-maxComputeWorkGroupInvocations"></span>
  `maxComputeWorkGroupInvocations` is the maximum total number of
  compute shader invocations in a single local workgroup. The product of
  the X, Y, and Z sizes, as specified by the `LocalSize` or
  `LocalSizeId` execution mode in shader modules or by the object
  decorated by the `WorkgroupSize` decoration, **must** be less than or
  equal to this limit.

- <span id="limits-maxComputeWorkGroupSize"></span>
  `maxComputeWorkGroupSize`\[3\] is the maximum size of a local compute
  workgroup, per dimension. These three values represent the maximum
  local workgroup size in the X, Y, and Z dimensions, respectively. The
  `x`, `y`, and `z` sizes, as specified by the `LocalSize` or
  `LocalSizeId` execution mode or by the object decorated by the
  `WorkgroupSize` decoration in shader modules, **must** be less than or
  equal to the corresponding limit.

- <span id="limits-subPixelPrecisionBits"></span>
  `subPixelPrecisionBits` is the number of bits of subpixel precision in
  framebuffer coordinates x<sub>f</sub> and y<sub>f</sub>. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast</a>.

- <span id="limits-subTexelPrecisionBits"></span>
  `subTexelPrecisionBits` is the number of bits of precision in the
  division along an axis of an image used for minification and
  magnification filters. 2<sup>`subTexelPrecisionBits`</sup> is the
  actual number of divisions along each axis of the image represented.
  Sub-texel values calculated during image sampling will snap to these
  locations when generating the filtered results.

- <span id="limits-mipmapPrecisionBits"></span> `mipmapPrecisionBits` is
  the number of bits of division that the LOD calculation for mipmap
  fetching get snapped to when determining the contribution from each
  mip level to the mip filtered results.
  2<sup>`mipmapPrecisionBits`</sup> is the actual number of divisions.

- <span id="limits-maxDrawIndexedIndexValue"></span>
  `maxDrawIndexedIndexValue` is the maximum index value that **can** be
  used for indexed draw calls when using 32-bit indices. This excludes
  the primitive restart index value of 0xFFFFFFFF. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-fullDrawIndexUint32"
  target="_blank" rel="noopener"><code>fullDrawIndexUint32</code></a>.

- <span id="limits-maxDrawIndirectCount"></span> `maxDrawIndirectCount`
  is the maximum draw count that is supported for indirect drawing
  calls. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiDrawIndirect"
  target="_blank" rel="noopener"><code>multiDrawIndirect</code></a>.

- <span id="limits-maxSamplerLodBias"></span> `maxSamplerLodBias` is the
  maximum absolute sampler LOD bias. The sum of the `mipLodBias` member
  of the [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html) structure and
  the `Bias` operand of image sampling operations in shader modules (or
  0 if no `Bias` operand is provided to an image sampling operation) are
  clamped to the range \[-`maxSamplerLodBias`,+`maxSamplerLodBias`\].
  See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#samplers-mipLodBias"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#samplers-mipLodBias</a>.

- <span id="limits-maxSamplerAnisotropy"></span> `maxSamplerAnisotropy`
  is the maximum degree of sampler anisotropy. The maximum degree of
  anisotropic filtering used for an image sampling operation is the
  minimum of the `maxAnisotropy` member of the
  [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html) structure and this
  limit. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#samplers-maxAnisotropy"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#samplers-maxAnisotropy</a>.

- <span id="limits-maxViewports"></span> `maxViewports` is the maximum
  number of active viewports. The `viewportCount` member of the
  [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html)
  structure that is provided at pipeline creation **must** be less than
  or equal to this limit.

- <span id="limits-maxViewportDimensions"></span>
  `maxViewportDimensions`\[2\] are the maximum viewport dimensions in
  the X (width) and Y (height) dimensions, respectively. The maximum
  viewport dimensions **must** be greater than or equal to the largest
  image which **can** be created and used as a framebuffer attachment.
  See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vertexpostproc-viewport"
  target="_blank" rel="noopener">Controlling the Viewport</a>.

- <span id="limits-viewportboundsrange"></span>
  `viewportBoundsRange`\[2\] is the \[minimum, maximum\] range that the
  corners of a viewport **must** be contained in. This range **must** be
  at least \[-2 × `size`, 2 × `size` - 1\], where `size` =
  max(`maxViewportDimensions`\[0\], `maxViewportDimensions`\[1\]). See
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vertexpostproc-viewport"
  target="_blank" rel="noopener">Controlling the Viewport</a>.

  <table>
  <colgroup>
  <col style="width: 50%" />
  <col style="width: 50%" />
  </colgroup>
  <tbody>
  <tr>
  <td class="icon"><em></em></td>
  <td class="content">Note
  <p>The intent of the <code>viewportBoundsRange</code> limit is to allow
  a maximum sized viewport to be arbitrarily shifted relative to the
  output target as long as at least some portion intersects. This would
  give a bounds limit of [-<code>size</code> + 1, 2 × <code>size</code> -
  1] which would allow all possible non-empty-set intersections of the
  output target and the viewport. Since these numbers are typically powers
  of two, picking the signed number range using the smallest possible
  number of bits ends up with the specified range.</p></td>
  </tr>
  </tbody>
  </table>

- <span id="limits-viewportSubPixelBits"></span> `viewportSubPixelBits`
  is the number of bits of subpixel precision for viewport bounds. The
  subpixel precision that floating-point viewport bounds are interpreted
  at is given by this limit.

- <span id="limits-minMemoryMapAlignment"></span>
  `minMemoryMapAlignment` is the minimum **required** alignment, in
  bytes, of host visible memory allocations within the host address
  space. When mapping a memory allocation with
  [vkMapMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkMapMemory.html), subtracting `offset` bytes from the
  returned pointer will always produce an integer multiple of this
  limit. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-device-hostaccess"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-device-hostaccess</a>.
  The value **must** be a power of two.

- <span id="limits-minTexelBufferOffsetAlignment"></span>
  `minTexelBufferOffsetAlignment` is the minimum **required** alignment,
  in bytes, for the `offset` member of the
  [VkBufferViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferViewCreateInfo.html) structure for
  texel buffers. The value **must** be a power of two. If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-texelBufferAlignment"
  target="_blank" rel="noopener"><code>texelBufferAlignment</code></a>
  is enabled, this limit is equivalent to the maximum of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-uniformTexelBufferOffsetAlignmentBytes"
  target="_blank"
  rel="noopener"><code>uniformTexelBufferOffsetAlignmentBytes</code></a>
  and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-storageTexelBufferOffsetAlignmentBytes"
  target="_blank"
  rel="noopener"><code>storageTexelBufferOffsetAlignmentBytes</code></a>
  members of
  [VkPhysicalDeviceTexelBufferAlignmentProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTexelBufferAlignmentProperties.html),
  but smaller alignment is **optionally** allowed by <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-storageTexelBufferOffsetSingleTexelAlignment"
  target="_blank"
  rel="noopener"><code>storageTexelBufferOffsetSingleTexelAlignment</code></a>
  and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-uniformTexelBufferOffsetSingleTexelAlignment"
  target="_blank"
  rel="noopener"><code>uniformTexelBufferOffsetSingleTexelAlignment</code></a>.
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-texelBufferAlignment"
  target="_blank" rel="noopener"><code>texelBufferAlignment</code></a>
  is not enabled,
  [VkBufferViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferViewCreateInfo.html)::`offset`
  **must** be a multiple of this value.

- <span id="limits-minUniformBufferOffsetAlignment"></span>
  `minUniformBufferOffsetAlignment` is the minimum **required**
  alignment, in bytes, for the `offset` member of the
  `VkDescriptorBufferInfo` structure for uniform buffers. When a
  descriptor of type `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER` or
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` is updated, the `offset`
  **must** be an integer multiple of this limit. Similarly, dynamic
  offsets for uniform buffers **must** be multiples of this limit. The
  value **must** be a power of two.

- <span id="limits-minStorageBufferOffsetAlignment"></span>
  `minStorageBufferOffsetAlignment` is the minimum **required**
  alignment, in bytes, for the `offset` member of the
  `VkDescriptorBufferInfo` structure for storage buffers. When a
  descriptor of type `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER` or
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` is updated, the `offset`
  **must** be an integer multiple of this limit. Similarly, dynamic
  offsets for storage buffers **must** be multiples of this limit. The
  value **must** be a power of two.

- <span id="limits-minTexelOffset"></span> `minTexelOffset` is the
  minimum offset value for the `ConstOffset` image operand of any of the
  `OpImageSample*` or `OpImageFetch*` image instructions.

- <span id="limits-maxTexelOffset"></span> `maxTexelOffset` is the
  maximum offset value for the `ConstOffset` image operand of any of the
  `OpImageSample*` or `OpImageFetch*` image instructions.

- <span id="limits-minTexelGatherOffset"></span> `minTexelGatherOffset`
  is the minimum offset value for the `Offset`, `ConstOffset`, or
  `ConstOffsets` image operands of any of the `OpImage*Gather` image
  instructions.

- <span id="limits-maxTexelGatherOffset"></span> `maxTexelGatherOffset`
  is the maximum offset value for the `Offset`, `ConstOffset`, or
  `ConstOffsets` image operands of any of the `OpImage*Gather` image
  instructions.

- <span id="limits-minInterpolationOffset"></span>
  `minInterpolationOffset` is the base minimum (inclusive) negative
  offset value for the `Offset` operand of the `InterpolateAtOffset`
  extended instruction.

- <span id="limits-maxInterpolationOffset"></span>
  `maxInterpolationOffset` is the base maximum (inclusive) positive
  offset value for the `Offset` operand of the `InterpolateAtOffset`
  extended instruction.

- <span id="limits-subPixelInterpolationOffsetBits"></span>
  `subPixelInterpolationOffsetBits` is the number of fractional bits
  that the `x` and `y` offsets to the `InterpolateAtOffset` extended
  instruction **may** be rounded to as fixed-point values.

- <span id="limits-maxFramebufferWidth"></span> `maxFramebufferWidth` is
  the maximum width for a framebuffer. The `width` member of the
  [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html) structure
  **must** be less than or equal to this limit.

- <span id="limits-maxFramebufferHeight"></span> `maxFramebufferHeight`
  is the maximum height for a framebuffer. The `height` member of the
  [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html) structure
  **must** be less than or equal to this limit.

- <span id="limits-maxFramebufferLayers"></span> `maxFramebufferLayers`
  is the maximum layer count for a layered framebuffer. The `layers`
  member of the [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateInfo.html)
  structure **must** be less than or equal to this limit.

- <span id="limits-framebufferColorSampleCounts"></span>
  `framebufferColorSampleCounts` is a bitmask<sup>1</sup> of
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) indicating the
  color sample counts that are supported for all framebuffer color
  attachments with floating- or fixed-point formats. For color
  attachments with integer formats, see <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-framebufferIntegerColorSampleCounts"
  target="_blank"
  rel="noopener"><code>framebufferIntegerColorSampleCounts</code></a>.

- <span id="limits-framebufferDepthSampleCounts"></span>
  `framebufferDepthSampleCounts` is a bitmask<sup>1</sup> of
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) indicating the
  supported depth sample counts for all framebuffer depth/stencil
  attachments, when the format includes a depth component.

- <span id="limits-framebufferStencilSampleCounts"></span>
  `framebufferStencilSampleCounts` is a bitmask<sup>1</sup> of
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) indicating the
  supported stencil sample counts for all framebuffer depth/stencil
  attachments, when the format includes a stencil component.

- <span id="limits-framebufferNoAttachmentsSampleCounts"></span>
  `framebufferNoAttachmentsSampleCounts` is a bitmask<sup>1</sup> of
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) indicating the
  supported sample counts for a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-noattachments"
  target="_blank" rel="noopener">subpass which uses no attachments</a>.

- <span id="limits-maxColorAttachments"></span> `maxColorAttachments` is
  the maximum number of color attachments that **can** be used by a
  subpass in a render pass. The `colorAttachmentCount` member of the
  `VkSubpassDescription` or `VkSubpassDescription2` structure **must**
  be less than or equal to this limit.

- <span id="limits-sampledImageColorSampleCounts"></span>
  `sampledImageColorSampleCounts` is a bitmask<sup>1</sup> of
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) indicating the
  sample counts supported for all 2D images created with
  `VK_IMAGE_TILING_OPTIMAL`, `usage` containing
  `VK_IMAGE_USAGE_SAMPLED_BIT`, and a non-integer color format.

- <span id="limits-sampledImageIntegerSampleCounts"></span>
  `sampledImageIntegerSampleCounts` is a bitmask<sup>1</sup> of
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) indicating the
  sample counts supported for all 2D images created with
  `VK_IMAGE_TILING_OPTIMAL`, `usage` containing
  `VK_IMAGE_USAGE_SAMPLED_BIT`, and an integer color format.

- <span id="limits-sampledImageDepthSampleCounts"></span>
  `sampledImageDepthSampleCounts` is a bitmask<sup>1</sup> of
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) indicating the
  sample counts supported for all 2D images created with
  `VK_IMAGE_TILING_OPTIMAL`, `usage` containing
  `VK_IMAGE_USAGE_SAMPLED_BIT`, and a depth format.

- <span id="limits-sampledImageStencilSampleCounts"></span>
  `sampledImageStencilSampleCounts` is a bitmask<sup>1</sup> of
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) indicating the
  sample counts supported for all 2D images created with
  `VK_IMAGE_TILING_OPTIMAL`, `usage` containing
  `VK_IMAGE_USAGE_SAMPLED_BIT`, and a stencil format.

- <span id="limits-storageImageSampleCounts"></span>
  `storageImageSampleCounts` is a bitmask<sup>1</sup> of
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) indicating the
  sample counts supported for all 2D images created with
  `VK_IMAGE_TILING_OPTIMAL`, and `usage` containing
  `VK_IMAGE_USAGE_STORAGE_BIT`.

- <span id="limits-maxSampleMaskWords"></span> `maxSampleMaskWords` is
  the maximum number of array elements of a variable decorated with the
  `SampleMask` built-in decoration.

- <span id="limits-timestampComputeAndGraphics"></span>
  `timestampComputeAndGraphics` specifies support for timestamps on all
  graphics and compute queues. If this limit is set to `VK_TRUE`, all
  queues that advertise the `VK_QUEUE_GRAPHICS_BIT` or
  `VK_QUEUE_COMPUTE_BIT` in the `VkQueueFamilyProperties`::`queueFlags`
  support `VkQueueFamilyProperties`::`timestampValidBits` of at
  least 36. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-timestamps"
  target="_blank" rel="noopener">Timestamp Queries</a>.

- <span id="limits-timestampPeriod"></span> `timestampPeriod` is the
  number of nanoseconds **required** for a timestamp query to be
  incremented by 1. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#queries-timestamps"
  target="_blank" rel="noopener">Timestamp Queries</a>.

- <span id="limits-maxClipDistances"></span> `maxClipDistances` is the
  maximum number of clip distances that **can** be used in a single
  shader stage. The size of any array declared with the `ClipDistance`
  built-in decoration in a shader module **must** be less than or equal
  to this limit.

- <span id="limits-maxCullDistances"></span> `maxCullDistances` is the
  maximum number of cull distances that **can** be used in a single
  shader stage. The size of any array declared with the `CullDistance`
  built-in decoration in a shader module **must** be less than or equal
  to this limit.

- <span id="limits-maxCombinedClipAndCullDistances"></span>
  `maxCombinedClipAndCullDistances` is the maximum combined number of
  clip and cull distances that **can** be used in a single shader stage.
  The sum of the sizes of all arrays declared with the `ClipDistance`
  and `CullDistance` built-in decoration used by a single shader stage
  in a shader module **must** be less than or equal to this limit.

- <span id="limits-discreteQueuePriorities"></span>
  `discreteQueuePriorities` is the number of discrete priorities that
  **can** be assigned to a queue based on the value of each member of
  [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateInfo.html)::`pQueuePriorities`.
  This **must** be at least 2, and levels **must** be spread evenly over
  the range, with at least one level at 1.0, and another at 0.0. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-priority"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-priority</a>.

- <span id="limits-pointSizeRange"></span> `pointSizeRange`\[2\] is the
  range \[`minimum`,`maximum`\] of supported sizes for points. Values
  written to variables decorated with the `PointSize` built-in
  decoration are clamped to this range.

- <span id="limits-lineWidthRange"></span> `lineWidthRange`\[2\] is the
  range \[`minimum`,`maximum`\] of supported widths for lines. Values
  specified by the `lineWidth` member of the
  [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)
  or the `lineWidth` parameter to `vkCmdSetLineWidth` are clamped to
  this range.

- <span id="limits-pointSizeGranularity"></span> `pointSizeGranularity`
  is the granularity of supported point sizes. Not all point sizes in
  the range defined by `pointSizeRange` are supported. This limit
  specifies the granularity (or increment) between successive supported
  point sizes.

- <span id="limits-lineWidthGranularity"></span> `lineWidthGranularity`
  is the granularity of supported line widths. Not all line widths in
  the range defined by `lineWidthRange` are supported. This limit
  specifies the granularity (or increment) between successive supported
  line widths.

- <span id="limits-strictLines"></span> `strictLines` specifies whether
  lines are rasterized according to the preferred method of
  rasterization. If set to `VK_FALSE`, lines **may** be rasterized under
  a relaxed set of rules. If set to `VK_TRUE`, lines are rasterized as
  per the strict definition. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-lines-basic"
  target="_blank" rel="noopener">Basic Line Segment Rasterization</a>.

- <span id="limits-standardSampleLocations"></span>
  `standardSampleLocations` specifies whether rasterization uses the
  standard sample locations as documented in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-multisampling"
  target="_blank" rel="noopener">Multisampling</a>. If set to `VK_TRUE`,
  the implementation uses the documented sample locations. If set to
  `VK_FALSE`, the implementation **may** use different sample locations.

- <span id="limits-optimalBufferCopyOffsetAlignment"></span>
  `optimalBufferCopyOffsetAlignment` is the optimal buffer offset
  alignment in bytes for
  [vkCmdCopyBufferToImage2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyBufferToImage2.html),
  [vkCmdCopyBufferToImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyBufferToImage.html),
  [vkCmdCopyImageToBuffer2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyImageToBuffer2.html), and
  [vkCmdCopyImageToBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyImageToBuffer.html). This value is
  also the optimal host memory offset alignment in bytes for
  [vkCopyMemoryToImageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyMemoryToImageEXT.html) and
  [vkCopyImageToMemoryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyImageToMemoryEXT.html). The per texel
  alignment requirements are enforced, but applications **should** use
  the optimal alignment for optimal performance and power use. The value
  **must** be a power of two.

- <span id="limits-optimalBufferCopyRowPitchAlignment"></span>
  `optimalBufferCopyRowPitchAlignment` is the optimal buffer row pitch
  alignment in bytes for
  [vkCmdCopyBufferToImage2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyBufferToImage2.html),
  [vkCmdCopyBufferToImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyBufferToImage.html),
  [vkCmdCopyImageToBuffer2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyImageToBuffer2.html), and
  [vkCmdCopyImageToBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyImageToBuffer.html). This value is
  also the optimal host memory row pitch alignment in bytes for
  [vkCopyMemoryToImageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyMemoryToImageEXT.html) and
  [vkCopyImageToMemoryEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCopyImageToMemoryEXT.html). Row pitch is
  the number of bytes between texels with the same X coordinate in
  adjacent rows (Y coordinates differ by one). The per texel alignment
  requirements are enforced, but applications **should** use the optimal
  alignment for optimal performance and power use. The value **must** be
  a power of two.

- <span id="limits-nonCoherentAtomSize"></span> `nonCoherentAtomSize` is
  the size and alignment in bytes that bounds concurrent access to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-device-hostaccess"
  target="_blank" rel="noopener">host-mapped device memory</a>. The
  value **must** be a power of two.

  1  
  For all bitmasks of
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html), the sample count
  limits defined above represent the minimum supported sample counts for
  each image type. Individual images **may** support additional sample
  counts, which are queried using
  [vkGetPhysicalDeviceImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties.html)
  as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-supported-sample-counts"
  target="_blank" rel="noopener">Supported Sample Counts</a>.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html),
[VkSampleCountFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceLimits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
