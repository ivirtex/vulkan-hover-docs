# VkBindDescriptorBufferEmbeddedSamplersInfoEXT(3) Manual Page

## Name

VkBindDescriptorBufferEmbeddedSamplersInfoEXT - Structure specifying
embedded immutable sampler offsets to set in a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBindDescriptorBufferEmbeddedSamplersInfoEXT` structure is defined
as:

``` c
// Provided by VK_KHR_maintenance6 with VK_EXT_descriptor_buffer
typedef struct VkBindDescriptorBufferEmbeddedSamplersInfoEXT {
    VkStructureType       sType;
    const void*           pNext;
    VkShaderStageFlags    stageFlags;
    VkPipelineLayout      layout;
    uint32_t              set;
} VkBindDescriptorBufferEmbeddedSamplersInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `stageFlags` is a bitmask of
  [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html) specifying the
  shader stages that will use the embedded immutable samplers.

- `layout` is a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) object used to
  program the bindings. If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-dynamicPipelineLayout"
  target="_blank" rel="noopener"><code>dynamicPipelineLayout</code></a>
  feature is enabled, `layout` **can** be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and the layout **must** be
  specified by chaining
  [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html)
  structure off the `pNext`

- `set` is the number of the set to be bound.

## <a href="#_description" class="anchor"></a>Description

If `stageFlags` specifies a subset of all stages corresponding to one or
more pipeline bind points, the binding operation still affects all
stages corresponding to the given pipeline bind point(s) as if the
equivalent original version of this command had been called with the
same parameters. For example, specifying a `stageFlags` value of
`VK_SHADER_STAGE_VERTEX_BIT` \| `VK_SHADER_STAGE_FRAGMENT_BIT` \|
`VK_SHADER_STAGE_COMPUTE_BIT` is equivalent to calling the original
version of this command once with `VK_PIPELINE_BIND_POINT_GRAPHICS` and
once with `VK_PIPELINE_BIND_POINT_COMPUTE`.

Valid Usage

- <a href="#VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-set-08070"
  id="VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-set-08070"></a>
  VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-set-08070  
  The [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html) at index `set`
  when `layout` was created **must** have been created with the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_EMBEDDED_IMMUTABLE_SAMPLERS_BIT_EXT`
  bit set

- <a href="#VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-set-08071"
  id="VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-set-08071"></a>
  VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-set-08071  
  `set` **must** be less than or equal to
  [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html)::`setLayoutCount`
  provided when `layout` was created

<!-- -->

- <a href="#VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-None-09495"
  id="VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-None-09495"></a>
  VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-None-09495  
  If the [`dynamicPipelineLayout`](#features-dynamicPipelineLayout)
  feature is not enabled, `layout` **must** be a valid
  [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) handle

- <a
  href="#VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-layout-09496"
  id="VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-layout-09496"></a>
  VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-layout-09496  
  If `layout` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the `pNext`
  chain **must** include a valid
  [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html)
  structure

Valid Usage (Implicit)

- <a
  href="#VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-sType-sType"
  id="VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-sType-sType"></a>
  VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_BIND_DESCRIPTOR_BUFFER_EMBEDDED_SAMPLERS_INFO_EXT`

- <a
  href="#VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-pNext-pNext"
  id="VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-pNext-pNext"></a>
  VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html)

- <a
  href="#VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-sType-unique"
  id="VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-sType-unique"></a>
  VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a
  href="#VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-stageFlags-parameter"
  id="VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-stageFlags-parameter"></a>
  VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-stageFlags-parameter  
  `stageFlags` **must** be a valid combination of
  [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlagBits.html) values

- <a
  href="#VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-stageFlags-requiredbitmask"
  id="VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-stageFlags-requiredbitmask"></a>
  VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-stageFlags-requiredbitmask  
  `stageFlags` **must** not be `0`

- <a
  href="#VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-layout-parameter"
  id="VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-layout-parameter"></a>
  VUID-VkBindDescriptorBufferEmbeddedSamplersInfoEXT-layout-parameter  
  If `layout` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `layout`
  **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_descriptor_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_buffer.html),
[VK_KHR_maintenance6](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance6.html),
[VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html),
[VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdBindDescriptorBufferEmbeddedSamplers2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindDescriptorBufferEmbeddedSamplers2EXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBindDescriptorBufferEmbeddedSamplersInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
