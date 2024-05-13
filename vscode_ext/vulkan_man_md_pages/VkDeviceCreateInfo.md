# VkDeviceCreateInfo(3) Manual Page

## Name

VkDeviceCreateInfo - Structure specifying parameters of a newly created
device



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDeviceCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkDeviceCreateInfo {
    VkStructureType                    sType;
    const void*                        pNext;
    VkDeviceCreateFlags                flags;
    uint32_t                           queueCreateInfoCount;
    const VkDeviceQueueCreateInfo*     pQueueCreateInfos;
    uint32_t                           enabledLayerCount;
    const char* const*                 ppEnabledLayerNames;
    uint32_t                           enabledExtensionCount;
    const char* const*                 ppEnabledExtensionNames;
    const VkPhysicalDeviceFeatures*    pEnabledFeatures;
} VkDeviceCreateInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `queueCreateInfoCount` is the unsigned integer size of the
  `pQueueCreateInfos` array. Refer to the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-queue-creation"
  target="_blank" rel="noopener">Queue Creation</a> section below for
  further details.

- `pQueueCreateInfos` is a pointer to an array of
  [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateInfo.html) structures
  describing the queues that are requested to be created along with the
  logical device. Refer to the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-queue-creation"
  target="_blank" rel="noopener">Queue Creation</a> section below for
  further details.

- `enabledLayerCount` is deprecated and ignored.

- `ppEnabledLayerNames` is deprecated and ignored. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-layers-devicelayerdeprecation"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-layers-devicelayerdeprecation</a>.

- `enabledExtensionCount` is the number of device extensions to enable.

- `ppEnabledExtensionNames` is a pointer to an array of
  `enabledExtensionCount` null-terminated UTF-8 strings containing the
  names of extensions to enable for the created device. See the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-extensions"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-extensions</a>
  section for further details.

- `pEnabledFeatures` is `NULL` or a pointer to a
  [VkPhysicalDeviceFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures.html) structure
  containing boolean indicators of all the features to be enabled. Refer
  to the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features"
  target="_blank" rel="noopener">Features</a> section for further
  details.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkDeviceCreateInfo-queueFamilyIndex-02802"
  id="VUID-VkDeviceCreateInfo-queueFamilyIndex-02802"></a>
  VUID-VkDeviceCreateInfo-queueFamilyIndex-02802  
  The `queueFamilyIndex` member of each element of `pQueueCreateInfos`
  **must** be unique within `pQueueCreateInfos` , except that two
  members can share the same `queueFamilyIndex` if one describes
  protected-capable queues and one describes queues that are not
  protected-capable

- <a href="#VUID-VkDeviceCreateInfo-pQueueCreateInfos-06755"
  id="VUID-VkDeviceCreateInfo-pQueueCreateInfos-06755"></a>
  VUID-VkDeviceCreateInfo-pQueueCreateInfos-06755  
  If multiple elements of `pQueueCreateInfos` share the same
  `queueFamilyIndex`, the sum of their `queueCount` members **must** be
  less than or equal to the `queueCount` member of the
  `VkQueueFamilyProperties` structure, as returned by
  `vkGetPhysicalDeviceQueueFamilyProperties` in the
  `pQueueFamilyProperties`\[queueFamilyIndex\]

- <a href="#VUID-VkDeviceCreateInfo-pQueueCreateInfos-06654"
  id="VUID-VkDeviceCreateInfo-pQueueCreateInfos-06654"></a>
  VUID-VkDeviceCreateInfo-pQueueCreateInfos-06654  
  If multiple elements of `pQueueCreateInfos` share the same
  `queueFamilyIndex`, then all of such elements **must** have the same
  global priority level, which **can** be specified explicitly by the
  including a
  [VkDeviceQueueGlobalPriorityCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueGlobalPriorityCreateInfoKHR.html)
  structure in the `pNext` chain, or by the implicit default value

- <a href="#VUID-VkDeviceCreateInfo-pNext-00373"
  id="VUID-VkDeviceCreateInfo-pNext-00373"></a>
  VUID-VkDeviceCreateInfo-pNext-00373  
  If the `pNext` chain includes a
  [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure,
  then `pEnabledFeatures` **must** be `NULL`

- <a href="#VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-01840"
  id="VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-01840"></a>
  VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-01840  
  If
  [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html)::`apiVersion`
  advertises Vulkan 1.1 or later, `ppEnabledExtensionNames` **must** not
  contain
  [`VK_AMD_negative_viewport_height`](VK_AMD_negative_viewport_height.html)

- <a href="#VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-00374"
  id="VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-00374"></a>
  VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-00374  
  `ppEnabledExtensionNames` **must** not contain both
  [`VK_KHR_maintenance1`](VK_KHR_maintenance1.html) and
  [`VK_AMD_negative_viewport_height`](VK_AMD_negative_viewport_height.html)

- <a href="#VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-03328"
  id="VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-03328"></a>
  VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-03328  
  `ppEnabledExtensionNames` **must** not contain both
  [`VK_KHR_buffer_device_address`](VK_KHR_buffer_device_address.html)
  and
  [`VK_EXT_buffer_device_address`](VK_EXT_buffer_device_address.html)

- <a href="#VUID-VkDeviceCreateInfo-pNext-04748"
  id="VUID-VkDeviceCreateInfo-pNext-04748"></a>
  VUID-VkDeviceCreateInfo-pNext-04748  
  If the `pNext` chain includes a
  [VkPhysicalDeviceVulkan12Features](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkan12Features.html)
  structure and
  [VkPhysicalDeviceVulkan12Features](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkan12Features.html)::`bufferDeviceAddress`
  is `VK_TRUE`, `ppEnabledExtensionNames` **must** not contain
  [`VK_EXT_buffer_device_address`](VK_EXT_buffer_device_address.html)

- <a href="#VUID-VkDeviceCreateInfo-pNext-02829"
  id="VUID-VkDeviceCreateInfo-pNext-02829"></a>
  VUID-VkDeviceCreateInfo-pNext-02829  
  If the `pNext` chain includes a
  [VkPhysicalDeviceVulkan11Features](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkan11Features.html)
  structure, then it **must** not include a
  [VkPhysicalDevice16BitStorageFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice16BitStorageFeatures.html),
  [VkPhysicalDeviceMultiviewFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMultiviewFeatures.html),
  [VkPhysicalDeviceVariablePointersFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVariablePointersFeatures.html),
  [VkPhysicalDeviceProtectedMemoryFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProtectedMemoryFeatures.html),
  [VkPhysicalDeviceSamplerYcbcrConversionFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSamplerYcbcrConversionFeatures.html),
  or
  [VkPhysicalDeviceShaderDrawParametersFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderDrawParametersFeatures.html)
  structure

- <a href="#VUID-VkDeviceCreateInfo-pNext-02830"
  id="VUID-VkDeviceCreateInfo-pNext-02830"></a>
  VUID-VkDeviceCreateInfo-pNext-02830  
  If the `pNext` chain includes a
  [VkPhysicalDeviceVulkan12Features](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkan12Features.html)
  structure, then it **must** not include a
  [VkPhysicalDevice8BitStorageFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice8BitStorageFeatures.html),
  [VkPhysicalDeviceShaderAtomicInt64Features](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderAtomicInt64Features.html),
  [VkPhysicalDeviceShaderFloat16Int8Features](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderFloat16Int8Features.html),
  [VkPhysicalDeviceDescriptorIndexingFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html),
  [VkPhysicalDeviceScalarBlockLayoutFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceScalarBlockLayoutFeatures.html),
  [VkPhysicalDeviceImagelessFramebufferFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImagelessFramebufferFeatures.html),
  [VkPhysicalDeviceUniformBufferStandardLayoutFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceUniformBufferStandardLayoutFeatures.html),
  [VkPhysicalDeviceShaderSubgroupExtendedTypesFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderSubgroupExtendedTypesFeatures.html),
  [VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures.html),
  [VkPhysicalDeviceHostQueryResetFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceHostQueryResetFeatures.html),
  [VkPhysicalDeviceTimelineSemaphoreFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTimelineSemaphoreFeatures.html),
  [VkPhysicalDeviceBufferDeviceAddressFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceBufferDeviceAddressFeatures.html),
  or
  [VkPhysicalDeviceVulkanMemoryModelFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkanMemoryModelFeatures.html)
  structure

- <a href="#VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-04476"
  id="VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-04476"></a>
  VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-04476  
  If `ppEnabledExtensionNames` contains
  `"VK_KHR_shader_draw_parameters"` and the `pNext` chain includes a
  [VkPhysicalDeviceVulkan11Features](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkan11Features.html)
  structure, then
  `VkPhysicalDeviceVulkan11Features`::`shaderDrawParameters` **must** be
  `VK_TRUE`

- <a href="#VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-02831"
  id="VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-02831"></a>
  VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-02831  
  If `ppEnabledExtensionNames` contains `"VK_KHR_draw_indirect_count"`
  and the `pNext` chain includes a
  [VkPhysicalDeviceVulkan12Features](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkan12Features.html)
  structure, then
  `VkPhysicalDeviceVulkan12Features`::`drawIndirectCount` **must** be
  `VK_TRUE`

- <a href="#VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-02832"
  id="VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-02832"></a>
  VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-02832  
  If `ppEnabledExtensionNames` contains
  `"VK_KHR_sampler_mirror_clamp_to_edge"` and the `pNext` chain includes
  a
  [VkPhysicalDeviceVulkan12Features](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkan12Features.html)
  structure, then
  `VkPhysicalDeviceVulkan12Features`::`samplerMirrorClampToEdge`
  **must** be `VK_TRUE`

- <a href="#VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-02833"
  id="VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-02833"></a>
  VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-02833  
  If `ppEnabledExtensionNames` contains `"VK_EXT_descriptor_indexing"`
  and the `pNext` chain includes a
  [VkPhysicalDeviceVulkan12Features](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkan12Features.html)
  structure, then
  `VkPhysicalDeviceVulkan12Features`::`descriptorIndexing` **must** be
  `VK_TRUE`

- <a href="#VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-02834"
  id="VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-02834"></a>
  VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-02834  
  If `ppEnabledExtensionNames` contains `"VK_EXT_sampler_filter_minmax"`
  and the `pNext` chain includes a
  [VkPhysicalDeviceVulkan12Features](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkan12Features.html)
  structure, then
  `VkPhysicalDeviceVulkan12Features`::`samplerFilterMinmax` **must** be
  `VK_TRUE`

- <a href="#VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-02835"
  id="VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-02835"></a>
  VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-02835  
  If `ppEnabledExtensionNames` contains
  `"VK_EXT_shader_viewport_index_layer"` and the `pNext` chain includes
  a
  [VkPhysicalDeviceVulkan12Features](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkan12Features.html)
  structure, then
  `VkPhysicalDeviceVulkan12Features`::`shaderOutputViewportIndex` and
  `VkPhysicalDeviceVulkan12Features`::`shaderOutputLayer` **must** both
  be `VK_TRUE`

- <a href="#VUID-VkDeviceCreateInfo-pNext-06532"
  id="VUID-VkDeviceCreateInfo-pNext-06532"></a>
  VUID-VkDeviceCreateInfo-pNext-06532  
  If the `pNext` chain includes a
  [VkPhysicalDeviceVulkan13Features](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkan13Features.html)
  structure, then it **must** not include a
  [VkPhysicalDeviceDynamicRenderingFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDynamicRenderingFeatures.html),
  [VkPhysicalDeviceImageRobustnessFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageRobustnessFeatures.html),
  [VkPhysicalDeviceInlineUniformBlockFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceInlineUniformBlockFeatures.html),
  [VkPhysicalDeviceMaintenance4Features](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance4Features.html),
  [VkPhysicalDevicePipelineCreationCacheControlFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePipelineCreationCacheControlFeatures.html),
  [VkPhysicalDevicePrivateDataFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePrivateDataFeatures.html),
  [VkPhysicalDeviceShaderDemoteToHelperInvocationFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderDemoteToHelperInvocationFeatures.html),
  [VkPhysicalDeviceShaderIntegerDotProductFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderIntegerDotProductFeatures.html),
  [VkPhysicalDeviceShaderTerminateInvocationFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderTerminateInvocationFeatures.html),
  [VkPhysicalDeviceSubgroupSizeControlFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSubgroupSizeControlFeatures.html),
  [VkPhysicalDeviceSynchronization2Features](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSynchronization2Features.html),
  [VkPhysicalDeviceTextureCompressionASTCHDRFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTextureCompressionASTCHDRFeatures.html),
  or
  [VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeatures.html)
  structure

- <a href="#VUID-VkDeviceCreateInfo-pProperties-04451"
  id="VUID-VkDeviceCreateInfo-pProperties-04451"></a>
  VUID-VkDeviceCreateInfo-pProperties-04451  
  If the [`VK_KHR_portability_subset`](VK_KHR_portability_subset.html)
  extension is included in `pProperties` of
  [vkEnumerateDeviceExtensionProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkEnumerateDeviceExtensionProperties.html),
  `ppEnabledExtensionNames` **must** include
  `"VK_KHR_portability_subset"`

- <a href="#VUID-VkDeviceCreateInfo-shadingRateImage-04478"
  id="VUID-VkDeviceCreateInfo-shadingRateImage-04478"></a>
  VUID-VkDeviceCreateInfo-shadingRateImage-04478  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shadingRateImage"
  target="_blank" rel="noopener"><code>shadingRateImage</code></a>
  feature is enabled, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelineFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>pipelineFragmentShadingRate</code></a> feature
  **must** not be enabled

- <a href="#VUID-VkDeviceCreateInfo-shadingRateImage-04479"
  id="VUID-VkDeviceCreateInfo-shadingRateImage-04479"></a>
  VUID-VkDeviceCreateInfo-shadingRateImage-04479  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shadingRateImage"
  target="_blank" rel="noopener"><code>shadingRateImage</code></a>
  feature is enabled, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-primitiveFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>primitiveFragmentShadingRate</code></a> feature
  **must** not be enabled

- <a href="#VUID-VkDeviceCreateInfo-shadingRateImage-04480"
  id="VUID-VkDeviceCreateInfo-shadingRateImage-04480"></a>
  VUID-VkDeviceCreateInfo-shadingRateImage-04480  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shadingRateImage"
  target="_blank" rel="noopener"><code>shadingRateImage</code></a>
  feature is enabled, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-attachmentFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>attachmentFragmentShadingRate</code></a> feature
  **must** not be enabled

- <a href="#VUID-VkDeviceCreateInfo-fragmentDensityMap-04481"
  id="VUID-VkDeviceCreateInfo-fragmentDensityMap-04481"></a>
  VUID-VkDeviceCreateInfo-fragmentDensityMap-04481  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-fragmentDensityMap"
  target="_blank" rel="noopener"><code>fragmentDensityMap</code></a>
  feature is enabled, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelineFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>pipelineFragmentShadingRate</code></a> feature
  **must** not be enabled

- <a href="#VUID-VkDeviceCreateInfo-fragmentDensityMap-04482"
  id="VUID-VkDeviceCreateInfo-fragmentDensityMap-04482"></a>
  VUID-VkDeviceCreateInfo-fragmentDensityMap-04482  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-fragmentDensityMap"
  target="_blank" rel="noopener"><code>fragmentDensityMap</code></a>
  feature is enabled, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-primitiveFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>primitiveFragmentShadingRate</code></a> feature
  **must** not be enabled

- <a href="#VUID-VkDeviceCreateInfo-fragmentDensityMap-04483"
  id="VUID-VkDeviceCreateInfo-fragmentDensityMap-04483"></a>
  VUID-VkDeviceCreateInfo-fragmentDensityMap-04483  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-fragmentDensityMap"
  target="_blank" rel="noopener"><code>fragmentDensityMap</code></a>
  feature is enabled, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-attachmentFragmentShadingRate"
  target="_blank"
  rel="noopener"><code>attachmentFragmentShadingRate</code></a> feature
  **must** not be enabled

- <a href="#VUID-VkDeviceCreateInfo-None-04896"
  id="VUID-VkDeviceCreateInfo-None-04896"></a>
  VUID-VkDeviceCreateInfo-None-04896  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-sparseImageInt64Atomics"
  target="_blank" rel="noopener"><code>sparseImageInt64Atomics</code></a>
  is enabled, <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderImageInt64Atomics"
  target="_blank" rel="noopener"><code>shaderImageInt64Atomics</code></a>
  **must** be enabled

- <a href="#VUID-VkDeviceCreateInfo-None-04897"
  id="VUID-VkDeviceCreateInfo-None-04897"></a>
  VUID-VkDeviceCreateInfo-None-04897  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-sparseImageFloat32Atomics"
  target="_blank"
  rel="noopener"><code>sparseImageFloat32Atomics</code></a> is enabled,
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderImageFloat32Atomics"
  target="_blank"
  rel="noopener"><code>shaderImageFloat32Atomics</code></a> **must** be
  enabled

- <a href="#VUID-VkDeviceCreateInfo-None-04898"
  id="VUID-VkDeviceCreateInfo-None-04898"></a>
  VUID-VkDeviceCreateInfo-None-04898  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-sparseImageFloat32AtomicAdd"
  target="_blank"
  rel="noopener"><code>sparseImageFloat32AtomicAdd</code></a> is
  enabled, <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderImageFloat32AtomicAdd"
  target="_blank"
  rel="noopener"><code>shaderImageFloat32AtomicAdd</code></a> **must**
  be enabled

- <a href="#VUID-VkDeviceCreateInfo-sparseImageFloat32AtomicMinMax-04975"
  id="VUID-VkDeviceCreateInfo-sparseImageFloat32AtomicMinMax-04975"></a>
  VUID-VkDeviceCreateInfo-sparseImageFloat32AtomicMinMax-04975  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-sparseImageFloat32AtomicMinMax"
  target="_blank"
  rel="noopener"><code>sparseImageFloat32AtomicMinMax</code></a> is
  enabled, <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderImageFloat32AtomicMinMax"
  target="_blank"
  rel="noopener"><code>shaderImageFloat32AtomicMinMax</code></a>
  **must** be enabled

- <a href="#VUID-VkDeviceCreateInfo-None-08095"
  id="VUID-VkDeviceCreateInfo-None-08095"></a>
  VUID-VkDeviceCreateInfo-None-08095  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-descriptorBuffer"
  target="_blank" rel="noopener"><code>descriptorBuffer</code></a> is
  enabled, `ppEnabledExtensionNames` **must** not contain
  [`VK_AMD_shader_fragment_mask`](VK_AMD_shader_fragment_mask.html)

- <a href="#VUID-VkDeviceCreateInfo-pNext-09396"
  id="VUID-VkDeviceCreateInfo-pNext-09396"></a>
  VUID-VkDeviceCreateInfo-pNext-09396  
  If the `pNext` chain includes a
  [VkDeviceQueueShaderCoreControlCreateInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueShaderCoreControlCreateInfoARM.html)
  structure, then it **must** not be included in the `pNext` chain of
  any of the [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateInfo.html)
  structures in `pQueueCreateInfos`

- <a href="#VUID-VkDeviceCreateInfo-pNext-09397"
  id="VUID-VkDeviceCreateInfo-pNext-09397"></a>
  VUID-VkDeviceCreateInfo-pNext-09397  
  If the `pNext` chain includes a
  [VkDeviceQueueShaderCoreControlCreateInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueShaderCoreControlCreateInfoARM.html)
  structure then
  [VkPhysicalDeviceSchedulingControlsPropertiesARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSchedulingControlsPropertiesARM.html)::`schedulingControlsFlags`
  **must** contain
  `VK_PHYSICAL_DEVICE_SCHEDULING_CONTROLS_SHADER_CORE_COUNT_ARM`

Valid Usage (Implicit)

- <a href="#VUID-VkDeviceCreateInfo-sType-sType"
  id="VUID-VkDeviceCreateInfo-sType-sType"></a>
  VUID-VkDeviceCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO`

- <a href="#VUID-VkDeviceCreateInfo-pNext-pNext"
  id="VUID-VkDeviceCreateInfo-pNext-pNext"></a>
  VUID-VkDeviceCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkDeviceDeviceMemoryReportCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceDeviceMemoryReportCreateInfoEXT.html),
  [VkDeviceDiagnosticsConfigCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceDiagnosticsConfigCreateInfoNV.html),
  [VkDeviceGroupDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupDeviceCreateInfo.html),
  [VkDeviceMemoryOverallocationCreateInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemoryOverallocationCreateInfoAMD.html),
  [VkDevicePrivateDataCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevicePrivateDataCreateInfo.html),
  [VkDeviceQueueShaderCoreControlCreateInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueShaderCoreControlCreateInfoARM.html),
  [VkPhysicalDevice16BitStorageFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice16BitStorageFeatures.html),
  [VkPhysicalDevice4444FormatsFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice4444FormatsFeaturesEXT.html),
  [VkPhysicalDevice8BitStorageFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice8BitStorageFeatures.html),
  [VkPhysicalDeviceASTCDecodeFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceASTCDecodeFeaturesEXT.html),
  [VkPhysicalDeviceAccelerationStructureFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceAccelerationStructureFeaturesKHR.html),
  [VkPhysicalDeviceAddressBindingReportFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceAddressBindingReportFeaturesEXT.html),
  [VkPhysicalDeviceAmigoProfilingFeaturesSEC](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceAmigoProfilingFeaturesSEC.html),
  [VkPhysicalDeviceAttachmentFeedbackLoopDynamicStateFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceAttachmentFeedbackLoopDynamicStateFeaturesEXT.html),
  [VkPhysicalDeviceAttachmentFeedbackLoopLayoutFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceAttachmentFeedbackLoopLayoutFeaturesEXT.html),
  [VkPhysicalDeviceBlendOperationAdvancedFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceBlendOperationAdvancedFeaturesEXT.html),
  [VkPhysicalDeviceBorderColorSwizzleFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceBorderColorSwizzleFeaturesEXT.html),
  [VkPhysicalDeviceBufferDeviceAddressFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceBufferDeviceAddressFeatures.html),
  [VkPhysicalDeviceBufferDeviceAddressFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceBufferDeviceAddressFeaturesEXT.html),
  [VkPhysicalDeviceClusterCullingShaderFeaturesHUAWEI](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceClusterCullingShaderFeaturesHUAWEI.html),
  [VkPhysicalDeviceCoherentMemoryFeaturesAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCoherentMemoryFeaturesAMD.html),
  [VkPhysicalDeviceColorWriteEnableFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceColorWriteEnableFeaturesEXT.html),
  [VkPhysicalDeviceComputeShaderDerivativesFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceComputeShaderDerivativesFeaturesNV.html),
  [VkPhysicalDeviceConditionalRenderingFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceConditionalRenderingFeaturesEXT.html),
  [VkPhysicalDeviceCooperativeMatrixFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCooperativeMatrixFeaturesKHR.html),
  [VkPhysicalDeviceCooperativeMatrixFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCooperativeMatrixFeaturesNV.html),
  [VkPhysicalDeviceCopyMemoryIndirectFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCopyMemoryIndirectFeaturesNV.html),
  [VkPhysicalDeviceCornerSampledImageFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCornerSampledImageFeaturesNV.html),
  [VkPhysicalDeviceCoverageReductionModeFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCoverageReductionModeFeaturesNV.html),
  [VkPhysicalDeviceCubicClampFeaturesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCubicClampFeaturesQCOM.html),
  [VkPhysicalDeviceCubicWeightsFeaturesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCubicWeightsFeaturesQCOM.html),
  [VkPhysicalDeviceCudaKernelLaunchFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCudaKernelLaunchFeaturesNV.html),
  [VkPhysicalDeviceCustomBorderColorFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCustomBorderColorFeaturesEXT.html),
  [VkPhysicalDeviceDedicatedAllocationImageAliasingFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDedicatedAllocationImageAliasingFeaturesNV.html),
  [VkPhysicalDeviceDepthBiasControlFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDepthBiasControlFeaturesEXT.html),
  [VkPhysicalDeviceDepthClampZeroOneFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDepthClampZeroOneFeaturesEXT.html),
  [VkPhysicalDeviceDepthClipControlFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDepthClipControlFeaturesEXT.html),
  [VkPhysicalDeviceDepthClipEnableFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDepthClipEnableFeaturesEXT.html),
  [VkPhysicalDeviceDescriptorBufferFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferFeaturesEXT.html),
  [VkPhysicalDeviceDescriptorIndexingFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html),
  [VkPhysicalDeviceDescriptorPoolOverallocationFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorPoolOverallocationFeaturesNV.html),
  [VkPhysicalDeviceDescriptorSetHostMappingFeaturesVALVE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorSetHostMappingFeaturesVALVE.html),
  [VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV.html),
  [VkPhysicalDeviceDeviceGeneratedCommandsFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDeviceGeneratedCommandsFeaturesNV.html),
  [VkPhysicalDeviceDeviceMemoryReportFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDeviceMemoryReportFeaturesEXT.html),
  [VkPhysicalDeviceDiagnosticsConfigFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDiagnosticsConfigFeaturesNV.html),
  [VkPhysicalDeviceDisplacementMicromapFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDisplacementMicromapFeaturesNV.html),
  [VkPhysicalDeviceDynamicRenderingFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDynamicRenderingFeatures.html),
  [VkPhysicalDeviceDynamicRenderingLocalReadFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDynamicRenderingLocalReadFeaturesKHR.html),
  [VkPhysicalDeviceDynamicRenderingUnusedAttachmentsFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDynamicRenderingUnusedAttachmentsFeaturesEXT.html),
  [VkPhysicalDeviceExclusiveScissorFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExclusiveScissorFeaturesNV.html),
  [VkPhysicalDeviceExtendedDynamicState2FeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExtendedDynamicState2FeaturesEXT.html),
  [VkPhysicalDeviceExtendedDynamicState3FeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExtendedDynamicState3FeaturesEXT.html),
  [VkPhysicalDeviceExtendedDynamicStateFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExtendedDynamicStateFeaturesEXT.html),
  [VkPhysicalDeviceExtendedSparseAddressSpaceFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExtendedSparseAddressSpaceFeaturesNV.html),
  [VkPhysicalDeviceExternalFormatResolveFeaturesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalFormatResolveFeaturesANDROID.html),
  [VkPhysicalDeviceExternalMemoryRDMAFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalMemoryRDMAFeaturesNV.html),
  [VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX.html),
  [VkPhysicalDeviceFaultFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFaultFeaturesEXT.html),
  [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkPhysicalDeviceFragmentDensityMap2FeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentDensityMap2FeaturesEXT.html),
  [VkPhysicalDeviceFragmentDensityMapFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentDensityMapFeaturesEXT.html),
  [VkPhysicalDeviceFragmentDensityMapOffsetFeaturesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentDensityMapOffsetFeaturesQCOM.html),
  [VkPhysicalDeviceFragmentShaderBarycentricFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentShaderBarycentricFeaturesKHR.html),
  [VkPhysicalDeviceFragmentShaderInterlockFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentShaderInterlockFeaturesEXT.html),
  [VkPhysicalDeviceFragmentShadingRateEnumsFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentShadingRateEnumsFeaturesNV.html),
  [VkPhysicalDeviceFragmentShadingRateFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentShadingRateFeaturesKHR.html),
  [VkPhysicalDeviceFrameBoundaryFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFrameBoundaryFeaturesEXT.html),
  [VkPhysicalDeviceGlobalPriorityQueryFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceGlobalPriorityQueryFeaturesKHR.html),
  [VkPhysicalDeviceGraphicsPipelineLibraryFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceGraphicsPipelineLibraryFeaturesEXT.html),
  [VkPhysicalDeviceHostImageCopyFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceHostImageCopyFeaturesEXT.html),
  [VkPhysicalDeviceHostQueryResetFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceHostQueryResetFeatures.html),
  [VkPhysicalDeviceImage2DViewOf3DFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImage2DViewOf3DFeaturesEXT.html),
  [VkPhysicalDeviceImageAlignmentControlFeaturesMESA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageAlignmentControlFeaturesMESA.html),
  [VkPhysicalDeviceImageCompressionControlFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageCompressionControlFeaturesEXT.html),
  [VkPhysicalDeviceImageCompressionControlSwapchainFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageCompressionControlSwapchainFeaturesEXT.html),
  [VkPhysicalDeviceImageProcessing2FeaturesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageProcessing2FeaturesQCOM.html),
  [VkPhysicalDeviceImageProcessingFeaturesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageProcessingFeaturesQCOM.html),
  [VkPhysicalDeviceImageRobustnessFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageRobustnessFeatures.html),
  [VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageSlicedViewOf3DFeaturesEXT.html),
  [VkPhysicalDeviceImageViewMinLodFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageViewMinLodFeaturesEXT.html),
  [VkPhysicalDeviceImagelessFramebufferFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImagelessFramebufferFeatures.html),
  [VkPhysicalDeviceIndexTypeUint8FeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceIndexTypeUint8FeaturesKHR.html),
  [VkPhysicalDeviceInheritedViewportScissorFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceInheritedViewportScissorFeaturesNV.html),
  [VkPhysicalDeviceInlineUniformBlockFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceInlineUniformBlockFeatures.html),
  [VkPhysicalDeviceInvocationMaskFeaturesHUAWEI](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceInvocationMaskFeaturesHUAWEI.html),
  [VkPhysicalDeviceLegacyDitheringFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLegacyDitheringFeaturesEXT.html),
  [VkPhysicalDeviceLegacyVertexAttributesFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLegacyVertexAttributesFeaturesEXT.html),
  [VkPhysicalDeviceLineRasterizationFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLineRasterizationFeaturesKHR.html),
  [VkPhysicalDeviceLinearColorAttachmentFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLinearColorAttachmentFeaturesNV.html),
  [VkPhysicalDeviceMaintenance4Features](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance4Features.html),
  [VkPhysicalDeviceMaintenance5FeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance5FeaturesKHR.html),
  [VkPhysicalDeviceMaintenance6FeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance6FeaturesKHR.html),
  [VkPhysicalDeviceMapMemoryPlacedFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMapMemoryPlacedFeaturesEXT.html),
  [VkPhysicalDeviceMemoryDecompressionFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryDecompressionFeaturesNV.html),
  [VkPhysicalDeviceMemoryPriorityFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryPriorityFeaturesEXT.html),
  [VkPhysicalDeviceMeshShaderFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMeshShaderFeaturesEXT.html),
  [VkPhysicalDeviceMeshShaderFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMeshShaderFeaturesNV.html),
  [VkPhysicalDeviceMultiDrawFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMultiDrawFeaturesEXT.html),
  [VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMultisampledRenderToSingleSampledFeaturesEXT.html),
  [VkPhysicalDeviceMultiviewFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMultiviewFeatures.html),
  [VkPhysicalDeviceMultiviewPerViewRenderAreasFeaturesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMultiviewPerViewRenderAreasFeaturesQCOM.html),
  [VkPhysicalDeviceMultiviewPerViewViewportsFeaturesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMultiviewPerViewViewportsFeaturesQCOM.html),
  [VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT.html),
  [VkPhysicalDeviceNestedCommandBufferFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceNestedCommandBufferFeaturesEXT.html),
  [VkPhysicalDeviceNonSeamlessCubeMapFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceNonSeamlessCubeMapFeaturesEXT.html),
  [VkPhysicalDeviceOpacityMicromapFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceOpacityMicromapFeaturesEXT.html),
  [VkPhysicalDeviceOpticalFlowFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceOpticalFlowFeaturesNV.html),
  [VkPhysicalDevicePageableDeviceLocalMemoryFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePageableDeviceLocalMemoryFeaturesEXT.html),
  [VkPhysicalDevicePerStageDescriptorSetFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePerStageDescriptorSetFeaturesNV.html),
  [VkPhysicalDevicePerformanceQueryFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePerformanceQueryFeaturesKHR.html),
  [VkPhysicalDevicePipelineCreationCacheControlFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePipelineCreationCacheControlFeatures.html),
  [VkPhysicalDevicePipelineExecutablePropertiesFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePipelineExecutablePropertiesFeaturesKHR.html),
  [VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT.html),
  [VkPhysicalDevicePipelinePropertiesFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePipelinePropertiesFeaturesEXT.html),
  [VkPhysicalDevicePipelineProtectedAccessFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePipelineProtectedAccessFeaturesEXT.html),
  [VkPhysicalDevicePipelineRobustnessFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePipelineRobustnessFeaturesEXT.html),
  [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html),
  [VkPhysicalDevicePresentBarrierFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePresentBarrierFeaturesNV.html),
  [VkPhysicalDevicePresentIdFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePresentIdFeaturesKHR.html),
  [VkPhysicalDevicePresentWaitFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePresentWaitFeaturesKHR.html),
  [VkPhysicalDevicePrimitiveTopologyListRestartFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePrimitiveTopologyListRestartFeaturesEXT.html),
  [VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePrimitivesGeneratedQueryFeaturesEXT.html),
  [VkPhysicalDevicePrivateDataFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePrivateDataFeatures.html),
  [VkPhysicalDeviceProtectedMemoryFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProtectedMemoryFeatures.html),
  [VkPhysicalDeviceProvokingVertexFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProvokingVertexFeaturesEXT.html),
  [VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRGBA10X6FormatsFeaturesEXT.html),
  [VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT.html),
  [VkPhysicalDeviceRawAccessChainsFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRawAccessChainsFeaturesNV.html),
  [VkPhysicalDeviceRayQueryFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayQueryFeaturesKHR.html),
  [VkPhysicalDeviceRayTracingInvocationReorderFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingInvocationReorderFeaturesNV.html),
  [VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingMaintenance1FeaturesKHR.html),
  [VkPhysicalDeviceRayTracingMotionBlurFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingMotionBlurFeaturesNV.html),
  [VkPhysicalDeviceRayTracingPipelineFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPipelineFeaturesKHR.html),
  [VkPhysicalDeviceRayTracingPositionFetchFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPositionFetchFeaturesKHR.html),
  [VkPhysicalDeviceRayTracingValidationFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingValidationFeaturesNV.html),
  [VkPhysicalDeviceRelaxedLineRasterizationFeaturesIMG](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRelaxedLineRasterizationFeaturesIMG.html),
  [VkPhysicalDeviceRenderPassStripedFeaturesARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRenderPassStripedFeaturesARM.html),
  [VkPhysicalDeviceRepresentativeFragmentTestFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRepresentativeFragmentTestFeaturesNV.html),
  [VkPhysicalDeviceRobustness2FeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRobustness2FeaturesEXT.html),
  [VkPhysicalDeviceSamplerYcbcrConversionFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSamplerYcbcrConversionFeatures.html),
  [VkPhysicalDeviceScalarBlockLayoutFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceScalarBlockLayoutFeatures.html),
  [VkPhysicalDeviceSchedulingControlsFeaturesARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSchedulingControlsFeaturesARM.html),
  [VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSeparateDepthStencilLayoutsFeatures.html),
  [VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV.html),
  [VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderAtomicFloat2FeaturesEXT.html),
  [VkPhysicalDeviceShaderAtomicFloatFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderAtomicFloatFeaturesEXT.html),
  [VkPhysicalDeviceShaderAtomicInt64Features](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderAtomicInt64Features.html),
  [VkPhysicalDeviceShaderClockFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderClockFeaturesKHR.html),
  [VkPhysicalDeviceShaderCoreBuiltinsFeaturesARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderCoreBuiltinsFeaturesARM.html),
  [VkPhysicalDeviceShaderDemoteToHelperInvocationFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderDemoteToHelperInvocationFeatures.html),
  [VkPhysicalDeviceShaderDrawParametersFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderDrawParametersFeatures.html),
  [VkPhysicalDeviceShaderEarlyAndLateFragmentTestsFeaturesAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderEarlyAndLateFragmentTestsFeaturesAMD.html),
  [VkPhysicalDeviceShaderEnqueueFeaturesAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderEnqueueFeaturesAMDX.html),
  [VkPhysicalDeviceShaderExpectAssumeFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderExpectAssumeFeaturesKHR.html),
  [VkPhysicalDeviceShaderFloat16Int8Features](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderFloat16Int8Features.html),
  [VkPhysicalDeviceShaderFloatControls2FeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderFloatControls2FeaturesKHR.html),
  [VkPhysicalDeviceShaderImageAtomicInt64FeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderImageAtomicInt64FeaturesEXT.html),
  [VkPhysicalDeviceShaderImageFootprintFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderImageFootprintFeaturesNV.html),
  [VkPhysicalDeviceShaderIntegerDotProductFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderIntegerDotProductFeatures.html),
  [VkPhysicalDeviceShaderIntegerFunctions2FeaturesINTEL](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderIntegerFunctions2FeaturesINTEL.html),
  [VkPhysicalDeviceShaderMaximalReconvergenceFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderMaximalReconvergenceFeaturesKHR.html),
  [VkPhysicalDeviceShaderModuleIdentifierFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderModuleIdentifierFeaturesEXT.html),
  [VkPhysicalDeviceShaderObjectFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderObjectFeaturesEXT.html),
  [VkPhysicalDeviceShaderQuadControlFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderQuadControlFeaturesKHR.html),
  [VkPhysicalDeviceShaderSMBuiltinsFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderSMBuiltinsFeaturesNV.html),
  [VkPhysicalDeviceShaderSubgroupExtendedTypesFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderSubgroupExtendedTypesFeatures.html),
  [VkPhysicalDeviceShaderSubgroupRotateFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderSubgroupRotateFeaturesKHR.html),
  [VkPhysicalDeviceShaderSubgroupUniformControlFlowFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderSubgroupUniformControlFlowFeaturesKHR.html),
  [VkPhysicalDeviceShaderTerminateInvocationFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderTerminateInvocationFeatures.html),
  [VkPhysicalDeviceShaderTileImageFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderTileImageFeaturesEXT.html),
  [VkPhysicalDeviceShadingRateImageFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShadingRateImageFeaturesNV.html),
  [VkPhysicalDeviceSubgroupSizeControlFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSubgroupSizeControlFeatures.html),
  [VkPhysicalDeviceSubpassMergeFeedbackFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSubpassMergeFeedbackFeaturesEXT.html),
  [VkPhysicalDeviceSubpassShadingFeaturesHUAWEI](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSubpassShadingFeaturesHUAWEI.html),
  [VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSwapchainMaintenance1FeaturesEXT.html),
  [VkPhysicalDeviceSynchronization2Features](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSynchronization2Features.html),
  [VkPhysicalDeviceTexelBufferAlignmentFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTexelBufferAlignmentFeaturesEXT.html),
  [VkPhysicalDeviceTextureCompressionASTCHDRFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTextureCompressionASTCHDRFeatures.html),
  [VkPhysicalDeviceTilePropertiesFeaturesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTilePropertiesFeaturesQCOM.html),
  [VkPhysicalDeviceTimelineSemaphoreFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTimelineSemaphoreFeatures.html),
  [VkPhysicalDeviceTransformFeedbackFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTransformFeedbackFeaturesEXT.html),
  [VkPhysicalDeviceUniformBufferStandardLayoutFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceUniformBufferStandardLayoutFeatures.html),
  [VkPhysicalDeviceVariablePointersFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVariablePointersFeatures.html),
  [VkPhysicalDeviceVertexAttributeDivisorFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVertexAttributeDivisorFeaturesKHR.html),
  [VkPhysicalDeviceVertexInputDynamicStateFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVertexInputDynamicStateFeaturesEXT.html),
  [VkPhysicalDeviceVideoMaintenance1FeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVideoMaintenance1FeaturesKHR.html),
  [VkPhysicalDeviceVulkan11Features](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkan11Features.html),
  [VkPhysicalDeviceVulkan12Features](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkan12Features.html),
  [VkPhysicalDeviceVulkan13Features](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkan13Features.html),
  [VkPhysicalDeviceVulkanMemoryModelFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkanMemoryModelFeatures.html),
  [VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR.html),
  [VkPhysicalDeviceYcbcr2Plane444FormatsFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceYcbcr2Plane444FormatsFeaturesEXT.html),
  [VkPhysicalDeviceYcbcrDegammaFeaturesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceYcbcrDegammaFeaturesQCOM.html),
  [VkPhysicalDeviceYcbcrImageArraysFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceYcbcrImageArraysFeaturesEXT.html),
  or
  [VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeatures.html)

- <a href="#VUID-VkDeviceCreateInfo-sType-unique"
  id="VUID-VkDeviceCreateInfo-sType-unique"></a>
  VUID-VkDeviceCreateInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique, with the exception of structures of type
  [VkDeviceDeviceMemoryReportCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceDeviceMemoryReportCreateInfoEXT.html)
  or [VkDevicePrivateDataCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevicePrivateDataCreateInfo.html)

- <a href="#VUID-VkDeviceCreateInfo-flags-zerobitmask"
  id="VUID-VkDeviceCreateInfo-flags-zerobitmask"></a>
  VUID-VkDeviceCreateInfo-flags-zerobitmask  
  `flags` **must** be `0`

- <a href="#VUID-VkDeviceCreateInfo-pQueueCreateInfos-parameter"
  id="VUID-VkDeviceCreateInfo-pQueueCreateInfos-parameter"></a>
  VUID-VkDeviceCreateInfo-pQueueCreateInfos-parameter  
  `pQueueCreateInfos` **must** be a valid pointer to an array of
  `queueCreateInfoCount` valid
  [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateInfo.html) structures

- <a href="#VUID-VkDeviceCreateInfo-ppEnabledLayerNames-parameter"
  id="VUID-VkDeviceCreateInfo-ppEnabledLayerNames-parameter"></a>
  VUID-VkDeviceCreateInfo-ppEnabledLayerNames-parameter  
  If `enabledLayerCount` is not `0`, `ppEnabledLayerNames` **must** be a
  valid pointer to an array of `enabledLayerCount` null-terminated UTF-8
  strings

- <a href="#VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-parameter"
  id="VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-parameter"></a>
  VUID-VkDeviceCreateInfo-ppEnabledExtensionNames-parameter  
  If `enabledExtensionCount` is not `0`, `ppEnabledExtensionNames`
  **must** be a valid pointer to an array of `enabledExtensionCount`
  null-terminated UTF-8 strings

- <a href="#VUID-VkDeviceCreateInfo-pEnabledFeatures-parameter"
  id="VUID-VkDeviceCreateInfo-pEnabledFeatures-parameter"></a>
  VUID-VkDeviceCreateInfo-pEnabledFeatures-parameter  
  If `pEnabledFeatures` is not `NULL`, `pEnabledFeatures` **must** be a
  valid pointer to a valid
  [VkPhysicalDeviceFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures.html) structure

- <a href="#VUID-VkDeviceCreateInfo-queueCreateInfoCount-arraylength"
  id="VUID-VkDeviceCreateInfo-queueCreateInfoCount-arraylength"></a>
  VUID-VkDeviceCreateInfo-queueCreateInfoCount-arraylength  
  `queueCreateInfoCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDeviceCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateFlags.html),
[VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateInfo.html),
[VkPhysicalDeviceFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
