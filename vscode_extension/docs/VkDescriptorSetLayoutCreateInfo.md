# VkDescriptorSetLayoutCreateInfo(3) Manual Page

## Name

VkDescriptorSetLayoutCreateInfo - Structure specifying parameters of a newly created descriptor set layout



## [](#_c_specification)C Specification

Information about the descriptor set layout is passed in a `VkDescriptorSetLayoutCreateInfo` structure:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkDescriptorSetLayoutCreateInfo {
    VkStructureType                        sType;
    const void*                            pNext;
    VkDescriptorSetLayoutCreateFlags       flags;
    uint32_t                               bindingCount;
    const VkDescriptorSetLayoutBinding*    pBindings;
} VkDescriptorSetLayoutCreateInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkDescriptorSetLayoutCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateFlagBits.html) specifying options for descriptor set layout creation.
- `bindingCount` is the number of elements in `pBindings`.
- `pBindings` is a pointer to an array of [VkDescriptorSetLayoutBinding](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutBinding.html) structures.

## [](#_description)Description

Valid Usage

- [](#VUID-VkDescriptorSetLayoutCreateInfo-binding-00279)VUID-VkDescriptorSetLayoutCreateInfo-binding-00279  
  If the [`perStageDescriptorSet`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-perStageDescriptorSet) feature is not enabled, or `flags` does not contain `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PER_STAGE_BIT_NV`, then the [VkDescriptorSetLayoutBinding](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutBinding.html)::`binding` members of the elements of the `pBindings` array **must** each have different values
- [](#VUID-VkDescriptorSetLayoutCreateInfo-flags-10354)VUID-VkDescriptorSetLayoutCreateInfo-flags-10354  
  If `flags` contains `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT`, and the [VK\_KHR\_push\_descriptor](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_push_descriptor.html) extension is not enabled, [`pushDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-pushDescriptor) **must** be enabled
- [](#VUID-VkDescriptorSetLayoutCreateInfo-flags-00280)VUID-VkDescriptorSetLayoutCreateInfo-flags-00280  
  If `flags` contains `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT`, then all elements of `pBindings` **must** not have a `descriptorType` of `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` or `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`
- [](#VUID-VkDescriptorSetLayoutCreateInfo-flags-02208)VUID-VkDescriptorSetLayoutCreateInfo-flags-02208  
  If `flags` contains `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT`, then all elements of `pBindings` **must** not have a `descriptorType` of `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`
- [](#VUID-VkDescriptorSetLayoutCreateInfo-flags-00281)VUID-VkDescriptorSetLayoutCreateInfo-flags-00281  
  If `flags` contains `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT`, then the total number of elements of all bindings **must** be less than or equal to [VkPhysicalDevicePushDescriptorProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePushDescriptorProperties.html)::`maxPushDescriptors`
- [](#VUID-VkDescriptorSetLayoutCreateInfo-flags-04590)VUID-VkDescriptorSetLayoutCreateInfo-flags-04590  
  If `flags` contains `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT`, `flags` **must** not contain `VK_DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_EXT`
- [](#VUID-VkDescriptorSetLayoutCreateInfo-flags-04591)VUID-VkDescriptorSetLayoutCreateInfo-flags-04591  
  If `flags` contains `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT`, `pBindings` **must** not have a `descriptorType` of `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`
- [](#VUID-VkDescriptorSetLayoutCreateInfo-flags-03000)VUID-VkDescriptorSetLayoutCreateInfo-flags-03000  
  If any binding has the `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` bit set, `flags` **must** include `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT`
- [](#VUID-VkDescriptorSetLayoutCreateInfo-descriptorType-03001)VUID-VkDescriptorSetLayoutCreateInfo-descriptorType-03001  
  If any binding has the `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` bit set, then all bindings **must** not have `descriptorType` of `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` or `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`
- [](#VUID-VkDescriptorSetLayoutCreateInfo-flags-04592)VUID-VkDescriptorSetLayoutCreateInfo-flags-04592  
  If `flags` contains `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT`, `flags` **must** not contain `VK_DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_EXT`
- [](#VUID-VkDescriptorSetLayoutCreateInfo-pBindings-07303)VUID-VkDescriptorSetLayoutCreateInfo-pBindings-07303  
  If any element `pBindings`\[i] has a `descriptorType` of `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`, then the `pNext` chain **must** include a [VkMutableDescriptorTypeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMutableDescriptorTypeCreateInfoEXT.html) structure, and `mutableDescriptorTypeListCount` **must** be greater than i
- [](#VUID-VkDescriptorSetLayoutCreateInfo-descriptorType-04594)VUID-VkDescriptorSetLayoutCreateInfo-descriptorType-04594  
  If a binding has a `descriptorType` value of `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`, then `pImmutableSamplers` **must** be `NULL`
- [](#VUID-VkDescriptorSetLayoutCreateInfo-mutableDescriptorType-04595)VUID-VkDescriptorSetLayoutCreateInfo-mutableDescriptorType-04595  
  If [VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT.html)::`mutableDescriptorType` is not enabled, `pBindings` **must** not contain a `descriptorType` of `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`
- [](#VUID-VkDescriptorSetLayoutCreateInfo-flags-04596)VUID-VkDescriptorSetLayoutCreateInfo-flags-04596  
  If `flags` contains `VK_DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_EXT`, [VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT.html)::`mutableDescriptorType` **must** be enabled
- [](#VUID-VkDescriptorSetLayoutCreateInfo-flags-08000)VUID-VkDescriptorSetLayoutCreateInfo-flags-08000  
  If `flags` contains `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`, then all elements of `pBindings` **must** not have a `descriptorType` of `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` or `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`
- [](#VUID-VkDescriptorSetLayoutCreateInfo-flags-08001)VUID-VkDescriptorSetLayoutCreateInfo-flags-08001  
  If `flags` contains `VK_DESCRIPTOR_SET_LAYOUT_CREATE_EMBEDDED_IMMUTABLE_SAMPLERS_BIT_EXT`, `flags` **must** also contain `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`
- [](#VUID-VkDescriptorSetLayoutCreateInfo-flags-08002)VUID-VkDescriptorSetLayoutCreateInfo-flags-08002  
  If `flags` contains `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`, then `flags` **must** not contain `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT`
- [](#VUID-VkDescriptorSetLayoutCreateInfo-flags-08003)VUID-VkDescriptorSetLayoutCreateInfo-flags-08003  
  If `flags` contains `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`, then `flags` **must** not contain `VK_DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_EXT`
- [](#VUID-VkDescriptorSetLayoutCreateInfo-flags-09463)VUID-VkDescriptorSetLayoutCreateInfo-flags-09463  
  If `flags` contains `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PER_STAGE_BIT_NV`, then [`perStageDescriptorSet`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-perStageDescriptorSet) **must** be enabled
- [](#VUID-VkDescriptorSetLayoutCreateInfo-flags-09464)VUID-VkDescriptorSetLayoutCreateInfo-flags-09464  
  If `flags` contains `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PER_STAGE_BIT_NV`, then there **must** not be any two elements of the `pBindings` array with the same [VkDescriptorSetLayoutBinding](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutBinding.html)::`binding` value and their [VkDescriptorSetLayoutBinding](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutBinding.html)::`stageFlags` containing the same bit

Valid Usage (Implicit)

- [](#VUID-VkDescriptorSetLayoutCreateInfo-sType-sType)VUID-VkDescriptorSetLayoutCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO`
- [](#VUID-VkDescriptorSetLayoutCreateInfo-pNext-pNext)VUID-VkDescriptorSetLayoutCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkDescriptorSetLayoutBindingFlagsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutBindingFlagsCreateInfo.html) or [VkMutableDescriptorTypeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMutableDescriptorTypeCreateInfoEXT.html)
- [](#VUID-VkDescriptorSetLayoutCreateInfo-sType-unique)VUID-VkDescriptorSetLayoutCreateInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkDescriptorSetLayoutCreateInfo-flags-parameter)VUID-VkDescriptorSetLayoutCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of [VkDescriptorSetLayoutCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateFlagBits.html) values
- [](#VUID-VkDescriptorSetLayoutCreateInfo-pBindings-parameter)VUID-VkDescriptorSetLayoutCreateInfo-pBindings-parameter  
  If `bindingCount` is not `0`, `pBindings` **must** be a valid pointer to an array of `bindingCount` valid [VkDescriptorSetLayoutBinding](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutBinding.html) structures

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDescriptorSetLayoutBinding](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutBinding.html), [VkDescriptorSetLayoutCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDescriptorSetLayout.html), [vkGetDescriptorSetLayoutSupport](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDescriptorSetLayoutSupport.html), [vkGetDescriptorSetLayoutSupportKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDescriptorSetLayoutSupportKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDescriptorSetLayoutCreateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0