# VkPhysicalDeviceSubgroupProperties(3) Manual Page

## Name

VkPhysicalDeviceSubgroupProperties - Structure describing subgroup
support for an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceSubgroupProperties` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkPhysicalDeviceSubgroupProperties {
    VkStructureType           sType;
    void*                     pNext;
    uint32_t                  subgroupSize;
    VkShaderStageFlags        supportedStages;
    VkSubgroupFeatureFlags    supportedOperations;
    VkBool32                  quadOperationsInAllStages;
} VkPhysicalDeviceSubgroupProperties;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-limits-subgroupSize"></span> `subgroupSize` is the
  default number of invocations in each subgroup. `subgroupSize` is at
  least 1 if any of the physical device’s queues support
  `VK_QUEUE_GRAPHICS_BIT` or `VK_QUEUE_COMPUTE_BIT`. `subgroupSize` is a
  power-of-two.

- <span id="extension-limits-subgroupSupportedStages"></span>
  `supportedStages` is a bitfield of
  [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html) describing the
  shader stages that <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-group-operations"
  target="_blank" rel="noopener">group operations</a> with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-scope-subgroup"
  target="_blank" rel="noopener">subgroup scope</a> are supported in.
  `supportedStages` will have the `VK_SHADER_STAGE_COMPUTE_BIT` bit set
  if any of the physical device’s queues support `VK_QUEUE_COMPUTE_BIT`.

- <span id="extension-limits-subgroupSupportedOperations"></span>
  `supportedOperations` is a bitmask of
  [VkSubgroupFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubgroupFeatureFlagBits.html) specifying
  the sets of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-group-operations"
  target="_blank" rel="noopener">group operations</a> with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-scope-subgroup"
  target="_blank" rel="noopener">subgroup scope</a> supported on this
  device. `supportedOperations` will have the
  `VK_SUBGROUP_FEATURE_BASIC_BIT` bit set if any of the physical
  device’s queues support `VK_QUEUE_GRAPHICS_BIT` or
  `VK_QUEUE_COMPUTE_BIT`.

- <span id="extension-limits-subgroupQuadOperationsInAllStages"></span>
  `quadOperationsInAllStages` is a boolean specifying whether <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-quad-operations"
  target="_blank" rel="noopener">quad group operations</a> are available
  in all stages, or are restricted to fragment and compute stages.

If the `VkPhysicalDeviceSubgroupProperties` structure is included in the
`pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

If `supportedOperations` includes <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-subgroup-quad"
target="_blank"
rel="noopener"><code>VK_SUBGROUP_FEATURE_QUAD_BIT</code></a>, or <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderSubgroupUniformControlFlow"
target="_blank"
rel="noopener"><code>shaderSubgroupUniformControlFlow</code></a> is
enabled, `subgroupSize` **must** be greater than or equal to 4.

If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderQuadControl"
target="_blank" rel="noopener"><code>shaderQuadControl</code></a>
feature is supported, `supportedOperations` **must** include <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-subgroup-quad"
target="_blank"
rel="noopener"><code>VK_SUBGROUP_FEATURE_QUAD_BIT</code></a>.

If [VK_KHR_shader_subgroup_rotate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shader_subgroup_rotate.html)
is supported, and the implementation advertises support with a
[VkExtensionProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtensionProperties.html)::`specVersion`
greater than or equal to 2, and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderSubgroupRotate"
target="_blank" rel="noopener"><code>shaderSubgroupRotate</code></a> is
supported, `VK_SUBGROUP_FEATURE_ROTATE_BIT_KHR` **must** be returned in
`supportedOperations`. If
[VK_KHR_shader_subgroup_rotate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shader_subgroup_rotate.html) is
supported, and the implementation advertises support with a
[VkExtensionProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtensionProperties.html)::`specVersion`
greater than or equal to 2, and <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderSubgroupRotateClustered"
target="_blank"
rel="noopener"><code>shaderSubgroupRotateClustered</code></a> is
supported, `VK_SUBGROUP_FEATURE_ROTATE_CLUSTERED_BIT_KHR` **must** be
returned in `supportedOperations`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p><code>VK_SUBGROUP_FEATURE_ROTATE_BIT_KHR</code> and
<code>VK_SUBGROUP_FEATURE_ROTATE_CLUSTERED_BIT_KHR</code> were added in
version 2 of the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shader_subgroup_rotate.html">VK_KHR_shader_subgroup_rotate</a>
extension, after the initial release, so there are implementations that
do not advertise these bits. Applications should use the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderSubgroupRotate"
target="_blank" rel="noopener"><code>shaderSubgroupRotate</code></a> and
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shaderSubgroupRotateClustered"
target="_blank"
rel="noopener"><code>shaderSubgroupRotateClustered</code></a> features
to determine and enable support. These bits are advertised here for
consistency and for future dependencies.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceSubgroupProperties-sType-sType"
  id="VUID-VkPhysicalDeviceSubgroupProperties-sType-sType"></a>
  VUID-VkPhysicalDeviceSubgroupProperties-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_PROPERTIES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSubgroupFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubgroupFeatureFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceSubgroupProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
