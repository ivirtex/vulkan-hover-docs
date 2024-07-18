# RuntimeSpirv(3) Manual Page

## Name

RuntimeSpirv - Runtime SPIR-V Validation



## <a href="#_description" class="anchor"></a>Description

The following rules **must** be validated at runtime. These rules depend
on knowledge of the implementation and its capabilities and knowledge of
runtime information, such as enabled features.

Valid Usage

- <a href="#VUID-RuntimeSpirv-vulkanMemoryModel-06265"
  id="VUID-RuntimeSpirv-vulkanMemoryModel-06265"></a>
  VUID-RuntimeSpirv-vulkanMemoryModel-06265  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-vulkanMemoryModel"
  target="_blank" rel="noopener"><code>vulkanMemoryModel</code></a> is
  enabled and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-vulkanMemoryModelDeviceScope"
  target="_blank"
  rel="noopener"><code>vulkanMemoryModelDeviceScope</code></a> is not
  enabled, `Device` memory scope **must** not be used

- <a href="#VUID-RuntimeSpirv-vulkanMemoryModel-06266"
  id="VUID-RuntimeSpirv-vulkanMemoryModel-06266"></a>
  VUID-RuntimeSpirv-vulkanMemoryModel-06266  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-vulkanMemoryModel"
  target="_blank" rel="noopener"><code>vulkanMemoryModel</code></a> is
  not enabled, `QueueFamily` memory scope **must** not be used

- <a href="#VUID-RuntimeSpirv-shaderSubgroupClock-06267"
  id="VUID-RuntimeSpirv-shaderSubgroupClock-06267"></a>
  VUID-RuntimeSpirv-shaderSubgroupClock-06267  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderSubgroupClock"
  target="_blank" rel="noopener"><code>shaderSubgroupClock</code></a> is
  not enabled, the `Subgroup` scope **must** not be used for
  `OpReadClockKHR`

- <a href="#VUID-RuntimeSpirv-shaderDeviceClock-06268"
  id="VUID-RuntimeSpirv-shaderDeviceClock-06268"></a>
  VUID-RuntimeSpirv-shaderDeviceClock-06268  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderDeviceClock"
  target="_blank" rel="noopener"><code>shaderDeviceClock</code></a> is
  not enabled, the `Device` scope **must** not be used for
  `OpReadClockKHR`

- <a href="#VUID-RuntimeSpirv-None-09558"
  id="VUID-RuntimeSpirv-None-09558"></a> VUID-RuntimeSpirv-None-09558  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-dynamicRenderingLocalRead"
  target="_blank"
  rel="noopener"><code>dynamicRenderingLocalRead</code></a> is not
  enabled, any variable created with a “Type” of `OpTypeImage` that has
  a “Dim” operand of `SubpassData` **must** be decorated with
  `InputAttachmentIndex`

- <a href="#VUID-RuntimeSpirv-OpTypeImage-09644"
  id="VUID-RuntimeSpirv-OpTypeImage-09644"></a>
  VUID-RuntimeSpirv-OpTypeImage-09644  
  Any variable declared as an `OpTypeArray` where the `Element` `Type`
  is an `OpTypeImage` with a “Dim” operand of `SubpassData` **must** be
  decorated with `InputAttachmentIndex`

- <a href="#VUID-RuntimeSpirv-apiVersion-07954"
  id="VUID-RuntimeSpirv-apiVersion-07954"></a>
  VUID-RuntimeSpirv-apiVersion-07954  
  If
  [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html)::`apiVersion`
  is less than Vulkan 1.3, the
  [VK_KHR_format_feature_flags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_format_feature_flags2.html)
  extension is not supported, and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderStorageImageWriteWithoutFormat"
  target="_blank"
  rel="noopener"><code>shaderStorageImageWriteWithoutFormat</code></a>
  is not enabled, any variable created with a “Type” of `OpTypeImage`
  that has a “Sampled” operand of 2 and an “Image Format” operand of
  `Unknown` **must** be decorated with `NonWritable`

- <a href="#VUID-RuntimeSpirv-apiVersion-07955"
  id="VUID-RuntimeSpirv-apiVersion-07955"></a>
  VUID-RuntimeSpirv-apiVersion-07955  
  If
  [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html)::`apiVersion`
  is less than Vulkan 1.3, the
  [VK_KHR_format_feature_flags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_format_feature_flags2.html)
  extension is not supported, and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderStorageImageReadWithoutFormat"
  target="_blank"
  rel="noopener"><code>shaderStorageImageReadWithoutFormat</code></a> is
  not enabled, any variable created with a “Type” of `OpTypeImage` that
  has a “Sampled” operand of 2 and an “Image Format” operand of
  `Unknown` **must** be decorated with `NonReadable`

- <a href="#VUID-RuntimeSpirv-OpImageWrite-07112"
  id="VUID-RuntimeSpirv-OpImageWrite-07112"></a>
  VUID-RuntimeSpirv-OpImageWrite-07112  
  `OpImageWrite` to any `Image` whose `Image` `Format` is not `Unknown`
  **must** have the `Texel` operand contain at least as many components
  as the corresponding [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) as given in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-image-formats"
  target="_blank" rel="noopener">SPIR-V Image Format compatibility
  table</a>

- <a href="#VUID-RuntimeSpirv-Location-06272"
  id="VUID-RuntimeSpirv-Location-06272"></a>
  VUID-RuntimeSpirv-Location-06272  
  The sum of `Location` and the number of locations the variable it
  decorates consumes **must** be less than or equal to the value for the
  matching `Execution` `Model` defined in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-iointerfaces-limits"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-iointerfaces-limits</a>

- <a href="#VUID-RuntimeSpirv-Location-06428"
  id="VUID-RuntimeSpirv-Location-06428"></a>
  VUID-RuntimeSpirv-Location-06428  
  The maximum number of storage buffers, storage images, and output
  `Location` decorated color attachments written to in the `Fragment`
  `Execution` `Model` **must** be less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxFragmentCombinedOutputResources"
  target="_blank"
  rel="noopener"><code>maxFragmentCombinedOutputResources</code></a>

- <a href="#VUID-RuntimeSpirv-NonUniform-06274"
  id="VUID-RuntimeSpirv-NonUniform-06274"></a>
  VUID-RuntimeSpirv-NonUniform-06274  
  If an instruction loads from or stores to a resource (including
  atomics and image instructions) and the resource descriptor being
  accessed is not dynamically uniform, then the operand corresponding to
  that resource (e.g. the pointer or sampled image operand) **must** be
  decorated with `NonUniform`

- <a href="#VUID-RuntimeSpirv-None-06275"
  id="VUID-RuntimeSpirv-None-06275"></a> VUID-RuntimeSpirv-None-06275  
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-subgroup-extended-types"
  target="_blank"
  rel="noopener"><code>shaderSubgroupExtendedTypes</code></a> **must**
  be enabled for <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-group-operations"
  target="_blank" rel="noopener">group operations</a> to use 8-bit
  integer, 16-bit integer, 64-bit integer, 16-bit floating-point, and
  vectors of these types

- <a href="#VUID-RuntimeSpirv-subgroupBroadcastDynamicId-06276"
  id="VUID-RuntimeSpirv-subgroupBroadcastDynamicId-06276"></a>
  VUID-RuntimeSpirv-subgroupBroadcastDynamicId-06276  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-subgroupBroadcastDynamicId"
  target="_blank"
  rel="noopener"><code>subgroupBroadcastDynamicId</code></a> is
  `VK_TRUE`, and the shader module version is 1.5 or higher, the “Index”
  for `OpGroupNonUniformQuadBroadcast` **must** be dynamically uniform
  within the derivative group. Otherwise, “Index” **must** be a constant

- <a href="#VUID-RuntimeSpirv-subgroupBroadcastDynamicId-06277"
  id="VUID-RuntimeSpirv-subgroupBroadcastDynamicId-06277"></a>
  VUID-RuntimeSpirv-subgroupBroadcastDynamicId-06277  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-subgroupBroadcastDynamicId"
  target="_blank"
  rel="noopener"><code>subgroupBroadcastDynamicId</code></a> is
  `VK_TRUE`, and the shader module version is 1.5 or higher, the “Id”
  for `OpGroupNonUniformBroadcast` **must** be dynamically uniform
  within the subgroup. Otherwise, “Id” **must** be a constant

- <a href="#VUID-RuntimeSpirv-None-06278"
  id="VUID-RuntimeSpirv-None-06278"></a> VUID-RuntimeSpirv-None-06278  
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderBufferInt64Atomics"
  target="_blank" rel="noopener"><code>shaderBufferInt64Atomics</code></a>
  **must** be enabled for 64-bit integer atomic operations to be
  supported on a *Pointer* with a `Storage` `Class` of `StorageBuffer`
  or `Uniform`

- <a href="#VUID-RuntimeSpirv-None-06279"
  id="VUID-RuntimeSpirv-None-06279"></a> VUID-RuntimeSpirv-None-06279  
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderSharedInt64Atomics"
  target="_blank" rel="noopener"><code>shaderSharedInt64Atomics</code></a>
  **must** be enabled for 64-bit integer atomic operations to be
  supported on a *Pointer* with a `Storage` `Class` of `Workgroup`

- <a href="#VUID-RuntimeSpirv-None-06284"
  id="VUID-RuntimeSpirv-None-06284"></a> VUID-RuntimeSpirv-None-06284  
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderBufferFloat32Atomics"
  target="_blank"
  rel="noopener"><code>shaderBufferFloat32Atomics</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderBufferFloat32AtomicAdd"
  target="_blank"
  rel="noopener"><code>shaderBufferFloat32AtomicAdd</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderBufferFloat64Atomics"
  target="_blank"
  rel="noopener"><code>shaderBufferFloat64Atomics</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderBufferFloat64AtomicAdd"
  target="_blank"
  rel="noopener"><code>shaderBufferFloat64AtomicAdd</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderBufferFloat16AtomicMinMax"
  target="_blank"
  rel="noopener"><code>shaderBufferFloat16Atomics</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderBufferFloat16AtomicMinMax"
  target="_blank"
  rel="noopener"><code>shaderBufferFloat16AtomicAdd</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderBufferFloat16AtomicMinMax"
  target="_blank"
  rel="noopener"><code>shaderBufferFloat16AtomicMinMax</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderBufferFloat32AtomicMinMax"
  target="_blank"
  rel="noopener"><code>shaderBufferFloat32AtomicMinMax</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderBufferFloat64AtomicMinMax"
  target="_blank"
  rel="noopener"><code>shaderBufferFloat64AtomicMinMax</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderFloat16VectorAtomics"
  target="_blank"
  rel="noopener"><code>shaderFloat16VectorAtomics</code></a> **must** be
  enabled for floating-point atomic operations to be supported on a
  *Pointer* with a `Storage` `Class` of `StorageBuffer`

- <a href="#VUID-RuntimeSpirv-None-06285"
  id="VUID-RuntimeSpirv-None-06285"></a> VUID-RuntimeSpirv-None-06285  
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderSharedFloat32Atomics"
  target="_blank"
  rel="noopener"><code>shaderSharedFloat32Atomics</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderSharedFloat32AtomicAdd"
  target="_blank"
  rel="noopener"><code>shaderSharedFloat32AtomicAdd</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderSharedFloat64Atomics"
  target="_blank"
  rel="noopener"><code>shaderSharedFloat64Atomics</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderSharedFloat64AtomicAdd"
  target="_blank"
  rel="noopener"><code>shaderSharedFloat64AtomicAdd</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderBufferFloat16AtomicMinMax"
  target="_blank"
  rel="noopener"><code>shaderSharedFloat16Atomics</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderBufferFloat16AtomicMinMax"
  target="_blank"
  rel="noopener"><code>shaderSharedFloat16AtomicAdd</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderBufferFloat16AtomicMinMax"
  target="_blank"
  rel="noopener"><code>shaderSharedFloat16AtomicMinMax</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderSharedFloat32AtomicMinMax"
  target="_blank"
  rel="noopener"><code>shaderSharedFloat32AtomicMinMax</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderSharedFloat64AtomicMinMax"
  target="_blank"
  rel="noopener"><code>shaderSharedFloat64AtomicMinMax</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderFloat16VectorAtomics"
  target="_blank"
  rel="noopener"><code>shaderFloat16VectorAtomics</code></a>, **must**
  be enabled for floating-point atomic operations to be supported on a
  *Pointer* with a `Storage` `Class` of `Workgroup`

- <a href="#VUID-RuntimeSpirv-None-06286"
  id="VUID-RuntimeSpirv-None-06286"></a> VUID-RuntimeSpirv-None-06286  
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderImageFloat32Atomics"
  target="_blank"
  rel="noopener"><code>shaderImageFloat32Atomics</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderImageFloat32AtomicAdd"
  target="_blank"
  rel="noopener"><code>shaderImageFloat32AtomicAdd</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderImageFloat32AtomicMinMax"
  target="_blank"
  rel="noopener"><code>shaderImageFloat32AtomicMinMax</code></a>,
  **must** be enabled for 32-bit floating-point atomic operations to be
  supported on a *Pointer* with a `Storage` `Class` of `Image`

- <a href="#VUID-RuntimeSpirv-None-06287"
  id="VUID-RuntimeSpirv-None-06287"></a> VUID-RuntimeSpirv-None-06287  
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-sparseImageFloat32Atomics"
  target="_blank"
  rel="noopener"><code>sparseImageFloat32Atomics</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-sparseImageFloat32AtomicAdd"
  target="_blank"
  rel="noopener"><code>sparseImageFloat32AtomicAdd</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-sparseImageFloat32AtomicMinMax"
  target="_blank"
  rel="noopener"><code>sparseImageFloat32AtomicMinMax</code></a>,
  **must** be enabled for 32-bit floating-point atomics to be supported
  on sparse images

- <a href="#VUID-RuntimeSpirv-None-06288"
  id="VUID-RuntimeSpirv-None-06288"></a> VUID-RuntimeSpirv-None-06288  
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderImageInt64Atomics"
  target="_blank" rel="noopener"><code>shaderImageInt64Atomics</code></a>
  **must** be enabled for 64-bit integer atomic operations to be
  supported on a *Pointer* with a `Storage` `Class` of `Image`

- <a href="#VUID-RuntimeSpirv-denormBehaviorIndependence-06289"
  id="VUID-RuntimeSpirv-denormBehaviorIndependence-06289"></a>
  VUID-RuntimeSpirv-denormBehaviorIndependence-06289  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-denormBehaviorIndependence"
  target="_blank"
  rel="noopener"><code>denormBehaviorIndependence</code></a> is
  `VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_32_BIT_ONLY`, then the entry
  point **must** use the same denormals `Execution` `Mode` for both
  16-bit and 64-bit floating-point types

- <a href="#VUID-RuntimeSpirv-denormBehaviorIndependence-06290"
  id="VUID-RuntimeSpirv-denormBehaviorIndependence-06290"></a>
  VUID-RuntimeSpirv-denormBehaviorIndependence-06290  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-denormBehaviorIndependence"
  target="_blank"
  rel="noopener"><code>denormBehaviorIndependence</code></a> is
  `VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_NONE`, then the entry point
  **must** use the same denormals `Execution` `Mode` for all
  floating-point types

- <a href="#VUID-RuntimeSpirv-roundingModeIndependence-06291"
  id="VUID-RuntimeSpirv-roundingModeIndependence-06291"></a>
  VUID-RuntimeSpirv-roundingModeIndependence-06291  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-roundingModeIndependence"
  target="_blank" rel="noopener"><code>roundingModeIndependence</code></a>
  is `VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_32_BIT_ONLY`, then the entry
  point **must** use the same rounding `Execution` `Mode` for both
  16-bit and 64-bit floating-point types

- <a href="#VUID-RuntimeSpirv-roundingModeIndependence-06292"
  id="VUID-RuntimeSpirv-roundingModeIndependence-06292"></a>
  VUID-RuntimeSpirv-roundingModeIndependence-06292  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-roundingModeIndependence"
  target="_blank" rel="noopener"><code>roundingModeIndependence</code></a>
  is `VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_NONE`, then the entry point
  **must** use the same rounding `Execution` `Mode` for all
  floating-point types

- <a href="#VUID-RuntimeSpirv-shaderSignedZeroInfNanPreserveFloat16-06293"
  id="VUID-RuntimeSpirv-shaderSignedZeroInfNanPreserveFloat16-06293"></a>
  VUID-RuntimeSpirv-shaderSignedZeroInfNanPreserveFloat16-06293  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-shaderSignedZeroInfNanPreserveFloat16"
  target="_blank"
  rel="noopener"><code>shaderSignedZeroInfNanPreserveFloat16</code></a>
  is `VK_FALSE`, then `SignedZeroInfNanPreserve` for 16-bit
  floating-point type **must** not be used

- <a href="#VUID-RuntimeSpirv-shaderSignedZeroInfNanPreserveFloat32-06294"
  id="VUID-RuntimeSpirv-shaderSignedZeroInfNanPreserveFloat32-06294"></a>
  VUID-RuntimeSpirv-shaderSignedZeroInfNanPreserveFloat32-06294  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-shaderSignedZeroInfNanPreserveFloat32"
  target="_blank"
  rel="noopener"><code>shaderSignedZeroInfNanPreserveFloat32</code></a>
  is `VK_FALSE`, then `SignedZeroInfNanPreserve` for 32-bit
  floating-point type **must** not be used

- <a href="#VUID-RuntimeSpirv-shaderSignedZeroInfNanPreserveFloat64-06295"
  id="VUID-RuntimeSpirv-shaderSignedZeroInfNanPreserveFloat64-06295"></a>
  VUID-RuntimeSpirv-shaderSignedZeroInfNanPreserveFloat64-06295  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-shaderSignedZeroInfNanPreserveFloat64"
  target="_blank"
  rel="noopener"><code>shaderSignedZeroInfNanPreserveFloat64</code></a>
  is `VK_FALSE`, then `SignedZeroInfNanPreserve` for 64-bit
  floating-point type **must** not be used

- <a href="#VUID-RuntimeSpirv-shaderDenormPreserveFloat16-06296"
  id="VUID-RuntimeSpirv-shaderDenormPreserveFloat16-06296"></a>
  VUID-RuntimeSpirv-shaderDenormPreserveFloat16-06296  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-shaderDenormPreserveFloat16"
  target="_blank"
  rel="noopener"><code>shaderDenormPreserveFloat16</code></a> is
  `VK_FALSE`, then `DenormPreserve` for 16-bit floating-point type
  **must** not be used

- <a href="#VUID-RuntimeSpirv-shaderDenormPreserveFloat32-06297"
  id="VUID-RuntimeSpirv-shaderDenormPreserveFloat32-06297"></a>
  VUID-RuntimeSpirv-shaderDenormPreserveFloat32-06297  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-shaderDenormPreserveFloat32"
  target="_blank"
  rel="noopener"><code>shaderDenormPreserveFloat32</code></a> is
  `VK_FALSE`, then `DenormPreserve` for 32-bit floating-point type
  **must** not be used

- <a href="#VUID-RuntimeSpirv-shaderDenormPreserveFloat64-06298"
  id="VUID-RuntimeSpirv-shaderDenormPreserveFloat64-06298"></a>
  VUID-RuntimeSpirv-shaderDenormPreserveFloat64-06298  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-shaderDenormPreserveFloat64"
  target="_blank"
  rel="noopener"><code>shaderDenormPreserveFloat64</code></a> is
  `VK_FALSE`, then `DenormPreserve` for 64-bit floating-point type
  **must** not be used

- <a href="#VUID-RuntimeSpirv-shaderDenormFlushToZeroFloat16-06299"
  id="VUID-RuntimeSpirv-shaderDenormFlushToZeroFloat16-06299"></a>
  VUID-RuntimeSpirv-shaderDenormFlushToZeroFloat16-06299  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-shaderDenormFlushToZeroFloat16"
  target="_blank"
  rel="noopener"><code>shaderDenormFlushToZeroFloat16</code></a> is
  `VK_FALSE`, then `DenormFlushToZero` for 16-bit floating-point type
  **must** not be used

- <a href="#VUID-RuntimeSpirv-shaderDenormFlushToZeroFloat32-06300"
  id="VUID-RuntimeSpirv-shaderDenormFlushToZeroFloat32-06300"></a>
  VUID-RuntimeSpirv-shaderDenormFlushToZeroFloat32-06300  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-shaderDenormFlushToZeroFloat32"
  target="_blank"
  rel="noopener"><code>shaderDenormFlushToZeroFloat32</code></a> is
  `VK_FALSE`, then `DenormFlushToZero` for 32-bit floating-point type
  **must** not be used

- <a href="#VUID-RuntimeSpirv-shaderDenormFlushToZeroFloat64-06301"
  id="VUID-RuntimeSpirv-shaderDenormFlushToZeroFloat64-06301"></a>
  VUID-RuntimeSpirv-shaderDenormFlushToZeroFloat64-06301  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-shaderDenormFlushToZeroFloat64"
  target="_blank"
  rel="noopener"><code>shaderDenormFlushToZeroFloat64</code></a> is
  `VK_FALSE`, then `DenormFlushToZero` for 64-bit floating-point type
  **must** not be used

- <a href="#VUID-RuntimeSpirv-shaderRoundingModeRTEFloat16-06302"
  id="VUID-RuntimeSpirv-shaderRoundingModeRTEFloat16-06302"></a>
  VUID-RuntimeSpirv-shaderRoundingModeRTEFloat16-06302  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-shaderRoundingModeRTEFloat16"
  target="_blank"
  rel="noopener"><code>shaderRoundingModeRTEFloat16</code></a> is
  `VK_FALSE`, then `RoundingModeRTE` for 16-bit floating-point type
  **must** not be used

- <a href="#VUID-RuntimeSpirv-shaderRoundingModeRTEFloat32-06303"
  id="VUID-RuntimeSpirv-shaderRoundingModeRTEFloat32-06303"></a>
  VUID-RuntimeSpirv-shaderRoundingModeRTEFloat32-06303  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-shaderRoundingModeRTEFloat32"
  target="_blank"
  rel="noopener"><code>shaderRoundingModeRTEFloat32</code></a> is
  `VK_FALSE`, then `RoundingModeRTE` for 32-bit floating-point type
  **must** not be used

- <a href="#VUID-RuntimeSpirv-shaderRoundingModeRTEFloat64-06304"
  id="VUID-RuntimeSpirv-shaderRoundingModeRTEFloat64-06304"></a>
  VUID-RuntimeSpirv-shaderRoundingModeRTEFloat64-06304  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-shaderRoundingModeRTEFloat64"
  target="_blank"
  rel="noopener"><code>shaderRoundingModeRTEFloat64</code></a> is
  `VK_FALSE`, then `RoundingModeRTE` for 64-bit floating-point type
  **must** not be used

- <a href="#VUID-RuntimeSpirv-shaderRoundingModeRTZFloat16-06305"
  id="VUID-RuntimeSpirv-shaderRoundingModeRTZFloat16-06305"></a>
  VUID-RuntimeSpirv-shaderRoundingModeRTZFloat16-06305  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-shaderRoundingModeRTZFloat16"
  target="_blank"
  rel="noopener"><code>shaderRoundingModeRTZFloat16</code></a> is
  `VK_FALSE`, then `RoundingModeRTZ` for 16-bit floating-point type
  **must** not be used

- <a href="#VUID-RuntimeSpirv-shaderRoundingModeRTZFloat32-06306"
  id="VUID-RuntimeSpirv-shaderRoundingModeRTZFloat32-06306"></a>
  VUID-RuntimeSpirv-shaderRoundingModeRTZFloat32-06306  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-shaderRoundingModeRTZFloat32"
  target="_blank"
  rel="noopener"><code>shaderRoundingModeRTZFloat32</code></a> is
  `VK_FALSE`, then `RoundingModeRTZ` for 32-bit floating-point type
  **must** not be used

- <a href="#VUID-RuntimeSpirv-shaderRoundingModeRTZFloat64-06307"
  id="VUID-RuntimeSpirv-shaderRoundingModeRTZFloat64-06307"></a>
  VUID-RuntimeSpirv-shaderRoundingModeRTZFloat64-06307  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-shaderRoundingModeRTZFloat64"
  target="_blank"
  rel="noopener"><code>shaderRoundingModeRTZFloat64</code></a> is
  `VK_FALSE`, then `RoundingModeRTZ` for 64-bit floating-point type
  **must** not be used

- <a href="#VUID-RuntimeSpirv-shaderSignedZeroInfNanPreserveFloat16-09559"
  id="VUID-RuntimeSpirv-shaderSignedZeroInfNanPreserveFloat16-09559"></a>
  VUID-RuntimeSpirv-shaderSignedZeroInfNanPreserveFloat16-09559  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-shaderSignedZeroInfNanPreserveFloat16"
  target="_blank"
  rel="noopener"><code>shaderSignedZeroInfNanPreserveFloat16</code></a>
  is `VK_FALSE` then any `FPFastMathDefault` execution mode with a type
  of 16-bit float **must** include the `NSZ`, `NotInf`, and `NotNaN`
  flags

- <a href="#VUID-RuntimeSpirv-shaderSignedZeroInfNanPreserveFloat16-09560"
  id="VUID-RuntimeSpirv-shaderSignedZeroInfNanPreserveFloat16-09560"></a>
  VUID-RuntimeSpirv-shaderSignedZeroInfNanPreserveFloat16-09560  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-shaderSignedZeroInfNanPreserveFloat16"
  target="_blank"
  rel="noopener"><code>shaderSignedZeroInfNanPreserveFloat16</code></a>
  is `VK_FALSE` then any `FPFastMathMode` decoration on an instruction
  with result type or any operand type that includes a 16-bit float
  **must** include the `NSZ`, `NotInf`, and `NotNaN` flags

- <a href="#VUID-RuntimeSpirv-shaderSignedZeroInfNanPreserveFloat32-09561"
  id="VUID-RuntimeSpirv-shaderSignedZeroInfNanPreserveFloat32-09561"></a>
  VUID-RuntimeSpirv-shaderSignedZeroInfNanPreserveFloat32-09561  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-shaderSignedZeroInfNanPreserveFloat32"
  target="_blank"
  rel="noopener"><code>shaderSignedZeroInfNanPreserveFloat32</code></a>
  is `VK_FALSE` then any `FPFastMathDefault` execution mode with a type
  of 32-bit float **must** include the `NSZ`, `NotInf`, and `NotNaN`
  flags

- <a href="#VUID-RuntimeSpirv-shaderSignedZeroInfNanPreserveFloat32-09562"
  id="VUID-RuntimeSpirv-shaderSignedZeroInfNanPreserveFloat32-09562"></a>
  VUID-RuntimeSpirv-shaderSignedZeroInfNanPreserveFloat32-09562  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-shaderSignedZeroInfNanPreserveFloat32"
  target="_blank"
  rel="noopener"><code>shaderSignedZeroInfNanPreserveFloat32</code></a>
  is `VK_FALSE` then any `FPFastMathMode` decoration on an instruction
  with result type or any operand type that includes a 32-bit float
  **must** include the `NSZ`, `NotInf`, and `NotNaN` flags

- <a href="#VUID-RuntimeSpirv-shaderSignedZeroInfNanPreserveFloat64-09563"
  id="VUID-RuntimeSpirv-shaderSignedZeroInfNanPreserveFloat64-09563"></a>
  VUID-RuntimeSpirv-shaderSignedZeroInfNanPreserveFloat64-09563  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-shaderSignedZeroInfNanPreserveFloat64"
  target="_blank"
  rel="noopener"><code>shaderSignedZeroInfNanPreserveFloat64</code></a>
  is `VK_FALSE` then any `FPFastMathDefault` execution mode with a type
  of 64-bit float **must** include the `NSZ`, `NotInf`, and `NotNaN`
  flags

- <a href="#VUID-RuntimeSpirv-shaderSignedZeroInfNanPreserveFloat64-09564"
  id="VUID-RuntimeSpirv-shaderSignedZeroInfNanPreserveFloat64-09564"></a>
  VUID-RuntimeSpirv-shaderSignedZeroInfNanPreserveFloat64-09564  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-shaderSignedZeroInfNanPreserveFloat64"
  target="_blank"
  rel="noopener"><code>shaderSignedZeroInfNanPreserveFloat64</code></a>
  is `VK_FALSE` then any `FPFastMathMode` decoration on an instruction
  with result type or any operand type that includes a 64-bit float
  **must** include the `NSZ`, `NotInf`, and `NotNaN` flags

- <a href="#VUID-RuntimeSpirv-Offset-06308"
  id="VUID-RuntimeSpirv-Offset-06308"></a>
  VUID-RuntimeSpirv-Offset-06308  
  The `Offset` plus size of the type of each variable, in the output
  interface of the entry point being compiled, decorated with
  `XfbBuffer` **must** not be greater than
  [VkPhysicalDeviceTransformFeedbackPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTransformFeedbackPropertiesEXT.html)::`maxTransformFeedbackBufferDataSize`

- <a href="#VUID-RuntimeSpirv-XfbBuffer-06309"
  id="VUID-RuntimeSpirv-XfbBuffer-06309"></a>
  VUID-RuntimeSpirv-XfbBuffer-06309  
  For any given `XfbBuffer` value, define the buffer data size to be
  smallest number of bytes such that, for all outputs decorated with the
  same `XfbBuffer` value, the size of the output interface variable plus
  the `Offset` is less than or equal to the buffer data size. For a
  given `Stream`, the sum of all the buffer data sizes for all buffers
  writing to that stream the **must** not exceed
  [VkPhysicalDeviceTransformFeedbackPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTransformFeedbackPropertiesEXT.html)::`maxTransformFeedbackStreamDataSize`

- <a href="#VUID-RuntimeSpirv-OpEmitStreamVertex-06310"
  id="VUID-RuntimeSpirv-OpEmitStreamVertex-06310"></a>
  VUID-RuntimeSpirv-OpEmitStreamVertex-06310  
  The Stream value to `OpEmitStreamVertex` and `OpEndStreamPrimitive`
  **must** be less than
  [VkPhysicalDeviceTransformFeedbackPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTransformFeedbackPropertiesEXT.html)::`maxTransformFeedbackStreams`

- <a
  href="#VUID-RuntimeSpirv-transformFeedbackStreamsLinesTriangles-06311"
  id="VUID-RuntimeSpirv-transformFeedbackStreamsLinesTriangles-06311"></a>
  VUID-RuntimeSpirv-transformFeedbackStreamsLinesTriangles-06311  
  If the geometry shader emits to more than one vertex stream and
  [VkPhysicalDeviceTransformFeedbackPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTransformFeedbackPropertiesEXT.html)::`transformFeedbackStreamsLinesTriangles`
  is `VK_FALSE`, then `Execution` `Mode` **must** be `OutputPoints`

- <a href="#VUID-RuntimeSpirv-Stream-06312"
  id="VUID-RuntimeSpirv-Stream-06312"></a>
  VUID-RuntimeSpirv-Stream-06312  
  The stream number value to `Stream` **must** be less than
  [VkPhysicalDeviceTransformFeedbackPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTransformFeedbackPropertiesEXT.html)::`maxTransformFeedbackStreams`

- <a href="#VUID-RuntimeSpirv-XfbStride-06313"
  id="VUID-RuntimeSpirv-XfbStride-06313"></a>
  VUID-RuntimeSpirv-XfbStride-06313  
  The XFB Stride value to `XfbStride` **must** be less than or equal to
  [VkPhysicalDeviceTransformFeedbackPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTransformFeedbackPropertiesEXT.html)::`maxTransformFeedbackBufferDataStride`

- <a href="#VUID-RuntimeSpirv-PhysicalStorageBuffer64-06314"
  id="VUID-RuntimeSpirv-PhysicalStorageBuffer64-06314"></a>
  VUID-RuntimeSpirv-PhysicalStorageBuffer64-06314  
  If the `PhysicalStorageBuffer64` addressing model is enabled any load
  or store through a physical pointer type **must** be aligned to a
  multiple of the size of the largest scalar type in the pointed-to type

- <a href="#VUID-RuntimeSpirv-PhysicalStorageBuffer64-06315"
  id="VUID-RuntimeSpirv-PhysicalStorageBuffer64-06315"></a>
  VUID-RuntimeSpirv-PhysicalStorageBuffer64-06315  
  If the `PhysicalStorageBuffer64` addressing model is enabled the
  pointer value of a memory access instruction **must** be at least as
  aligned as specified by the `Aligned` memory access operand

- <a href="#VUID-RuntimeSpirv-OpTypeCooperativeMatrixNV-06316"
  id="VUID-RuntimeSpirv-OpTypeCooperativeMatrixNV-06316"></a>
  VUID-RuntimeSpirv-OpTypeCooperativeMatrixNV-06316  
  For `OpTypeCooperativeMatrixNV`, the component type, scope, number of
  rows, and number of columns **must** match one of the matrices in any
  of the supported
  [VkCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesNV.html)

- <a href="#VUID-RuntimeSpirv-OpTypeCooperativeMatrixMulAddNV-10059"
  id="VUID-RuntimeSpirv-OpTypeCooperativeMatrixMulAddNV-10059"></a>
  VUID-RuntimeSpirv-OpTypeCooperativeMatrixMulAddNV-10059  
  For `OpTypeCooperativeMatrixMulAddNV`, the operands **must** match a
  supported
  [VkCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesNV.html),
  such that:

  - The type of `A` **must** have `Rows` match
    [VkCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesNV.html)::`MSize`,
    `Columns` match
    [VkCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesNV.html)::`KSize`,
    and `ComponentType` match
    [VkCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesNV.html)::`AType`.

  - The type of `B` **must** have `Rows` match
    [VkCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesNV.html)::`KSize`,
    `Columns` match
    [VkCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesNV.html)::`NSize`,
    and `ComponentType` match
    [VkCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesNV.html)::`BType`.

  - The type of `C` **must** have `Rows` match
    [VkCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesNV.html)::`MSize`,
    `Columns` match
    [VkCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesNV.html)::`NSize`,
    and `ComponentType` match
    [VkCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesNV.html)::`CType`.

  - The type of `Result` **must** have `Rows` match
    [VkCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesNV.html)::`MSize`,
    `Columns` match
    [VkCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesNV.html)::`NSize`,
    and `ComponentType` match
    [VkCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesNV.html)::`DType`.

  - The scope of all cooperative matrix operands **must** be
    [VkScopeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkScopeNV.html)::`VK_SCOPE_SUBGROUP_NV`.

  - If `ComponentType` of `A`, `B`, `C`, or `Result` is a signed
    integral type, the `Signedness` operand of the `OpTypeInt` must be
    1.

  - If `ComponentType` of `A`, `B`, `C`, or `Result` is an unsigned
    integral type, the `Signedness` operand of the `OpTypeInt` must be
    0.

- <a href="#VUID-RuntimeSpirv-OpTypeCooperativeMatrixNV-06322"
  id="VUID-RuntimeSpirv-OpTypeCooperativeMatrixNV-06322"></a>
  VUID-RuntimeSpirv-OpTypeCooperativeMatrixNV-06322  
  `OpTypeCooperativeMatrixNV` and `OpCooperativeMatrix*` instructions
  **must** not be used in shader stages not included in
  [VkPhysicalDeviceCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCooperativeMatrixPropertiesNV.html)::`cooperativeMatrixSupportedStages`

- <a href="#VUID-RuntimeSpirv-OpTypeCooperativeMatrixKHR-08974"
  id="VUID-RuntimeSpirv-OpTypeCooperativeMatrixKHR-08974"></a>
  VUID-RuntimeSpirv-OpTypeCooperativeMatrixKHR-08974  
  For `OpTypeCooperativeMatrixKHR`, the component type, scope, number of
  rows, and number of columns **must** match one of the matrices in any
  of the supported
  [VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesKHR.html).

- <a href="#VUID-RuntimeSpirv-OpCooperativeMatrixMulAddKHR-10060"
  id="VUID-RuntimeSpirv-OpCooperativeMatrixMulAddKHR-10060"></a>
  VUID-RuntimeSpirv-OpCooperativeMatrixMulAddKHR-10060  
  For `OpCooperativeMatrixMulAddKHR`, the operands **must** match a
  supported
  [VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesKHR.html),
  such that:

  - The type of `A` **must** have `Rows` match
    [VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesKHR.html)::`MSize`,
    `Columns` match
    [VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesKHR.html)::`KSize`,
    `Use` be `MatrixAKHR`, and `ComponentType` match
    [VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesKHR.html)::`AType`.

  - The type of `B` **must** have `Rows` match
    [VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesKHR.html)::`KSize`,
    `Columns` match
    [VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesKHR.html)::`NSize`,
    `Use` be `MatrixBKHR`, and `ComponentType` match
    [VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesKHR.html)::`BType`.

  - The type of `C` **must** have `Rows` match
    [VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesKHR.html)::`MSize`,
    `Columns` match
    [VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesKHR.html)::`NSize`,
    `Use` be `MatrixAccumulatorKHR`, and `ComponentType` match
    [VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesKHR.html)::`CType`.

  - The type of `Result` **must** have `Rows` match
    [VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesKHR.html)::`MSize`,
    `Columns` match
    [VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesKHR.html)::`NSize`,
    `Use` be `MatrixAccumulatorKHR`, and `ComponentType` match
    [VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesKHR.html)::`ResultType`.

  - If and only if
    [VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesKHR.html)::`AType`
    is a signed integer type, `MatrixASignedComponents` **must** be
    used.

  - If and only if
    [VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesKHR.html)::`BType`
    is a signed integer type, `MatrixBSignedComponents` **must** be
    used.

  - If and only if
    [VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesKHR.html)::`CType`
    is a signed integer type, `MatrixCSignedComponents` **must** be
    used.

  - If and only if
    [VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesKHR.html)::`ResultType`
    is a signed integer type, `MatrixResultSignedComponents` **must** be
    used.

  - If and only if
    [VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesKHR.html)::`saturatingAccumulation`
    is `VK_TRUE`, `SaturatingAccumulationKHR` **must** be used.

  - If and only if
    [VkCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesKHR.html)::`saturatingAccumulation`
    is `VK_FALSE`, `SaturatingAccumulationKHR` **must** not be used.

  - The scope of all cooperative matrix operands **must** be
    [VkScopeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkScopeKHR.html)::`VK_SCOPE_SUBGROUP_KHR`.

- <a href="#VUID-RuntimeSpirv-cooperativeMatrixSupportedStages-08985"
  id="VUID-RuntimeSpirv-cooperativeMatrixSupportedStages-08985"></a>
  VUID-RuntimeSpirv-cooperativeMatrixSupportedStages-08985  
  `OpTypeCooperativeMatrixKHR` and `OpCooperativeMatrix*` instructions
  **must** not be used in shader stages not included in
  [VkPhysicalDeviceCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCooperativeMatrixPropertiesKHR.html)::`cooperativeMatrixSupportedStages`

- <a href="#VUID-RuntimeSpirv-DescriptorSet-06323"
  id="VUID-RuntimeSpirv-DescriptorSet-06323"></a>
  VUID-RuntimeSpirv-DescriptorSet-06323  
  `DescriptorSet` and `Binding` decorations **must** obey the
  constraints on `Storage` `Class`, type, and descriptor type described
  in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-resources-setandbinding"
  target="_blank" rel="noopener">DescriptorSet and Binding Assignment</a>

- <a href="#VUID-RuntimeSpirv-OpCooperativeMatrixLoadNV-06324"
  id="VUID-RuntimeSpirv-OpCooperativeMatrixLoadNV-06324"></a>
  VUID-RuntimeSpirv-OpCooperativeMatrixLoadNV-06324  
  For `OpCooperativeMatrixLoadNV` and `OpCooperativeMatrixStoreNV`
  instructions, the `Pointer` and `Stride` operands **must** be aligned
  to at least the lesser of 16 bytes or the natural alignment of a row
  or column (depending on `ColumnMajor`) of the matrix (where the
  natural alignment is the number of columns/rows multiplied by the
  component size)

- <a href="#VUID-RuntimeSpirv-MeshNV-07113"
  id="VUID-RuntimeSpirv-MeshNV-07113"></a>
  VUID-RuntimeSpirv-MeshNV-07113  
  For mesh shaders using the `MeshNV` `Execution` `Model` the
  `OutputVertices` `OpExecutionMode` **must** be less than or equal to
  [VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`maxMeshOutputVertices`

- <a href="#VUID-RuntimeSpirv-MeshNV-07114"
  id="VUID-RuntimeSpirv-MeshNV-07114"></a>
  VUID-RuntimeSpirv-MeshNV-07114  
  For mesh shaders using the `MeshNV` `Execution` `Model` the
  `OutputPrimitivesNV` `OpExecutionMode` **must** be less than or equal
  to
  [VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html)::`maxMeshOutputPrimitives`

- <a href="#VUID-RuntimeSpirv-MeshEXT-07115"
  id="VUID-RuntimeSpirv-MeshEXT-07115"></a>
  VUID-RuntimeSpirv-MeshEXT-07115  
  For mesh shaders using the `MeshEXT` `Execution` `Model` the
  `OutputVertices` `OpExecutionMode` **must** be less than or equal to
  [VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshOutputVertices`

- <a href="#VUID-RuntimeSpirv-MeshEXT-07332"
  id="VUID-RuntimeSpirv-MeshEXT-07332"></a>
  VUID-RuntimeSpirv-MeshEXT-07332  
  For mesh shaders using the `MeshEXT` `Execution` `Model` the “Vertex
  Count” operand of `OpSetMeshOutputsEXT` **must** be less than or equal
  to `OutputVertices` `OpExecutionMode`

- <a href="#VUID-RuntimeSpirv-MeshEXT-07116"
  id="VUID-RuntimeSpirv-MeshEXT-07116"></a>
  VUID-RuntimeSpirv-MeshEXT-07116  
  For mesh shaders using the `MeshEXT` `Execution` `Model` the
  `OutputPrimitivesEXT` `OpExecutionMode` **must** be less than or equal
  to
  [VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshOutputPrimitives`

- <a href="#VUID-RuntimeSpirv-MeshEXT-07333"
  id="VUID-RuntimeSpirv-MeshEXT-07333"></a>
  VUID-RuntimeSpirv-MeshEXT-07333  
  For mesh shaders using the `MeshEXT` `Execution` `Model` the
  “Primitive Count” operand of `OpSetMeshOutputsEXT` **must** be less
  than or equal to `OutputPrimitivesEXT` `OpExecutionMode`

- <a href="#VUID-RuntimeSpirv-TaskEXT-07117"
  id="VUID-RuntimeSpirv-TaskEXT-07117"></a>
  VUID-RuntimeSpirv-TaskEXT-07117  
  In task shaders using the `TaskEXT` `Execution` `Model`
  `OpEmitMeshTasksEXT` **must** be called exactly once under dynamically
  uniform conditions

- <a href="#VUID-RuntimeSpirv-MeshEXT-07118"
  id="VUID-RuntimeSpirv-MeshEXT-07118"></a>
  VUID-RuntimeSpirv-MeshEXT-07118  
  In mesh shaders using the `MeshEXT` `Execution` `Model`
  `OpSetMeshOutputsEXT` **must** be called at most once under
  dynamically uniform conditions

- <a href="#VUID-RuntimeSpirv-TaskEXT-07291"
  id="VUID-RuntimeSpirv-TaskEXT-07291"></a>
  VUID-RuntimeSpirv-TaskEXT-07291  
  In task shaders using the `TaskEXT` `Execution` `Model` the `x` size
  in `LocalSize` or `LocalSizeId` **must** be less than or equal to
  [VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxTaskWorkGroupSize`\[0\]

- <a href="#VUID-RuntimeSpirv-TaskEXT-07292"
  id="VUID-RuntimeSpirv-TaskEXT-07292"></a>
  VUID-RuntimeSpirv-TaskEXT-07292  
  In task shaders using the `TaskEXT` `Execution` `Model` the `y` size
  in `LocalSize` or `LocalSizeId` **must** be less than or equal to
  [VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxTaskWorkGroupSize`\[1\]

- <a href="#VUID-RuntimeSpirv-TaskEXT-07293"
  id="VUID-RuntimeSpirv-TaskEXT-07293"></a>
  VUID-RuntimeSpirv-TaskEXT-07293  
  In task shaders using the `TaskEXT` `Execution` `Model` the `z` size
  in `LocalSize` or `LocalSizeId` **must** be less than or equal to
  [VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxTaskWorkGroupSize`\[2\]

- <a href="#VUID-RuntimeSpirv-TaskEXT-07294"
  id="VUID-RuntimeSpirv-TaskEXT-07294"></a>
  VUID-RuntimeSpirv-TaskEXT-07294  
  In task shaders using the `TaskEXT` `Execution` `Model` the product of
  `x` size, `y` size, and `z` size in `LocalSize` or `LocalSizeId`
  **must** be less than or equal to
  [VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxTaskWorkGroupInvocations`

- <a href="#VUID-RuntimeSpirv-MeshEXT-07295"
  id="VUID-RuntimeSpirv-MeshEXT-07295"></a>
  VUID-RuntimeSpirv-MeshEXT-07295  
  For mesh shaders using the `MeshEXT` `Execution` `Model` the `x` size
  in `LocalSize` or `LocalSizeId` **must** be less than or equal to
  [VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshWorkGroupSize`\[0\]

- <a href="#VUID-RuntimeSpirv-MeshEXT-07296"
  id="VUID-RuntimeSpirv-MeshEXT-07296"></a>
  VUID-RuntimeSpirv-MeshEXT-07296  
  For mesh shaders using the `MeshEXT` `Execution` `Model` the `y` size
  in `LocalSize` or `LocalSizeId` **must** be less than or equal to
  [VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshWorkGroupSize`\[1\]

- <a href="#VUID-RuntimeSpirv-MeshEXT-07297"
  id="VUID-RuntimeSpirv-MeshEXT-07297"></a>
  VUID-RuntimeSpirv-MeshEXT-07297  
  For mesh shaders using the `MeshEXT` `Execution` `Model` the `z` size
  in `LocalSize` or `LocalSizeId` **must** be less than or equal to
  [VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshWorkGroupSize`\[2\]

- <a href="#VUID-RuntimeSpirv-MeshEXT-07298"
  id="VUID-RuntimeSpirv-MeshEXT-07298"></a>
  VUID-RuntimeSpirv-MeshEXT-07298  
  For mesh shaders using the `MeshEXT` `Execution` `Model` the product
  of `x` size, `y` size, and `z` size in `LocalSize` or `LocalSizeId`
  **must** be less than or equal to
  [VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshWorkGroupInvocations`

- <a href="#VUID-RuntimeSpirv-TaskEXT-07299"
  id="VUID-RuntimeSpirv-TaskEXT-07299"></a>
  VUID-RuntimeSpirv-TaskEXT-07299  
  In task shaders using the `TaskEXT` `Execution` `Model` the value of
  the “Group Count X” operand of `OpEmitMeshTasksEXT` **must** be less
  than or equal to
  [VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshWorkGroupCount`\[0\]

- <a href="#VUID-RuntimeSpirv-TaskEXT-07300"
  id="VUID-RuntimeSpirv-TaskEXT-07300"></a>
  VUID-RuntimeSpirv-TaskEXT-07300  
  In task shaders using the `TaskEXT` `Execution` `Model` the value of
  the “Group Count Y” operand of `OpEmitMeshTasksEXT` **must** be less
  than or equal to
  [VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshWorkGroupCount`\[1\]

- <a href="#VUID-RuntimeSpirv-TaskEXT-07301"
  id="VUID-RuntimeSpirv-TaskEXT-07301"></a>
  VUID-RuntimeSpirv-TaskEXT-07301  
  In task shaders using the `TaskEXT` `Execution` `Model` the value of
  the “Group Count Z” operand of `OpEmitMeshTasksEXT` **must** be less
  than or equal to
  [VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshWorkGroupCount`\[2\]

- <a href="#VUID-RuntimeSpirv-TaskEXT-07302"
  id="VUID-RuntimeSpirv-TaskEXT-07302"></a>
  VUID-RuntimeSpirv-TaskEXT-07302  
  In task shaders using the `TaskEXT` `Execution` `Model` the product of
  the “Group Count” operands of `OpEmitMeshTasksEXT` **must** be less
  than or equal to
  [VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html)::`maxMeshWorkGroupTotalCount`

- <a href="#VUID-RuntimeSpirv-maxMeshSharedMemorySize-08754"
  id="VUID-RuntimeSpirv-maxMeshSharedMemorySize-08754"></a>
  VUID-RuntimeSpirv-maxMeshSharedMemorySize-08754  
  The sum of size in bytes for variables and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#workgroup-padding"
  target="_blank" rel="noopener">padding</a> in the `Workgroup`
  `Storage` `Class` in the `MeshEXT` `Execution` `Model` **must** be
  less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxMeshSharedMemorySize"
  target="_blank" rel="noopener"><code>maxMeshSharedMemorySize</code></a>

- <a href="#VUID-RuntimeSpirv-maxMeshPayloadAndSharedMemorySize-08755"
  id="VUID-RuntimeSpirv-maxMeshPayloadAndSharedMemorySize-08755"></a>
  VUID-RuntimeSpirv-maxMeshPayloadAndSharedMemorySize-08755  
  The sum of size in bytes for variables and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#workgroup-padding"
  target="_blank" rel="noopener">padding</a> in the
  `TaskPayloadWorkgroupEXT` or `Workgroup` `Storage` `Class` in the
  `MeshEXT` `Execution` `Model` **must** be less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxMeshPayloadAndSharedMemorySize"
  target="_blank"
  rel="noopener"><code>maxMeshPayloadAndSharedMemorySize</code></a>

- <a href="#VUID-RuntimeSpirv-maxMeshOutputMemorySize-08756"
  id="VUID-RuntimeSpirv-maxMeshOutputMemorySize-08756"></a>
  VUID-RuntimeSpirv-maxMeshOutputMemorySize-08756  
  The sum of size in bytes for variables in the `Output` `Storage`
  `Class` in the `MeshEXT` `Execution` `Model` **must** be less than or
  equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxMeshOutputMemorySize"
  target="_blank" rel="noopener"><code>maxMeshOutputMemorySize</code></a>
  according to the formula in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#mesh-output"
  target="_blank" rel="noopener">Mesh Shader Output</a>

- <a href="#VUID-RuntimeSpirv-maxMeshPayloadAndOutputMemorySize-08757"
  id="VUID-RuntimeSpirv-maxMeshPayloadAndOutputMemorySize-08757"></a>
  VUID-RuntimeSpirv-maxMeshPayloadAndOutputMemorySize-08757  
  The sum of size in bytes for variables and in the
  `TaskPayloadWorkgroupEXT` or `Output` `Storage` `Class` in the
  `MeshEXT` `Execution` `Model` **must** be less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxMeshPayloadAndOutputMemorySize"
  target="_blank"
  rel="noopener"><code>maxMeshPayloadAndOutputMemorySize</code></a>
  according to the formula in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#mesh-output"
  target="_blank" rel="noopener">Mesh Shader Output</a>

- <a href="#VUID-RuntimeSpirv-maxTaskPayloadSize-08758"
  id="VUID-RuntimeSpirv-maxTaskPayloadSize-08758"></a>
  VUID-RuntimeSpirv-maxTaskPayloadSize-08758  
  The sum of size in bytes for variables and in the
  `TaskPayloadWorkgroupEXT` `Storage` `Class` in the `TaskEXT`
  `Execution` `Model` **must** be less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxTaskPayloadSize"
  target="_blank" rel="noopener"><code>maxTaskPayloadSize</code></a>

- <a href="#VUID-RuntimeSpirv-maxTaskSharedMemorySize-08759"
  id="VUID-RuntimeSpirv-maxTaskSharedMemorySize-08759"></a>
  VUID-RuntimeSpirv-maxTaskSharedMemorySize-08759  
  The sum of size in bytes for variables and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#workgroup-padding"
  target="_blank" rel="noopener">padding</a> in the `Workgroup`
  `Storage` `Class` in the `TaskEXT` `Execution` `Model` **must** be
  less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxTaskSharedMemorySize"
  target="_blank" rel="noopener"><code>maxTaskSharedMemorySize</code></a>

- <a href="#VUID-RuntimeSpirv-maxTaskPayloadAndSharedMemorySize-08760"
  id="VUID-RuntimeSpirv-maxTaskPayloadAndSharedMemorySize-08760"></a>
  VUID-RuntimeSpirv-maxTaskPayloadAndSharedMemorySize-08760  
  The sum of size in bytes for variables and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#workgroup-padding"
  target="_blank" rel="noopener">padding</a> in the
  `TaskPayloadWorkgroupEXT` or `Workgroup` `Storage` `Class` in the
  `TaskEXT` `Execution` `Model` **must** be less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxTaskPayloadAndSharedMemorySize"
  target="_blank"
  rel="noopener"><code>maxTaskPayloadAndSharedMemorySize</code></a>

- <a href="#VUID-RuntimeSpirv-OpCooperativeMatrixLoadKHR-08986"
  id="VUID-RuntimeSpirv-OpCooperativeMatrixLoadKHR-08986"></a>
  VUID-RuntimeSpirv-OpCooperativeMatrixLoadKHR-08986  
  For `OpCooperativeMatrixLoadKHR` and `OpCooperativeMatrixStoreKHR`
  instructions, the `Pointer` and `Stride` operands **must** be aligned
  to at least the lesser of 16 bytes or the natural alignment of a row
  or column (depending on `ColumnMajor`) of the matrix (where the
  natural alignment is the number of columns/rows multiplied by the
  component size)

- <a
  href="#VUID-RuntimeSpirv-shaderSampleRateInterpolationFunctions-06325"
  id="VUID-RuntimeSpirv-shaderSampleRateInterpolationFunctions-06325"></a>
  VUID-RuntimeSpirv-shaderSampleRateInterpolationFunctions-06325  
  If the [`VK_KHR_portability_subset`](VK_KHR_portability_subset.html)
  extension is enabled, and
  [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`shaderSampleRateInterpolationFunctions`
  is `VK_FALSE`, then `GLSL.std.450` fragment interpolation functions
  are not supported by the implementation and `OpCapability` **must**
  not be set to `InterpolationFunction`

- <a href="#VUID-RuntimeSpirv-tessellationShader-06326"
  id="VUID-RuntimeSpirv-tessellationShader-06326"></a>
  VUID-RuntimeSpirv-tessellationShader-06326  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-tessellationShader"
  target="_blank" rel="noopener"><code>tessellationShader</code></a> is
  enabled, and the
  [`VK_KHR_portability_subset`](VK_KHR_portability_subset.html)
  extension is enabled, and
  [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`tessellationIsolines`
  is `VK_FALSE`, then `OpExecutionMode` **must** not be set to
  `IsoLines`

- <a href="#VUID-RuntimeSpirv-tessellationShader-06327"
  id="VUID-RuntimeSpirv-tessellationShader-06327"></a>
  VUID-RuntimeSpirv-tessellationShader-06327  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-tessellationShader"
  target="_blank" rel="noopener"><code>tessellationShader</code></a> is
  enabled, and the
  [`VK_KHR_portability_subset`](VK_KHR_portability_subset.html)
  extension is enabled, and
  [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`tessellationPointMode`
  is `VK_FALSE`, then `OpExecutionMode` **must** not be set to
  `PointMode`

- <a href="#VUID-RuntimeSpirv-storageBuffer8BitAccess-06328"
  id="VUID-RuntimeSpirv-storageBuffer8BitAccess-06328"></a>
  VUID-RuntimeSpirv-storageBuffer8BitAccess-06328  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-storageBuffer8BitAccess"
  target="_blank" rel="noopener"><code>storageBuffer8BitAccess</code></a>
  is `VK_FALSE`, then objects containing an 8-bit integer element
  **must** not have `Storage` `Class` of `StorageBuffer`,
  `ShaderRecordBufferKHR`, or `PhysicalStorageBuffer`

- <a href="#VUID-RuntimeSpirv-uniformAndStorageBuffer8BitAccess-06329"
  id="VUID-RuntimeSpirv-uniformAndStorageBuffer8BitAccess-06329"></a>
  VUID-RuntimeSpirv-uniformAndStorageBuffer8BitAccess-06329  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-uniformAndStorageBuffer8BitAccess"
  target="_blank"
  rel="noopener"><code>uniformAndStorageBuffer8BitAccess</code></a> is
  `VK_FALSE`, then objects in the `Uniform` `Storage` `Class` with the
  `Block` decoration **must** not have an 8-bit integer member

- <a href="#VUID-RuntimeSpirv-storagePushConstant8-06330"
  id="VUID-RuntimeSpirv-storagePushConstant8-06330"></a>
  VUID-RuntimeSpirv-storagePushConstant8-06330  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-storagePushConstant8"
  target="_blank" rel="noopener"><code>storagePushConstant8</code></a>
  is `VK_FALSE`, then objects containing an 8-bit integer element
  **must** not have `Storage` `Class` of `PushConstant`

- <a href="#VUID-RuntimeSpirv-storageBuffer16BitAccess-06331"
  id="VUID-RuntimeSpirv-storageBuffer16BitAccess-06331"></a>
  VUID-RuntimeSpirv-storageBuffer16BitAccess-06331  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-storageBuffer16BitAccess"
  target="_blank" rel="noopener"><code>storageBuffer16BitAccess</code></a>
  is `VK_FALSE`, then objects containing 16-bit integer or 16-bit
  floating-point elements **must** not have `Storage` `Class` of
  `StorageBuffer`, `ShaderRecordBufferKHR`, or `PhysicalStorageBuffer`

- <a href="#VUID-RuntimeSpirv-uniformAndStorageBuffer16BitAccess-06332"
  id="VUID-RuntimeSpirv-uniformAndStorageBuffer16BitAccess-06332"></a>
  VUID-RuntimeSpirv-uniformAndStorageBuffer16BitAccess-06332  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-uniformAndStorageBuffer16BitAccess"
  target="_blank"
  rel="noopener"><code>uniformAndStorageBuffer16BitAccess</code></a> is
  `VK_FALSE`, then objects in the `Uniform` `Storage` `Class` with the
  `Block` decoration **must** not have 16-bit integer or 16-bit
  floating-point members

- <a href="#VUID-RuntimeSpirv-storagePushConstant16-06333"
  id="VUID-RuntimeSpirv-storagePushConstant16-06333"></a>
  VUID-RuntimeSpirv-storagePushConstant16-06333  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-storagePushConstant16"
  target="_blank" rel="noopener"><code>storagePushConstant16</code></a>
  is `VK_FALSE`, then objects containing 16-bit integer or 16-bit
  floating-point elements **must** not have `Storage` `Class` of
  `PushConstant`

- <a href="#VUID-RuntimeSpirv-storageInputOutput16-06334"
  id="VUID-RuntimeSpirv-storageInputOutput16-06334"></a>
  VUID-RuntimeSpirv-storageInputOutput16-06334  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-storageInputOutput16"
  target="_blank" rel="noopener"><code>storageInputOutput16</code></a>
  is `VK_FALSE`, then objects containing 16-bit integer or 16-bit
  floating-point elements **must** not have `Storage` `Class` of `Input`
  or `Output`

- <a href="#VUID-RuntimeSpirv-None-06337"
  id="VUID-RuntimeSpirv-None-06337"></a> VUID-RuntimeSpirv-None-06337  
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderBufferFloat16Atomics"
  target="_blank"
  rel="noopener"><code>shaderBufferFloat16Atomics</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderBufferFloat16AtomicAdd"
  target="_blank"
  rel="noopener"><code>shaderBufferFloat16AtomicAdd</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderBufferFloat16AtomicMinMax"
  target="_blank"
  rel="noopener"><code>shaderBufferFloat16AtomicMinMax</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderSharedFloat16Atomics"
  target="_blank"
  rel="noopener"><code>shaderSharedFloat16Atomics</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderSharedFloat16AtomicAdd"
  target="_blank"
  rel="noopener"><code>shaderSharedFloat16AtomicAdd</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderSharedFloat16AtomicMinMax"
  target="_blank"
  rel="noopener"><code>shaderSharedFloat16AtomicMinMax</code></a>
  **must** be enabled for 16-bit floating-point atomic operations

- <a href="#VUID-RuntimeSpirv-None-06338"
  id="VUID-RuntimeSpirv-None-06338"></a> VUID-RuntimeSpirv-None-06338  
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderBufferFloat32Atomics"
  target="_blank"
  rel="noopener"><code>shaderBufferFloat32Atomics</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderBufferFloat32AtomicAdd"
  target="_blank"
  rel="noopener"><code>shaderBufferFloat32AtomicAdd</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderSharedFloat32Atomics"
  target="_blank"
  rel="noopener"><code>shaderSharedFloat32Atomics</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderSharedFloat32AtomicAdd"
  target="_blank"
  rel="noopener"><code>shaderSharedFloat32AtomicAdd</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderImageFloat32Atomics"
  target="_blank"
  rel="noopener"><code>shaderImageFloat32Atomics</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderImageFloat32AtomicAdd"
  target="_blank"
  rel="noopener"><code>shaderImageFloat32AtomicAdd</code></a> or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderBufferFloat32AtomicMinMax"
  target="_blank"
  rel="noopener"><code>shaderBufferFloat32AtomicMinMax</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderSharedFloat32AtomicMinMax"
  target="_blank"
  rel="noopener"><code>shaderSharedFloat32AtomicMinMax</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderImageFloat32AtomicMinMax"
  target="_blank"
  rel="noopener"><code>shaderImageFloat32AtomicMinMax</code></a>
  **must** be enabled for 32-bit floating-point atomic operations

- <a href="#VUID-RuntimeSpirv-None-06339"
  id="VUID-RuntimeSpirv-None-06339"></a> VUID-RuntimeSpirv-None-06339  
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderBufferFloat64Atomics"
  target="_blank"
  rel="noopener"><code>shaderBufferFloat64Atomics</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderBufferFloat64AtomicAdd"
  target="_blank"
  rel="noopener"><code>shaderBufferFloat64AtomicAdd</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderSharedFloat64Atomics"
  target="_blank"
  rel="noopener"><code>shaderSharedFloat64Atomics</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderSharedFloat64AtomicAdd"
  target="_blank"
  rel="noopener"><code>shaderSharedFloat64AtomicAdd</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderBufferFloat64AtomicMinMax"
  target="_blank"
  rel="noopener"><code>shaderBufferFloat64AtomicMinMax</code></a>, or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderSharedFloat64AtomicMinMax"
  target="_blank"
  rel="noopener"><code>shaderSharedFloat64AtomicMinMax</code></a>,
  **must** be enabled for 64-bit floating-point atomic operations

- <a href="#VUID-RuntimeSpirv-shaderFloat16VectorAtomics-09581"
  id="VUID-RuntimeSpirv-shaderFloat16VectorAtomics-09581"></a>
  VUID-RuntimeSpirv-shaderFloat16VectorAtomics-09581  
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderFloat16VectorAtomics"
  target="_blank"
  rel="noopener"><code>shaderFloat16VectorAtomics</code></a>, **must**
  be enabled for 16-bit floating-point, 2- and 4-component vector atomic
  operations to be supported

- <a href="#VUID-RuntimeSpirv-NonWritable-06340"
  id="VUID-RuntimeSpirv-NonWritable-06340"></a>
  VUID-RuntimeSpirv-NonWritable-06340  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-fragmentStoresAndAtomics"
  target="_blank" rel="noopener"><code>fragmentStoresAndAtomics</code></a>
  is not enabled, then all storage image, storage texel buffer, and
  storage buffer variables in the fragment stage **must** be decorated
  with the `NonWritable` decoration

- <a href="#VUID-RuntimeSpirv-NonWritable-06341"
  id="VUID-RuntimeSpirv-NonWritable-06341"></a>
  VUID-RuntimeSpirv-NonWritable-06341  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-vertexPipelineStoresAndAtomics"
  target="_blank"
  rel="noopener"><code>vertexPipelineStoresAndAtomics</code></a> is not
  enabled, then all storage image, storage texel buffer, and storage
  buffer variables in the vertex, tessellation, and geometry stages
  **must** be decorated with the `NonWritable` decoration

- <a href="#VUID-RuntimeSpirv-None-06342"
  id="VUID-RuntimeSpirv-None-06342"></a> VUID-RuntimeSpirv-None-06342  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-subgroupQuadOperationsInAllStages"
  target="_blank"
  rel="noopener"><code>subgroupQuadOperationsInAllStages</code></a> is
  `VK_FALSE`, then <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-subgroup-quad"
  target="_blank" rel="noopener">quad subgroup operations</a> **must**
  not be used except for in fragment and compute stages

- <a href="#VUID-RuntimeSpirv-None-06343"
  id="VUID-RuntimeSpirv-None-06343"></a> VUID-RuntimeSpirv-None-06343  
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-group-operations"
  target="_blank" rel="noopener">Group operations</a> with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-scope-subgroup"
  target="_blank" rel="noopener">subgroup scope</a> **must** not be used
  if the shader stage is not in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-subgroupSupportedStages"
  target="_blank" rel="noopener"><code>subgroupSupportedStages</code></a>

- <a href="#VUID-RuntimeSpirv-Offset-06344"
  id="VUID-RuntimeSpirv-Offset-06344"></a>
  VUID-RuntimeSpirv-Offset-06344  
  The first element of the `Offset` operand of `InterpolateAtOffset`
  **must** be greater than or equal to:  
  frag<sub>width</sub> × <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-minInterpolationOffset"
  target="_blank" rel="noopener"><code>minInterpolationOffset</code></a>  
  where frag<sub>width</sub> is the width of the current fragment in
  pixels

- <a href="#VUID-RuntimeSpirv-Offset-06345"
  id="VUID-RuntimeSpirv-Offset-06345"></a>
  VUID-RuntimeSpirv-Offset-06345  
  The first element of the `Offset` operand of `InterpolateAtOffset`
  **must** be less than or equal to  
  frag<sub>width</sub> × (<a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxInterpolationOffset"
  target="_blank" rel="noopener"><code>maxInterpolationOffset</code></a> +
  ULP ) - ULP  
  where frag<sub>width</sub> is the width of the current fragment in
  pixels and
  ULP = 1 / 2
  [`subPixelInterpolationOffsetBits`](https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-subPixelInterpolationOffsetBits)^

- <a href="#VUID-RuntimeSpirv-Offset-06346"
  id="VUID-RuntimeSpirv-Offset-06346"></a>
  VUID-RuntimeSpirv-Offset-06346  
  The second element of the `Offset` operand of `InterpolateAtOffset`
  **must** be greater than or equal to  
  frag<sub>height</sub> × <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-minInterpolationOffset"
  target="_blank" rel="noopener"><code>minInterpolationOffset</code></a>  
  where frag<sub>height</sub> is the height of the current fragment in
  pixels

- <a href="#VUID-RuntimeSpirv-Offset-06347"
  id="VUID-RuntimeSpirv-Offset-06347"></a>
  VUID-RuntimeSpirv-Offset-06347  
  The second element of the `Offset` operand of `InterpolateAtOffset`
  **must** be less than or equal to  
  frag<sub>height</sub> × (<a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxInterpolationOffset"
  target="_blank" rel="noopener"><code>maxInterpolationOffset</code></a> +
  ULP ) - ULP  
  where frag<sub>height</sub> is the height of the current fragment in
  pixels and
  ULP = 1 / 2
  [`subPixelInterpolationOffsetBits`](https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-subPixelInterpolationOffsetBits)^

- <a href="#VUID-RuntimeSpirv-OpRayQueryInitializeKHR-06348"
  id="VUID-RuntimeSpirv-OpRayQueryInitializeKHR-06348"></a>
  VUID-RuntimeSpirv-OpRayQueryInitializeKHR-06348  
  For `OpRayQueryInitializeKHR` instructions, all components of the
  `RayOrigin` and `RayDirection` operands **must** be finite
  floating-point values

- <a href="#VUID-RuntimeSpirv-OpRayQueryInitializeKHR-06349"
  id="VUID-RuntimeSpirv-OpRayQueryInitializeKHR-06349"></a>
  VUID-RuntimeSpirv-OpRayQueryInitializeKHR-06349  
  For `OpRayQueryInitializeKHR` instructions, the `RayTmin` and
  `RayTmax` operands **must** be non-negative floating-point values

- <a href="#VUID-RuntimeSpirv-OpRayQueryInitializeKHR-06350"
  id="VUID-RuntimeSpirv-OpRayQueryInitializeKHR-06350"></a>
  VUID-RuntimeSpirv-OpRayQueryInitializeKHR-06350  
  For `OpRayQueryInitializeKHR` instructions, the `RayTmin` operand
  **must** be less than or equal to the `RayTmax` operand

- <a href="#VUID-RuntimeSpirv-OpRayQueryInitializeKHR-06351"
  id="VUID-RuntimeSpirv-OpRayQueryInitializeKHR-06351"></a>
  VUID-RuntimeSpirv-OpRayQueryInitializeKHR-06351  
  For `OpRayQueryInitializeKHR` instructions, `RayOrigin`,
  `RayDirection`, `RayTmin`, and `RayTmax` operands **must** not contain
  NaNs

- <a href="#VUID-RuntimeSpirv-OpRayQueryInitializeKHR-06352"
  id="VUID-RuntimeSpirv-OpRayQueryInitializeKHR-06352"></a>
  VUID-RuntimeSpirv-OpRayQueryInitializeKHR-06352  
  For `OpRayQueryInitializeKHR` instructions, `Acceleration` `Structure`
  **must** be an acceleration structure built as a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#acceleration-structure-top-level"
  target="_blank" rel="noopener">top-level acceleration structure</a>

- <a href="#VUID-RuntimeSpirv-OpRayQueryInitializeKHR-06889"
  id="VUID-RuntimeSpirv-OpRayQueryInitializeKHR-06889"></a>
  VUID-RuntimeSpirv-OpRayQueryInitializeKHR-06889  
  For `OpRayQueryInitializeKHR` instructions, the `Rayflags` operand
  **must** not contain both `SkipTrianglesKHR` and `SkipAABBsKHR`

- <a href="#VUID-RuntimeSpirv-OpRayQueryInitializeKHR-06890"
  id="VUID-RuntimeSpirv-OpRayQueryInitializeKHR-06890"></a>
  VUID-RuntimeSpirv-OpRayQueryInitializeKHR-06890  
  For `OpRayQueryInitializeKHR` instructions, the `Rayflags` operand
  **must** not contain more than one of `SkipTrianglesKHR`,
  `CullBackFacingTrianglesKHR`, and `CullFrontFacingTrianglesKHR`

- <a href="#VUID-RuntimeSpirv-OpRayQueryInitializeKHR-06891"
  id="VUID-RuntimeSpirv-OpRayQueryInitializeKHR-06891"></a>
  VUID-RuntimeSpirv-OpRayQueryInitializeKHR-06891  
  For `OpRayQueryInitializeKHR` instructions, the `Rayflags` operand
  **must** not contain more than one of `OpaqueKHR`, `NoOpaqueKHR`,
  `CullOpaqueKHR`, and `CullNoOpaqueKHR`

- <a href="#VUID-RuntimeSpirv-OpRayQueryGenerateIntersectionKHR-06353"
  id="VUID-RuntimeSpirv-OpRayQueryGenerateIntersectionKHR-06353"></a>
  VUID-RuntimeSpirv-OpRayQueryGenerateIntersectionKHR-06353  
  For `OpRayQueryGenerateIntersectionKHR` instructions, `Hit` `T`
  **must** satisfy the condition `RayTmin` ≤ `Hit` `T` ≤ `RayTmax`,
  where `RayTmin` is equal to the value returned by
  `OpRayQueryGetRayTMinKHR` with the same ray query object, and
  `RayTmax` is equal to the value of `OpRayQueryGetIntersectionTKHR` for
  the current committed intersection with the same ray query object

- <a href="#VUID-RuntimeSpirv-OpRayQueryGenerateIntersectionKHR-06354"
  id="VUID-RuntimeSpirv-OpRayQueryGenerateIntersectionKHR-06354"></a>
  VUID-RuntimeSpirv-OpRayQueryGenerateIntersectionKHR-06354  
  For `OpRayQueryGenerateIntersectionKHR` instructions, `Acceleration`
  `Structure` **must** not be built with
  `VK_BUILD_ACCELERATION_STRUCTURE_MOTION_BIT_NV` in `flags`

- <a href="#VUID-RuntimeSpirv-flags-08761"
  id="VUID-RuntimeSpirv-flags-08761"></a>
  VUID-RuntimeSpirv-flags-08761  
  For `OpRayQueryGetIntersectionTriangleVertexPositionsKHR`
  instructions, `Acceleration` `Structure` **must** have been built with
  `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_DATA_ACCESS_KHR` in `flags`

- <a href="#VUID-RuntimeSpirv-OpTraceRayKHR-06355"
  id="VUID-RuntimeSpirv-OpTraceRayKHR-06355"></a>
  VUID-RuntimeSpirv-OpTraceRayKHR-06355  
  For `OpTraceRayKHR` instructions, all components of the `RayOrigin`
  and `RayDirection` operands **must** be finite floating-point values

- <a href="#VUID-RuntimeSpirv-OpTraceRayKHR-06356"
  id="VUID-RuntimeSpirv-OpTraceRayKHR-06356"></a>
  VUID-RuntimeSpirv-OpTraceRayKHR-06356  
  For `OpTraceRayKHR` instructions, the `RayTmin` and `RayTmax` operands
  **must** be non-negative floating-point values

- <a href="#VUID-RuntimeSpirv-OpTraceRayKHR-06552"
  id="VUID-RuntimeSpirv-OpTraceRayKHR-06552"></a>
  VUID-RuntimeSpirv-OpTraceRayKHR-06552  
  For `OpTraceRayKHR` instructions, the `Rayflags` operand **must** not
  contain both `SkipTrianglesKHR` and `SkipAABBsKHR`

- <a href="#VUID-RuntimeSpirv-OpTraceRayKHR-06892"
  id="VUID-RuntimeSpirv-OpTraceRayKHR-06892"></a>
  VUID-RuntimeSpirv-OpTraceRayKHR-06892  
  For `OpTraceRayKHR` instructions, the `Rayflags` operand **must** not
  contain more than one of `SkipTrianglesKHR`,
  `CullBackFacingTrianglesKHR`, and `CullFrontFacingTrianglesKHR`

- <a href="#VUID-RuntimeSpirv-OpTraceRayKHR-06893"
  id="VUID-RuntimeSpirv-OpTraceRayKHR-06893"></a>
  VUID-RuntimeSpirv-OpTraceRayKHR-06893  
  For `OpTraceRayKHR` instructions, the `Rayflags` operand **must** not
  contain more than one of `OpaqueKHR`, `NoOpaqueKHR`, `CullOpaqueKHR`,
  and `CullNoOpaqueKHR`

- <a href="#VUID-RuntimeSpirv-OpTraceRayKHR-06553"
  id="VUID-RuntimeSpirv-OpTraceRayKHR-06553"></a>
  VUID-RuntimeSpirv-OpTraceRayKHR-06553  
  For `OpTraceRayKHR` instructions, if the `Rayflags` operand contains
  `SkipTrianglesKHR`, the pipeline **must** not have been created with
  `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_AABBS_BIT_KHR` set

- <a href="#VUID-RuntimeSpirv-OpTraceRayKHR-06554"
  id="VUID-RuntimeSpirv-OpTraceRayKHR-06554"></a>
  VUID-RuntimeSpirv-OpTraceRayKHR-06554  
  For `OpTraceRayKHR` instructions, if the `Rayflags` operand contains
  `SkipAABBsKHR`, the pipeline **must** not have been created with
  `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR` set

- <a href="#VUID-RuntimeSpirv-OpTraceRayKHR-06357"
  id="VUID-RuntimeSpirv-OpTraceRayKHR-06357"></a>
  VUID-RuntimeSpirv-OpTraceRayKHR-06357  
  For `OpTraceRayKHR` instructions, the `RayTmin` operand **must** be
  less than or equal to the `RayTmax` operand

- <a href="#VUID-RuntimeSpirv-OpTraceRayKHR-06358"
  id="VUID-RuntimeSpirv-OpTraceRayKHR-06358"></a>
  VUID-RuntimeSpirv-OpTraceRayKHR-06358  
  For `OpTraceRayKHR` instructions, `RayOrigin`, `RayDirection`,
  `RayTmin`, and `RayTmax` operands **must** not contain NaNs

- <a href="#VUID-RuntimeSpirv-OpTraceRayKHR-06359"
  id="VUID-RuntimeSpirv-OpTraceRayKHR-06359"></a>
  VUID-RuntimeSpirv-OpTraceRayKHR-06359  
  For `OpTraceRayKHR` instructions, `Acceleration` `Structure` **must**
  be an acceleration structure built as a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#acceleration-structure-top-level"
  target="_blank" rel="noopener">top-level acceleration structure</a>

- <a href="#VUID-RuntimeSpirv-OpReportIntersectionKHR-06998"
  id="VUID-RuntimeSpirv-OpReportIntersectionKHR-06998"></a>
  VUID-RuntimeSpirv-OpReportIntersectionKHR-06998  
  The value of the “Hit Kind” operand of `OpReportIntersectionKHR`
  **must** be in the range \[0,127\]

- <a href="#VUID-RuntimeSpirv-OpTraceRayKHR-06360"
  id="VUID-RuntimeSpirv-OpTraceRayKHR-06360"></a>
  VUID-RuntimeSpirv-OpTraceRayKHR-06360  
  For `OpTraceRayKHR` instructions, if `Acceleration` `Structure` was
  built with `VK_BUILD_ACCELERATION_STRUCTURE_MOTION_BIT_NV` in `flags`,
  the pipeline **must** have been created with
  `VK_PIPELINE_CREATE_RAY_TRACING_ALLOW_MOTION_BIT_NV` set

- <a href="#VUID-RuntimeSpirv-OpTraceRayMotionNV-06361"
  id="VUID-RuntimeSpirv-OpTraceRayMotionNV-06361"></a>
  VUID-RuntimeSpirv-OpTraceRayMotionNV-06361  
  For `OpTraceRayMotionNV` instructions, all components of the
  `RayOrigin` and `RayDirection` operands **must** be finite
  floating-point values

- <a href="#VUID-RuntimeSpirv-OpTraceRayMotionNV-06362"
  id="VUID-RuntimeSpirv-OpTraceRayMotionNV-06362"></a>
  VUID-RuntimeSpirv-OpTraceRayMotionNV-06362  
  For `OpTraceRayMotionNV` instructions, the `RayTmin` and `RayTmax`
  operands **must** be non-negative floating-point values

- <a href="#VUID-RuntimeSpirv-OpTraceRayMotionNV-06363"
  id="VUID-RuntimeSpirv-OpTraceRayMotionNV-06363"></a>
  VUID-RuntimeSpirv-OpTraceRayMotionNV-06363  
  For `OpTraceRayMotionNV` instructions, the `RayTmin` operand **must**
  be less than or equal to the `RayTmax` operand

- <a href="#VUID-RuntimeSpirv-OpTraceRayMotionNV-06364"
  id="VUID-RuntimeSpirv-OpTraceRayMotionNV-06364"></a>
  VUID-RuntimeSpirv-OpTraceRayMotionNV-06364  
  For `OpTraceRayMotionNV` instructions, `RayOrigin`, `RayDirection`,
  `RayTmin`, and `RayTmax` operands **must** not contain NaNs

- <a href="#VUID-RuntimeSpirv-OpTraceRayMotionNV-06365"
  id="VUID-RuntimeSpirv-OpTraceRayMotionNV-06365"></a>
  VUID-RuntimeSpirv-OpTraceRayMotionNV-06365  
  For `OpTraceRayMotionNV` instructions, `Acceleration` `Structure`
  **must** be an acceleration structure built as a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#acceleration-structure-top-level"
  target="_blank" rel="noopener">top-level acceleration structure</a>
  with `VK_BUILD_ACCELERATION_STRUCTURE_MOTION_BIT_NV` in `flags`

- <a href="#VUID-RuntimeSpirv-OpTraceRayMotionNV-06366"
  id="VUID-RuntimeSpirv-OpTraceRayMotionNV-06366"></a>
  VUID-RuntimeSpirv-OpTraceRayMotionNV-06366  
  For `OpTraceRayMotionNV` instructions the `time` operand **must** be
  between 0.0 and 1.0

- <a href="#VUID-RuntimeSpirv-OpTraceRayMotionNV-06367"
  id="VUID-RuntimeSpirv-OpTraceRayMotionNV-06367"></a>
  VUID-RuntimeSpirv-OpTraceRayMotionNV-06367  
  For `OpTraceRayMotionNV` instructions the pipeline **must** have been
  created with `VK_PIPELINE_CREATE_RAY_TRACING_ALLOW_MOTION_BIT_NV` set

- <a href="#VUID-RuntimeSpirv-OpHitObjectTraceRayMotionNV-07704"
  id="VUID-RuntimeSpirv-OpHitObjectTraceRayMotionNV-07704"></a>
  VUID-RuntimeSpirv-OpHitObjectTraceRayMotionNV-07704  
  For `OpHitObjectTraceRayMotionNV` instructions, if `Acceleration`
  `Structure` was built with
  `VK_BUILD_ACCELERATION_STRUCTURE_MOTION_BIT_NV` in `flags`, the
  pipeline **must** have been created with
  `VK_PIPELINE_CREATE_RAY_TRACING_ALLOW_MOTION_BIT_NV` set

- <a href="#VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07705"
  id="VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07705"></a>
  VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07705  
  For `OpHitObjectTraceRayNV` and `OpHitObjectTraceRayMotionNV`
  instructions, all components of the `RayOrigin` and `RayDirection`
  operands **must** be finite floating-point values

- <a href="#VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07706"
  id="VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07706"></a>
  VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07706  
  For `OpHitObjectTraceRayNV` and `OpHitObjectTraceRayMotionNV`
  instructions, the `RayTmin` and `RayTmax` operands **must** be
  non-negative floating-point values

- <a href="#VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07707"
  id="VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07707"></a>
  VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07707  
  For `OpHitObjectTraceRayNV` and `OpHitObjectTraceRayMotionNV`
  instructions, the `RayTmin` operand **must** be less than or equal to
  the `RayTmax` operand

- <a href="#VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07708"
  id="VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07708"></a>
  VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07708  
  For `OpHitObjectTraceRayNV` and `OpHitObjectTraceRayMotionNV`
  instructions, `RayOrigin`, `RayDirection`, `RayTmin`, and `RayTmax`
  operands **must** not contain NaNs

- <a href="#VUID-RuntimeSpirv-OpHitObjectTraceRayMotionNV-07709"
  id="VUID-RuntimeSpirv-OpHitObjectTraceRayMotionNV-07709"></a>
  VUID-RuntimeSpirv-OpHitObjectTraceRayMotionNV-07709  
  For `OpHitObjectTraceRayMotionNV` instructions, `Acceleration`
  `Structure` **must** be an acceleration structure built as a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#acceleration-structure-top-level"
  target="_blank" rel="noopener">top-level acceleration structure</a>
  with `VK_BUILD_ACCELERATION_STRUCTURE_MOTION_BIT_NV` in `flags`

- <a href="#VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07710"
  id="VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07710"></a>
  VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07710  
  For `OpHitObjectTraceRayNV` and `OpHitObjectTraceRayMotionNV`
  instructions the `time` operand **must** be between 0.0 and 1.0

- <a href="#VUID-RuntimeSpirv-OpHitObjectTraceRayMotionNV-07711"
  id="VUID-RuntimeSpirv-OpHitObjectTraceRayMotionNV-07711"></a>
  VUID-RuntimeSpirv-OpHitObjectTraceRayMotionNV-07711  
  For `OpHitObjectTraceRayMotionNV` instructions the pipeline **must**
  have been created with
  `VK_PIPELINE_CREATE_RAY_TRACING_ALLOW_MOTION_BIT_NV` set

- <a href="#VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07712"
  id="VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07712"></a>
  VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07712  
  For `OpHitObjectTraceRayNV` and `OpHitObjectTraceRayMotionNV`
  instructions, the `Rayflags` operand **must** not contain both
  `SkipTrianglesKHR` and `SkipAABBsKHR`

- <a href="#VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07713"
  id="VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07713"></a>
  VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07713  
  For `OpHitObjectTraceRayNV` and `OpHitObjectTraceRayMotionNV`
  instructions, the `Rayflags` operand **must** not contain more than
  one of `SkipTrianglesKHR`, `CullBackFacingTrianglesKHR`, and
  `CullFrontFacingTrianglesKHR`

- <a href="#VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07714"
  id="VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07714"></a>
  VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07714  
  For `OpHitObjectTraceRayNV` and `OpHitObjectTraceRayMotionNV`
  instructions, the `Rayflags` operand **must** not contain more than
  one of `OpaqueKHR`, `NoOpaqueKHR`, `CullOpaqueKHR`, and
  `CullNoOpaqueKHR`

- <a href="#VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07715"
  id="VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07715"></a>
  VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07715  
  For `OpHitObjectTraceRayNV` and `OpHitObjectTraceRayMotionNV`
  instructions, if the `Rayflags` operand contains `SkipTrianglesKHR`,
  the pipeline **must** not have been created with
  `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_AABBS_BIT_KHR` set

- <a href="#VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07716"
  id="VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07716"></a>
  VUID-RuntimeSpirv-OpHitObjectTraceRayNV-07716  
  For `OpHitObjectTraceRayNV` and `OpHitObjectTraceRayMotionNV`
  instructions, if the `Rayflags` operand contains `SkipAABBsKHR`, the
  pipeline **must** not have been created with
  `VK_PIPELINE_CREATE_RAY_TRACING_SKIP_TRIANGLES_BIT_KHR` set

- <a href="#VUID-RuntimeSpirv-x-06429" id="VUID-RuntimeSpirv-x-06429"></a>
  VUID-RuntimeSpirv-x-06429  
  In compute shaders using the `GLCompute` `Execution` `Model` the `x`
  size in `LocalSize` or `LocalSizeId` **must** be less than or equal to
  [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)::`maxComputeWorkGroupSize`\[0\]

- <a href="#VUID-RuntimeSpirv-y-06430" id="VUID-RuntimeSpirv-y-06430"></a>
  VUID-RuntimeSpirv-y-06430  
  In compute shaders using the `GLCompute` `Execution` `Model` the `y`
  size in `LocalSize` or `LocalSizeId` **must** be less than or equal to
  [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)::`maxComputeWorkGroupSize`\[1\]

- <a href="#VUID-RuntimeSpirv-z-06431" id="VUID-RuntimeSpirv-z-06431"></a>
  VUID-RuntimeSpirv-z-06431  
  In compute shaders using the `GLCompute` `Execution` `Model` the `z`
  size in `LocalSize` or `LocalSizeId` **must** be less than or equal to
  [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)::`maxComputeWorkGroupSize`\[2\]

- <a href="#VUID-RuntimeSpirv-x-06432" id="VUID-RuntimeSpirv-x-06432"></a>
  VUID-RuntimeSpirv-x-06432  
  In compute shaders using the `GLCompute` `Execution` `Model` the
  product of `x` size, `y` size, and `z` size in `LocalSize` or
  `LocalSizeId` **must** be less than or equal to
  [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)::`maxComputeWorkGroupInvocations`

- <a href="#VUID-RuntimeSpirv-LocalSizeId-06434"
  id="VUID-RuntimeSpirv-LocalSizeId-06434"></a>
  VUID-RuntimeSpirv-LocalSizeId-06434  
  If `Execution` `Mode` `LocalSizeId` is used, <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance4"
  target="_blank" rel="noopener"><code>maintenance4</code></a> **must**
  be enabled

- <a href="#VUID-RuntimeSpirv-maintenance4-06817"
  id="VUID-RuntimeSpirv-maintenance4-06817"></a>
  VUID-RuntimeSpirv-maintenance4-06817  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance4"
  target="_blank" rel="noopener"><code>maintenance4</code></a> is not
  enabled, any `OpTypeVector` output interface variables **must** not
  have a higher `Component` `Count` than a matching `OpTypeVector` input
  interface variable

- <a href="#VUID-RuntimeSpirv-OpEntryPoint-08743"
  id="VUID-RuntimeSpirv-OpEntryPoint-08743"></a>
  VUID-RuntimeSpirv-OpEntryPoint-08743  
  Any <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-iointerfaces-user"
  target="_blank" rel="noopener">user-defined variables</a> shared
  between the `OpEntryPoint` of two shader stages, and declared with
  `Input` as its `Storage` `Class` for the subsequent shader stage,
  **must** have all `Location` slots and `Component` words declared in
  the preceding shader stage’s `OpEntryPoint` with `Output` as the
  `Storage` `Class`

- <a href="#VUID-RuntimeSpirv-OpEntryPoint-07754"
  id="VUID-RuntimeSpirv-OpEntryPoint-07754"></a>
  VUID-RuntimeSpirv-OpEntryPoint-07754  
  Any <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-iointerfaces-user"
  target="_blank" rel="noopener">user-defined variables</a> between the
  `OpEntryPoint` of two shader stages **must** have the same type and
  width for each `Component`

- <a href="#VUID-RuntimeSpirv-OpVariable-08746"
  id="VUID-RuntimeSpirv-OpVariable-08746"></a>
  VUID-RuntimeSpirv-OpVariable-08746  
  Any `OpVariable`, `Block`-decorated `OpTypeStruct`, or
  `Block`-decorated `OpTypeStruct` members shared between the
  `OpEntryPoint` of two shader stages **must** have matching decorations
  as defined in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-iointerfaces-matching"
  target="_blank" rel="noopener">interface matching</a>

- <a href="#VUID-RuntimeSpirv-Workgroup-06530"
  id="VUID-RuntimeSpirv-Workgroup-06530"></a>
  VUID-RuntimeSpirv-Workgroup-06530  
  The sum of size in bytes for variables and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#workgroup-padding"
  target="_blank" rel="noopener">padding</a> in the `Workgroup`
  `Storage` `Class` in the `GLCompute` `Execution` `Model` **must** be
  less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxComputeSharedMemorySize"
  target="_blank"
  rel="noopener"><code>maxComputeSharedMemorySize</code></a>

- <a href="#VUID-RuntimeSpirv-shaderZeroInitializeWorkgroupMemory-06372"
  id="VUID-RuntimeSpirv-shaderZeroInitializeWorkgroupMemory-06372"></a>
  VUID-RuntimeSpirv-shaderZeroInitializeWorkgroupMemory-06372  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderZeroInitializeWorkgroupMemory"
  target="_blank"
  rel="noopener"><code>shaderZeroInitializeWorkgroupMemory</code></a> is
  not enabled, any `OpVariable` with `Workgroup` as its `Storage`
  `Class` **must** not have an `Initializer` operand

- <a href="#VUID-RuntimeSpirv-OpImage-06376"
  id="VUID-RuntimeSpirv-OpImage-06376"></a>
  VUID-RuntimeSpirv-OpImage-06376  
  If an `OpImage*Gather` operation has an image operand of `Offset`,
  `ConstOffset`, or `ConstOffsets` the offset value **must** be greater
  than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-minTexelGatherOffset"
  target="_blank" rel="noopener"><code>minTexelGatherOffset</code></a>

- <a href="#VUID-RuntimeSpirv-OpImage-06377"
  id="VUID-RuntimeSpirv-OpImage-06377"></a>
  VUID-RuntimeSpirv-OpImage-06377  
  If an `OpImage*Gather` operation has an image operand of `Offset`,
  `ConstOffset`, or `ConstOffsets` the offset value **must** be less
  than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxTexelGatherOffset"
  target="_blank" rel="noopener"><code>maxTexelGatherOffset</code></a>

- <a href="#VUID-RuntimeSpirv-OpImageSample-06435"
  id="VUID-RuntimeSpirv-OpImageSample-06435"></a>
  VUID-RuntimeSpirv-OpImageSample-06435  
  If an `OpImageSample*` or `OpImageFetch*` operation has an image
  operand of `ConstOffset` then the offset value **must** be greater
  than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-minTexelOffset"
  target="_blank" rel="noopener"><code>minTexelOffset</code></a>

- <a href="#VUID-RuntimeSpirv-OpImageSample-06436"
  id="VUID-RuntimeSpirv-OpImageSample-06436"></a>
  VUID-RuntimeSpirv-OpImageSample-06436  
  If an `OpImageSample*` or `OpImageFetch*` operation has an image
  operand of `ConstOffset` then the offset value **must** be less than
  or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxTexelOffset"
  target="_blank" rel="noopener"><code>maxTexelOffset</code></a>

- <a href="#VUID-RuntimeSpirv-samples-08725"
  id="VUID-RuntimeSpirv-samples-08725"></a>
  VUID-RuntimeSpirv-samples-08725  
  If an `OpTypeImage` has an `MS` operand 0, its bound image **must**
  have been created with
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`samples` as
  `VK_SAMPLE_COUNT_1_BIT`

- <a href="#VUID-RuntimeSpirv-samples-08726"
  id="VUID-RuntimeSpirv-samples-08726"></a>
  VUID-RuntimeSpirv-samples-08726  
  If an `OpTypeImage` has an `MS` operand 1, its bound image **must**
  not have been created with
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`samples` as
  `VK_SAMPLE_COUNT_1_BIT`

- <a href="#VUID-RuntimeSpirv-SampleRateShading-06378"
  id="VUID-RuntimeSpirv-SampleRateShading-06378"></a>
  VUID-RuntimeSpirv-SampleRateShading-06378  
  If the subpass description contains
  `VK_SUBPASS_DESCRIPTION_FRAGMENT_REGION_BIT_QCOM`, then the SPIR-V
  fragment shader Capability `SampleRateShading` **must** not be enabled

- <a href="#VUID-RuntimeSpirv-SubgroupUniformControlFlowKHR-06379"
  id="VUID-RuntimeSpirv-SubgroupUniformControlFlowKHR-06379"></a>
  VUID-RuntimeSpirv-SubgroupUniformControlFlowKHR-06379  
  The `Execution` `Mode` `SubgroupUniformControlFlowKHR` **must** not be
  applied to an entry point unless <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderSubgroupUniformControlFlow"
  target="_blank"
  rel="noopener"><code>shaderSubgroupUniformControlFlow</code></a> is
  enabled, the corresponding shader stage bit is set in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-subgroupSupportedStages"
  target="_blank" rel="noopener"><code>subgroupSupportedStages</code></a>,
  and the entry point does not execute any <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ray-tracing-repack"
  target="_blank" rel="noopener"><em>invocation repack
  instructions</em></a>

- <a href="#VUID-RuntimeSpirv-shaderEarlyAndLateFragmentTests-06767"
  id="VUID-RuntimeSpirv-shaderEarlyAndLateFragmentTests-06767"></a>
  VUID-RuntimeSpirv-shaderEarlyAndLateFragmentTests-06767  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderEarlyAndLateFragmentTests"
  target="_blank"
  rel="noopener"><code>shaderEarlyAndLateFragmentTests</code></a> is not
  enabled, the `EarlyAndLateFragmentTestsEXT` `Execution` `Mode`
  **must** not be used

- <a href="#VUID-RuntimeSpirv-shaderEarlyAndLateFragmentTests-06768"
  id="VUID-RuntimeSpirv-shaderEarlyAndLateFragmentTests-06768"></a>
  VUID-RuntimeSpirv-shaderEarlyAndLateFragmentTests-06768  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderEarlyAndLateFragmentTests"
  target="_blank"
  rel="noopener"><code>shaderEarlyAndLateFragmentTests</code></a>
  feature is not enabled, the `StencilRefUnchangedFrontEXT` `Execution`
  `Mode` **must** not be used

- <a href="#VUID-RuntimeSpirv-shaderEarlyAndLateFragmentTests-06769"
  id="VUID-RuntimeSpirv-shaderEarlyAndLateFragmentTests-06769"></a>
  VUID-RuntimeSpirv-shaderEarlyAndLateFragmentTests-06769  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderEarlyAndLateFragmentTests"
  target="_blank"
  rel="noopener"><code>shaderEarlyAndLateFragmentTests</code></a> is not
  enabled, the `StencilRefUnchangedBackEXT` `Execution` `Mode` **must**
  not be used

- <a href="#VUID-RuntimeSpirv-shaderEarlyAndLateFragmentTests-06770"
  id="VUID-RuntimeSpirv-shaderEarlyAndLateFragmentTests-06770"></a>
  VUID-RuntimeSpirv-shaderEarlyAndLateFragmentTests-06770  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderEarlyAndLateFragmentTests"
  target="_blank"
  rel="noopener"><code>shaderEarlyAndLateFragmentTests</code></a> is not
  enabled, the `StencilRefGreaterFrontEXT` `Execution` `Mode` **must**
  not be used

- <a href="#VUID-RuntimeSpirv-shaderEarlyAndLateFragmentTests-06771"
  id="VUID-RuntimeSpirv-shaderEarlyAndLateFragmentTests-06771"></a>
  VUID-RuntimeSpirv-shaderEarlyAndLateFragmentTests-06771  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderEarlyAndLateFragmentTests"
  target="_blank"
  rel="noopener"><code>shaderEarlyAndLateFragmentTests</code></a> is not
  enabled, the `StencilRefGreaterBackEXT` `Execution` `Mode` **must**
  not be used

- <a href="#VUID-RuntimeSpirv-shaderEarlyAndLateFragmentTests-06772"
  id="VUID-RuntimeSpirv-shaderEarlyAndLateFragmentTests-06772"></a>
  VUID-RuntimeSpirv-shaderEarlyAndLateFragmentTests-06772  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderEarlyAndLateFragmentTests"
  target="_blank"
  rel="noopener"><code>shaderEarlyAndLateFragmentTests</code></a> is not
  enabled, the `StencilRefLessFrontEXT` `Execution` `Mode` **must** not
  be used

- <a href="#VUID-RuntimeSpirv-shaderEarlyAndLateFragmentTests-06773"
  id="VUID-RuntimeSpirv-shaderEarlyAndLateFragmentTests-06773"></a>
  VUID-RuntimeSpirv-shaderEarlyAndLateFragmentTests-06773  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderEarlyAndLateFragmentTests"
  target="_blank"
  rel="noopener"><code>shaderEarlyAndLateFragmentTests</code></a> is not
  enabled, the `StencilRefLessBackEXT` `Execution` `Mode` **must** not
  be used

- <a href="#VUID-RuntimeSpirv-OpImageWeightedSampleQCOM-06979"
  id="VUID-RuntimeSpirv-OpImageWeightedSampleQCOM-06979"></a>
  VUID-RuntimeSpirv-OpImageWeightedSampleQCOM-06979  
  If an `OpImageWeightedSampleQCOM` operation is used, then the
  `Texture` `Sampled` `Image` and `Weight` `Image` parameters **must**
  both be *dynamically uniform* for the quad

- <a href="#VUID-RuntimeSpirv-OpImageWeightedSampleQCOM-06980"
  id="VUID-RuntimeSpirv-OpImageWeightedSampleQCOM-06980"></a>
  VUID-RuntimeSpirv-OpImageWeightedSampleQCOM-06980  
  If an `OpImageWeightedSampleQCOM` operation is used, then the `Weight`
  `Image` parameter **must** be of `Storage` `Class` `UniformConstant`
  and type `OpTypeImage` with `Depth`=0, `Dim`=`2D`, `Arrayed`=1,
  `MS`=0, and `Sampled`=1

- <a href="#VUID-RuntimeSpirv-OpImageWeightedSampleQCOM-06981"
  id="VUID-RuntimeSpirv-OpImageWeightedSampleQCOM-06981"></a>
  VUID-RuntimeSpirv-OpImageWeightedSampleQCOM-06981  
  If an `OpImageWeightedSampleQCOM` operation is used, then the `Weight`
  `Image` parameter **must** be decorated with `WeightTextureQCOM`

- <a href="#VUID-RuntimeSpirv-OpImageBlockMatchSADQCOM-06982"
  id="VUID-RuntimeSpirv-OpImageBlockMatchSADQCOM-06982"></a>
  VUID-RuntimeSpirv-OpImageBlockMatchSADQCOM-06982  
  If an `OpImageBlockMatchSADQCOM` or `OpImageBlockMatchSSDQCOM`
  operation is used, then the `target` `sampled` `image`, `reference`
  `sampled` `image`, and `Block` `Size` parameters **must** both be
  *dynamically uniform* for the quad

- <a href="#VUID-RuntimeSpirv-OpImageBlockMatchSSDQCOM-06983"
  id="VUID-RuntimeSpirv-OpImageBlockMatchSSDQCOM-06983"></a>
  VUID-RuntimeSpirv-OpImageBlockMatchSSDQCOM-06983  
  If an `OpImageBlockMatchSSDQCOM` or `OpImageBlockMatchSADQCOM`
  operation is used, then `target` `sampled` `image` and `reference`
  `sampled` `image` parameters **must** be of storage class
  `UniformConstant` and type `OpTypeImage` with `Depth`=0, `Dim`=`2D`,
  `Arrayed`=0, `MS`=0, and `Sampled`=1

- <a href="#VUID-RuntimeSpirv-OpImageBlockMatchSSDQCOM-06984"
  id="VUID-RuntimeSpirv-OpImageBlockMatchSSDQCOM-06984"></a>
  VUID-RuntimeSpirv-OpImageBlockMatchSSDQCOM-06984  
  If an `OpImageBlockMatchSSDQCOM` or `OpImageBlockMatchSADQCOM`
  operation is used, then the `target` `sampled` `image` and `reference`
  `sampled` `image` parameters **must** be decorated with
  `BlockMatchTextureQCOM`

- <a href="#VUID-RuntimeSpirv-OpImageBlockMatchSSDQCOM-06985"
  id="VUID-RuntimeSpirv-OpImageBlockMatchSSDQCOM-06985"></a>
  VUID-RuntimeSpirv-OpImageBlockMatchSSDQCOM-06985  
  If an `OpImageBlockMatchSSDQCOM` or `OpImageBlockMatchSADQCOM`
  operation is used, then `target` `sampled` `image` and `reference`
  `sampled` `image` parameters **must** have been created using an
  identical sampler object

- <a href="#VUID-RuntimeSpirv-OpImageBlockMatchSSDQCOM-06986"
  id="VUID-RuntimeSpirv-OpImageBlockMatchSSDQCOM-06986"></a>
  VUID-RuntimeSpirv-OpImageBlockMatchSSDQCOM-06986  
  If an `OpImageBlockMatchSSDQCOM` or `OpImageBlockMatchSADQCOM`
  operation is used, then `target` `sampled` `image` and `reference`
  `sampled` `image` parameters **must** have been created with a sampler
  object with `unnormalizedCoordinates` equal to `VK_TRUE`

- <a href="#VUID-RuntimeSpirv-OpImageBlockMatchSSDQCOM-06987"
  id="VUID-RuntimeSpirv-OpImageBlockMatchSSDQCOM-06987"></a>
  VUID-RuntimeSpirv-OpImageBlockMatchSSDQCOM-06987  
  If an `OpImageBlockMatchSSDQCOM` or `OpImageBlockMatchSADQCOM`
  operation is used, then `target` `sampled` `image` and `reference`
  `sampled` `image` parameters **must** have been created with a sampler
  object with `unnormalizedCoordinates` equal to `VK_TRUE`

- <a href="#VUID-RuntimeSpirv-OpImageBlockMatchSSDQCOM-06988"
  id="VUID-RuntimeSpirv-OpImageBlockMatchSSDQCOM-06988"></a>
  VUID-RuntimeSpirv-OpImageBlockMatchSSDQCOM-06988  
  If an `OpImageBlockMatchSSDQCOM` or `OpImageBlockMatchSADQCOM`
  operation is used, then `Block` `Size` less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-blockmatch-maxblocksize"
  target="_blank" rel="noopener"><code>maxBlockMatchRegion</code></a>

- <a href="#VUID-RuntimeSpirv-OpImageBoxFilterQCOM-06989"
  id="VUID-RuntimeSpirv-OpImageBoxFilterQCOM-06989"></a>
  VUID-RuntimeSpirv-OpImageBoxFilterQCOM-06989  
  If an `OpImageBoxFilterQCOM` operation is used, then `Box` `Size.y`
  **must** be equal to or greater than 1.0 and less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-boxfilter-maxblocksize"
  target="_blank" rel="noopener"><code>maxBoxFilterBlockSize</code></a>.`height`

- <a href="#VUID-RuntimeSpirv-OpImageBoxFilterQCOM-06990"
  id="VUID-RuntimeSpirv-OpImageBoxFilterQCOM-06990"></a>
  VUID-RuntimeSpirv-OpImageBoxFilterQCOM-06990  
  If an `OpImageBoxFilterQCOM` operation is used, then `Sampled`
  `Texture` `Image` and `Box` `Size` parameters **must** be *dynamically
  uniform*

- <a href="#VUID-RuntimeSpirv-OpEntryPoint-08727"
  id="VUID-RuntimeSpirv-OpEntryPoint-08727"></a>
  VUID-RuntimeSpirv-OpEntryPoint-08727  
  Each `OpEntryPoint` **must** not have more than one variable decorated
  with `InputAttachmentIndex` per image aspect of the attachment image
  bound to it, either explicitly or implicitly as described by <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-inputattachment"
  target="_blank" rel="noopener">input attachment interface</a>

- <a href="#VUID-RuntimeSpirv-shaderTileImageColorReadAccess-08728"
  id="VUID-RuntimeSpirv-shaderTileImageColorReadAccess-08728"></a>
  VUID-RuntimeSpirv-shaderTileImageColorReadAccess-08728  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderTileImageColorReadAccess"
  target="_blank"
  rel="noopener"><code>shaderTileImageColorReadAccess</code></a> is not
  enabled, `OpColorAttachmentReadEXT` operation **must** not be used

- <a href="#VUID-RuntimeSpirv-shaderTileImageDepthReadAccess-08729"
  id="VUID-RuntimeSpirv-shaderTileImageDepthReadAccess-08729"></a>
  VUID-RuntimeSpirv-shaderTileImageDepthReadAccess-08729  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderTileImageDepthReadAccess"
  target="_blank"
  rel="noopener"><code>shaderTileImageDepthReadAccess</code></a> is not
  enabled, `OpDepthAttachmentReadEXT` operation **must** not be used

- <a href="#VUID-RuntimeSpirv-shaderTileImageStencilReadAccess-08730"
  id="VUID-RuntimeSpirv-shaderTileImageStencilReadAccess-08730"></a>
  VUID-RuntimeSpirv-shaderTileImageStencilReadAccess-08730  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderTileImageStencilReadAccess"
  target="_blank"
  rel="noopener"><code>shaderTileImageStencilReadAccess</code></a> is
  not enabled, `OpStencilAttachmentReadEXT` operation **must** not be
  used

- <a href="#VUID-RuntimeSpirv-minSampleShading-08731"
  id="VUID-RuntimeSpirv-minSampleShading-08731"></a>
  VUID-RuntimeSpirv-minSampleShading-08731  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-sampleshading"
  target="_blank" rel="noopener">sample shading</a> is enabled and
  `minSampleShading` is 1.0, the `sample` operand of any
  `OpColorAttachmentReadEXT`, `OpDepthAttachmentReadEXT`, or
  `OpStencilAttachmentReadEXT` operation **must** evaluate to the value
  of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-multisampling-coverage-mask"
  target="_blank" rel="noopener">coverage index</a> for any given
  fragment invocation

- <a href="#VUID-RuntimeSpirv-minSampleShading-08732"
  id="VUID-RuntimeSpirv-minSampleShading-08732"></a>
  VUID-RuntimeSpirv-minSampleShading-08732  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-sampleshading"
  target="_blank" rel="noopener">sample shading</a> is enabled and any
  of the `OpColorAttachmentReadEXT`, `OpDepthAttachmentReadEXT`, or
  `OpStencilAttachmentReadEXT` operations are used, then
  `minSampleShading` **must** be 1.0

- <a href="#VUID-RuntimeSpirv-MeshEXT-09218"
  id="VUID-RuntimeSpirv-MeshEXT-09218"></a>
  VUID-RuntimeSpirv-MeshEXT-09218  
  In mesh shaders using the `MeshEXT` or `MeshNV` `Execution` `Model`
  and the `OutputPoints` `Execution` `Mode`, if <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance5"
  target="_blank" rel="noopener"><code>maintenance5</code></a> is not
  enabled, and if the number of output points is greater than 0, a
  `PointSize` decorated variable **must** be written to for each output
  point

- <a href="#VUID-RuntimeSpirv-maintenance5-09190"
  id="VUID-RuntimeSpirv-maintenance5-09190"></a>
  VUID-RuntimeSpirv-maintenance5-09190  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance5"
  target="_blank" rel="noopener"><code>maintenance5</code></a> is
  enabled and a `PointSize` decorated variable is written to, all
  execution paths **must** write to a `PointSize` decorated variable

- <a href="#VUID-RuntimeSpirv-ShaderEnqueueAMDX-09191"
  id="VUID-RuntimeSpirv-ShaderEnqueueAMDX-09191"></a>
  VUID-RuntimeSpirv-ShaderEnqueueAMDX-09191  
  The `ShaderEnqueueAMDX` capability **must** only be used in shaders
  with the `GLCompute` execution model

- <a href="#VUID-RuntimeSpirv-NodePayloadAMDX-09192"
  id="VUID-RuntimeSpirv-NodePayloadAMDX-09192"></a>
  VUID-RuntimeSpirv-NodePayloadAMDX-09192  
  Variables in the `NodePayloadAMDX` storage class **must** only be
  declared in the `GLCompute` execution model

- <a href="#VUID-RuntimeSpirv-maxExecutionGraphShaderPayloadSize-09193"
  id="VUID-RuntimeSpirv-maxExecutionGraphShaderPayloadSize-09193"></a>
  VUID-RuntimeSpirv-maxExecutionGraphShaderPayloadSize-09193  
  Variables declared in the `NodePayloadAMDX` storage class **must** not
  be larger than the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxExecutionGraphShaderPayloadSize"
  target="_blank"
  rel="noopener"><code>maxExecutionGraphShaderPayloadSize</code></a>
  limit

- <a href="#VUID-RuntimeSpirv-maxExecutionGraphShaderPayloadSize-09194"
  id="VUID-RuntimeSpirv-maxExecutionGraphShaderPayloadSize-09194"></a>
  VUID-RuntimeSpirv-maxExecutionGraphShaderPayloadSize-09194  
  Variables declared in the `NodeOutputPayloadAMDX` storage class
  **must** not be larger than the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxExecutionGraphShaderPayloadSize"
  target="_blank"
  rel="noopener"><code>maxExecutionGraphShaderPayloadSize</code></a>
  limit

- <a href="#VUID-RuntimeSpirv-maxExecutionGraphShaderPayloadSize-09195"
  id="VUID-RuntimeSpirv-maxExecutionGraphShaderPayloadSize-09195"></a>
  VUID-RuntimeSpirv-maxExecutionGraphShaderPayloadSize-09195  
  For a given entry point, the sum of the size of any variable in the
  `NodePayloadAMDX` storage class, and the combined size of all
  statically initialized variables in the `NodeOutputPayloadAMDX`
  storage class **must** not be greater than <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxExecutionGraphShaderPayloadSize"
  target="_blank"
  rel="noopener"><code>maxExecutionGraphShaderPayloadSize</code></a>

- <a href="#VUID-RuntimeSpirv-maxExecutionGraphShaderPayloadCount-09196"
  id="VUID-RuntimeSpirv-maxExecutionGraphShaderPayloadCount-09196"></a>
  VUID-RuntimeSpirv-maxExecutionGraphShaderPayloadCount-09196  
  Shaders **must** not statically initialize more than <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxExecutionGraphShaderPayloadCount"
  target="_blank"
  rel="noopener"><code>maxExecutionGraphShaderPayloadCount</code></a>
  variables in the `NodeOutputPayloadAMDX` storage class

- <a href="#VUID-RuntimeSpirv-maxExecutionGraphShaderOutputNodes-09197"
  id="VUID-RuntimeSpirv-maxExecutionGraphShaderOutputNodes-09197"></a>
  VUID-RuntimeSpirv-maxExecutionGraphShaderOutputNodes-09197  
  Shaders **must** not include more than <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxExecutionGraphShaderOutputNodes"
  target="_blank"
  rel="noopener"><code>maxExecutionGraphShaderOutputNodes</code></a>
  instances of `OpInitializeNodePayloadsAMDX`

- <a href="#VUID-RuntimeSpirv-OpImageBlockMatchWindow-09219"
  id="VUID-RuntimeSpirv-OpImageBlockMatchWindow-09219"></a>
  VUID-RuntimeSpirv-OpImageBlockMatchWindow-09219  
  If a `OpImageBlockMatchWindow*QCOM` or `OpImageBlockMatchGather*QCOM`
  operation is used, then the `target` `sampled` `image`, `reference`
  `sampled` `image`, and `Block` `Size` parameters **must** both be
  *dynamically uniform* for the quad

- <a href="#VUID-RuntimeSpirv-OpImageBlockMatchWindow-09220"
  id="VUID-RuntimeSpirv-OpImageBlockMatchWindow-09220"></a>
  VUID-RuntimeSpirv-OpImageBlockMatchWindow-09220  
  If a `OpImageBlockMatchWindow*QCOM` or `OpImageBlockMatchGather*QCOM`
  operation is used, then `target` `sampled` `image` and `reference`
  `sampled` `image` parameters **must** be of storage class
  `UniformConstant` and type `OpTypeImage` with `Depth`=0, `Dim`=`2D`,
  `Arrayed`=0, `MS`=0, and `Sampled`=1

- <a href="#VUID-RuntimeSpirv-OpImageBlockMatchWindow-09221"
  id="VUID-RuntimeSpirv-OpImageBlockMatchWindow-09221"></a>
  VUID-RuntimeSpirv-OpImageBlockMatchWindow-09221  
  If a `OpImageBlockMatchWindow*QCOM` or `OpImageBlockMatchGather*QCOM`
  operation is used, then the `target` `sampled` `image` and `reference`
  `sampled` `image` parameters **must** be decorated with
  `BlockMatchTextureQCOM`

- <a href="#VUID-RuntimeSpirv-OpImageBlockMatchWindow-09222"
  id="VUID-RuntimeSpirv-OpImageBlockMatchWindow-09222"></a>
  VUID-RuntimeSpirv-OpImageBlockMatchWindow-09222  
  If a `OpImageBlockMatchWindow*QCOM` or `OpImageBlockMatchGather*QCOM`
  operation is used, then `target` `sampled` `image` and `reference`
  `sampled` `image` parameters **must** have been created using an
  identical sampler object

- <a href="#VUID-RuntimeSpirv-OpImageBlockMatchWindow-09223"
  id="VUID-RuntimeSpirv-OpImageBlockMatchWindow-09223"></a>
  VUID-RuntimeSpirv-OpImageBlockMatchWindow-09223  
  If a `OpImageBlockMatchWindow*QCOM` or `OpImageBlockMatchGather*QCOM`
  operation is used, then `target` `sampled` `image` and `reference`
  `sampled` `image` parameters **must** have been created with a sampler
  object with `unnormalizedCoordinates` equal to `VK_TRUE`

- <a href="#VUID-RuntimeSpirv-OpImageBlockMatchWindow-09224"
  id="VUID-RuntimeSpirv-OpImageBlockMatchWindow-09224"></a>
  VUID-RuntimeSpirv-OpImageBlockMatchWindow-09224  
  If a `OpImageBlockMatchWindow*QCOM` or `OpImageBlockMatchGather*QCOM`
  operation is used, then `target` `sampled` `image` and `reference`
  `sampled` `image` parameters **must** have been created with sampler
  object with `unnormalizedCoordinates` equal to `VK_TRUE`

- <a href="#VUID-RuntimeSpirv-maxBlockMatchRegion-09225"
  id="VUID-RuntimeSpirv-maxBlockMatchRegion-09225"></a>
  VUID-RuntimeSpirv-maxBlockMatchRegion-09225  
  If a `OpImageBlockMatchWindow*QCOM` or `OpImageBlockMatchGather*QCOM`
  operation is used, then `Block` `Size` less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-blockmatch-maxblocksize"
  target="_blank" rel="noopener"><code>maxBlockMatchRegion</code></a>

- <a href="#VUID-RuntimeSpirv-pNext-09226"
  id="VUID-RuntimeSpirv-pNext-09226"></a>
  VUID-RuntimeSpirv-pNext-09226  
  If a `OpImageBlockMatchWindow*QCOM` operation is used, then `target`
  `sampled` `image` **must** have been created using asampler object
  that included
  [VkSamplerBlockMatchWindowCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerBlockMatchWindowCreateInfoQCOM.html)
  in the `pNext` chain

- <a href="#VUID-RuntimeSpirv-MaximallyReconvergesKHR-09565"
  id="VUID-RuntimeSpirv-MaximallyReconvergesKHR-09565"></a>
  VUID-RuntimeSpirv-MaximallyReconvergesKHR-09565  
  The execution mode `MaximallyReconvergesKHR` **must** not be applied
  to an entry point unless the entry point does not execute any <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ray-tracing-repack"
  target="_blank" rel="noopener"><em>invocation repack
  instructions</em></a>

- <a href="#VUID-RuntimeSpirv-shaderSubgroupRotateClustered-09566"
  id="VUID-RuntimeSpirv-shaderSubgroupRotateClustered-09566"></a>
  VUID-RuntimeSpirv-shaderSubgroupRotateClustered-09566  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderSubgroupRotateClustered"
  target="_blank"
  rel="noopener"><code>shaderSubgroupRotateClustered</code></a> is
  `VK_FALSE`, then the `ClusterSize` operand to
  `OpGroupNonUniformRotateKHR` **must** not be used

- <a href="#VUID-RuntimeSpirv-protectedNoFault-09645"
  id="VUID-RuntimeSpirv-protectedNoFault-09645"></a>
  VUID-RuntimeSpirv-protectedNoFault-09645  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-protectedNoFault"
  target="_blank" rel="noopener"><code>protectedNoFault</code></a> is
  not supported, the `Storage` `Class` of the `PhysicalStorageBuffer`
  **must** not be used if the buffer being accessed is <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-protected-memory"
  target="_blank" rel="noopener">protected</a>

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#RuntimeSpirv"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
