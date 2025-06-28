# VkDescriptorSetLayoutBindingFlagsCreateInfo(3) Manual Page

## Name

VkDescriptorSetLayoutBindingFlagsCreateInfo - Structure specifying creation flags for descriptor set layout bindings



## [](#_c_specification)C Specification

If the `pNext` chain of a [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateInfo.html) structure includes a [VkDescriptorSetLayoutBindingFlagsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutBindingFlagsCreateInfo.html) structure, then that structure includes an array of flags, one for each descriptor set layout binding.

The [VkDescriptorSetLayoutBindingFlagsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutBindingFlagsCreateInfo.html) structure is defined as:

```c++
// Provided by VK_VERSION_1_2
typedef struct VkDescriptorSetLayoutBindingFlagsCreateInfo {
    VkStructureType                    sType;
    const void*                        pNext;
    uint32_t                           bindingCount;
    const VkDescriptorBindingFlags*    pBindingFlags;
} VkDescriptorSetLayoutBindingFlagsCreateInfo;
```

or the equivalent

```c++
// Provided by VK_EXT_descriptor_indexing
typedef VkDescriptorSetLayoutBindingFlagsCreateInfo VkDescriptorSetLayoutBindingFlagsCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `bindingCount` is zero or the number of elements in `pBindingFlags`.
- `pBindingFlags` is a pointer to an array of [VkDescriptorBindingFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBindingFlags.html) bitfields, one for each descriptor set layout binding.

## [](#_description)Description

If `bindingCount` is zero or if this structure is not included in the `pNext` chain, the [VkDescriptorBindingFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBindingFlags.html) for each descriptor set layout binding is considered to be zero. Otherwise, the descriptor set layout binding at [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateInfo.html)::`pBindings`\[i] uses the flags in `pBindingFlags`\[i].

Valid Usage

- [](#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-bindingCount-03002)VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-bindingCount-03002  
  If `bindingCount` is not zero, `bindingCount` **must** equal [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateInfo.html)::`bindingCount`
- [](#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-flags-03003)VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-flags-03003  
  If [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateInfo.html)::`flags` includes `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT`, then all elements of `pBindingFlags` **must** not include `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`, `VK_DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT`, or `VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT`
- [](#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-pBindingFlags-03004)VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-pBindingFlags-03004  
  If an element of `pBindingFlags` includes `VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT`, then it **must** be the element with the highest `binding` number
- [](#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingUniformBufferUpdateAfterBind-03005)VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingUniformBufferUpdateAfterBind-03005  
  If [VkPhysicalDeviceDescriptorIndexingFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html)::`descriptorBindingUniformBufferUpdateAfterBind` is not enabled, all bindings with descriptor type `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER` **must** not use `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`
- [](#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingSampledImageUpdateAfterBind-03006)VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingSampledImageUpdateAfterBind-03006  
  If [VkPhysicalDeviceDescriptorIndexingFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html)::`descriptorBindingSampledImageUpdateAfterBind` is not enabled, all bindings with descriptor type `VK_DESCRIPTOR_TYPE_SAMPLER`, `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, or `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE` **must** not use `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`
- [](#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingStorageImageUpdateAfterBind-03007)VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingStorageImageUpdateAfterBind-03007  
  If [VkPhysicalDeviceDescriptorIndexingFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html)::`descriptorBindingStorageImageUpdateAfterBind` is not enabled, all bindings with descriptor type `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE` **must** not use `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`
- [](#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingStorageBufferUpdateAfterBind-03008)VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingStorageBufferUpdateAfterBind-03008  
  If [VkPhysicalDeviceDescriptorIndexingFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html)::`descriptorBindingStorageBufferUpdateAfterBind` is not enabled, all bindings with descriptor type `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER` **must** not use `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`
- [](#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingUniformTexelBufferUpdateAfterBind-03009)VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingUniformTexelBufferUpdateAfterBind-03009  
  If [VkPhysicalDeviceDescriptorIndexingFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html)::`descriptorBindingUniformTexelBufferUpdateAfterBind` is not enabled, all bindings with descriptor type `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` **must** not use `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`
- [](#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingStorageTexelBufferUpdateAfterBind-03010)VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingStorageTexelBufferUpdateAfterBind-03010  
  If [VkPhysicalDeviceDescriptorIndexingFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html)::`descriptorBindingStorageTexelBufferUpdateAfterBind` is not enabled, all bindings with descriptor type `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` **must** not use `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`
- [](#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingInlineUniformBlockUpdateAfterBind-02211)VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingInlineUniformBlockUpdateAfterBind-02211  
  If [VkPhysicalDeviceInlineUniformBlockFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceInlineUniformBlockFeatures.html)::`descriptorBindingInlineUniformBlockUpdateAfterBind` is not enabled, all bindings with descriptor type `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` **must** not use `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`
- [](#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingAccelerationStructureUpdateAfterBind-03570)VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingAccelerationStructureUpdateAfterBind-03570  
  If [VkPhysicalDeviceAccelerationStructureFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceAccelerationStructureFeaturesKHR.html)::`descriptorBindingAccelerationStructureUpdateAfterBind` is not enabled, all bindings with descriptor type `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR` or `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_NV` **must** not use `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`
- [](#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-None-03011)VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-None-03011  
  All bindings with descriptor type `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`, `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`, or `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` **must** not use `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`
- [](#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingUpdateUnusedWhilePending-03012)VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingUpdateUnusedWhilePending-03012  
  If [VkPhysicalDeviceDescriptorIndexingFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html)::`descriptorBindingUpdateUnusedWhilePending` is not enabled, all elements of `pBindingFlags` **must** not include `VK_DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT`
- [](#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingPartiallyBound-03013)VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingPartiallyBound-03013  
  If [VkPhysicalDeviceDescriptorIndexingFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html)::`descriptorBindingPartiallyBound` is not enabled, all elements of `pBindingFlags` **must** not include `VK_DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT`
- [](#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingVariableDescriptorCount-03014)VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingVariableDescriptorCount-03014  
  If [VkPhysicalDeviceDescriptorIndexingFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html)::`descriptorBindingVariableDescriptorCount` is not enabled, all elements of `pBindingFlags` **must** not include `VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT`
- [](#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-pBindingFlags-03015)VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-pBindingFlags-03015  
  If an element of `pBindingFlags` includes `VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT`, that element’s `descriptorType` **must** not be `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` or `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`
- [](#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingStorageTensorUpdateAfterBind-09697)VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingStorageTensorUpdateAfterBind-09697  
  If [VkPhysicalDeviceTensorFeaturesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTensorFeaturesARM.html)::`descriptorBindingStorageTensorUpdateAfterBind` is not enabled, all bindings with descriptor type `VK_DESCRIPTOR_TYPE_TENSOR_ARM` **must** not use `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`

Valid Usage (Implicit)

- [](#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-sType-sType)VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_BINDING_FLAGS_CREATE_INFO`
- [](#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-pBindingFlags-parameter)VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-pBindingFlags-parameter  
  If `bindingCount` is not `0`, `pBindingFlags` **must** be a valid pointer to an array of `bindingCount` valid combinations of [VkDescriptorBindingFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBindingFlagBits.html) values

## [](#_see_also)See Also

[VK\_EXT\_descriptor\_indexing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_indexing.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkDescriptorBindingFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBindingFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDescriptorSetLayoutBindingFlagsCreateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0