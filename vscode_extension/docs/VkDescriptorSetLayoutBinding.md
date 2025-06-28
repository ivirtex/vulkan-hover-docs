# VkDescriptorSetLayoutBinding(3) Manual Page

## Name

VkDescriptorSetLayoutBinding - Structure specifying a descriptor set layout binding



## [](#_c_specification)C Specification

The `VkDescriptorSetLayoutBinding` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkDescriptorSetLayoutBinding {
    uint32_t              binding;
    VkDescriptorType      descriptorType;
    uint32_t              descriptorCount;
    VkShaderStageFlags    stageFlags;
    const VkSampler*      pImmutableSamplers;
} VkDescriptorSetLayoutBinding;
```

## [](#_members)Members

- `binding` is the binding number of this entry and corresponds to a resource of the same binding number in the shader stages.
- `descriptorType` is a [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html) specifying which type of resource descriptors are used for this binding.
- `descriptorCount` is the number of descriptors contained in the binding, accessed in a shader as an array, except if `descriptorType` is `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` in which case `descriptorCount` is the size in bytes of the inline uniform block. If `descriptorCount` is zero this binding entry is reserved and the resource **must** not be accessed from any stage via this binding within any pipeline using the set layout.
- `stageFlags` member is a bitmask of [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) specifying which pipeline shader stages **can** access a resource for this binding. `VK_SHADER_STAGE_ALL` is a shorthand specifying that all defined shader stages, including any additional stages defined by extensions, **can** access the resource.
  
  If a shader stage is not included in `stageFlags`, then a resource **must** not be accessed from that stage via this binding within any pipeline using the set layout. Other than input attachments which are limited to the fragment shader, there are no limitations on what combinations of stages **can** use a descriptor binding, and in particular a binding **can** be used by both graphics stages and the compute stage.
- `pImmutableSamplers` affects initialization of samplers. If `descriptorType` specifies a `VK_DESCRIPTOR_TYPE_SAMPLER` or `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` type descriptor, then `pImmutableSamplers` **can** be used to initialize a set of *immutable samplers*. Immutable samplers are permanently bound into the set layout and **must** not be changed; updating a `VK_DESCRIPTOR_TYPE_SAMPLER` descriptor with immutable samplers is not allowed and updates to a `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` descriptor with immutable samplers does not modify the samplers (the image views are updated, but the sampler updates are ignored). If `pImmutableSamplers` is not `NULL`, then it is a pointer to an array of sampler handles that will be copied into the set layout and used for the corresponding binding. Only the sampler handles are copied; the sampler objects **must** not be destroyed before the final use of the set layout and any descriptor pools and sets created using it. If `pImmutableSamplers` is `NULL`, then the sampler slots are dynamic and sampler handles **must** be bound into descriptor sets using this layout. If `descriptorType` is not one of these descriptor types, then `pImmutableSamplers` is ignored.

## [](#_description)Description

The above layout definition allows the descriptor bindings to be specified sparsely such that not all binding numbers between 0 and the maximum binding number need to be specified in the `pBindings` array. Bindings that are not specified have a `descriptorCount` and `stageFlags` of zero, and the value of `descriptorType` is undefined. However, all binding numbers between 0 and the maximum binding number in the [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateInfo.html)::`pBindings` array **may** consume memory in the descriptor set layout even if not all descriptor bindings are used, though it **should** not consume additional memory from the descriptor pool.

Note

The maximum binding number specified **should** be as compact as possible to avoid wasted memory.

Valid Usage

- [](#VUID-VkDescriptorSetLayoutBinding-descriptorType-00282)VUID-VkDescriptorSetLayoutBinding-descriptorType-00282  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_SAMPLER` or `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, and `descriptorCount` is not `0` and `pImmutableSamplers` is not `NULL`, `pImmutableSamplers` **must** be a valid pointer to an array of `descriptorCount` valid `VkSampler` handles
- [](#VUID-VkDescriptorSetLayoutBinding-descriptorType-04604)VUID-VkDescriptorSetLayoutBinding-descriptorType-04604  
  If the [`inlineUniformBlock`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-inlineUniformBlock) feature is not enabled, `descriptorType` **must** not be `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`
- [](#VUID-VkDescriptorSetLayoutBinding-descriptorType-02209)VUID-VkDescriptorSetLayoutBinding-descriptorType-02209  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` then `descriptorCount` **must** be a multiple of `4`
- [](#VUID-VkDescriptorSetLayoutBinding-descriptorType-08004)VUID-VkDescriptorSetLayoutBinding-descriptorType-08004  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` and [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateInfo.html)::`flags` does not contain `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT` then `descriptorCount` **must** be less than or equal to `VkPhysicalDeviceInlineUniformBlockProperties`::`maxInlineUniformBlockSize`
- [](#VUID-VkDescriptorSetLayoutBinding-flags-08005)VUID-VkDescriptorSetLayoutBinding-flags-08005  
  If [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateInfo.html)::`flags` contains `VK_DESCRIPTOR_SET_LAYOUT_CREATE_EMBEDDED_IMMUTABLE_SAMPLERS_BIT_EXT`, `descriptorType` **must** be `VK_DESCRIPTOR_TYPE_SAMPLER`
- [](#VUID-VkDescriptorSetLayoutBinding-flags-08006)VUID-VkDescriptorSetLayoutBinding-flags-08006  
  If [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateInfo.html)::`flags` contains `VK_DESCRIPTOR_SET_LAYOUT_CREATE_EMBEDDED_IMMUTABLE_SAMPLERS_BIT_EXT`, `descriptorCount` **must** less than or equal to `1`
- [](#VUID-VkDescriptorSetLayoutBinding-flags-08007)VUID-VkDescriptorSetLayoutBinding-flags-08007  
  If [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateInfo.html)::`flags` contains `VK_DESCRIPTOR_SET_LAYOUT_CREATE_EMBEDDED_IMMUTABLE_SAMPLERS_BIT_EXT`, and `descriptorCount` is equal to `1`, `pImmutableSamplers` **must** not be `NULL`
- [](#VUID-VkDescriptorSetLayoutBinding-descriptorCount-09465)VUID-VkDescriptorSetLayoutBinding-descriptorCount-09465  
  If `descriptorCount` is not `0`, `stageFlags` **must** be `VK_SHADER_STAGE_ALL` or a valid combination of other [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) values
- [](#VUID-VkDescriptorSetLayoutBinding-descriptorType-01510)VUID-VkDescriptorSetLayoutBinding-descriptorType-01510  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT` and `descriptorCount` is not `0`, then `stageFlags` **must** be `0` or `VK_SHADER_STAGE_FRAGMENT_BIT`
- [](#VUID-VkDescriptorSetLayoutBinding-pImmutableSamplers-04009)VUID-VkDescriptorSetLayoutBinding-pImmutableSamplers-04009  
  The sampler objects indicated by `pImmutableSamplers` **must** not have a `borderColor` with one of the values `VK_BORDER_COLOR_FLOAT_CUSTOM_EXT` or `VK_BORDER_COLOR_INT_CUSTOM_EXT`
- [](#VUID-VkDescriptorSetLayoutBinding-descriptorType-04605)VUID-VkDescriptorSetLayoutBinding-descriptorType-04605  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`, then `pImmutableSamplers` **must** be `NULL`
- [](#VUID-VkDescriptorSetLayoutBinding-flags-09466)VUID-VkDescriptorSetLayoutBinding-flags-09466  
  If [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateInfo.html)::`flags` contains `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PER_STAGE_BIT_NV`, and `descriptorCount` is not `0`, then `stageFlags` **must** be a valid combination of `VK_SHADER_STAGE_VERTEX_BIT`, `VK_SHADER_STAGE_TESSELLATION_CONTROL_BIT`, `VK_SHADER_STAGE_TESSELLATION_EVALUATION_BIT`, `VK_SHADER_STAGE_GEOMETRY_BIT`, `VK_SHADER_STAGE_FRAGMENT_BIT` and `VK_SHADER_STAGE_COMPUTE_BIT` values

Valid Usage (Implicit)

- [](#VUID-VkDescriptorSetLayoutBinding-descriptorType-parameter)VUID-VkDescriptorSetLayoutBinding-descriptorType-parameter  
  `descriptorType` **must** be a valid [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html) value

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateInfo.html), [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html), [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html), [VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDescriptorSetLayoutBinding)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0