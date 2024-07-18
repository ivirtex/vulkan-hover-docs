# VkDescriptorSetLayoutBindingFlagsCreateInfo(3) Manual Page

## Name

VkDescriptorSetLayoutBindingFlagsCreateInfo - Structure specifying
creation flags for descriptor set layout bindings



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain of a
[VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateInfo.html)
structure includes a
[VkDescriptorSetLayoutBindingFlagsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutBindingFlagsCreateInfo.html)
structure, then that structure includes an array of flags, one for each
descriptor set layout binding.

The
[VkDescriptorSetLayoutBindingFlagsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutBindingFlagsCreateInfo.html)
structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkDescriptorSetLayoutBindingFlagsCreateInfo {
    VkStructureType                    sType;
    const void*                        pNext;
    uint32_t                           bindingCount;
    const VkDescriptorBindingFlags*    pBindingFlags;
} VkDescriptorSetLayoutBindingFlagsCreateInfo;
```

or the equivalent

``` c
// Provided by VK_EXT_descriptor_indexing
typedef VkDescriptorSetLayoutBindingFlagsCreateInfo VkDescriptorSetLayoutBindingFlagsCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `bindingCount` is zero or the number of elements in `pBindingFlags`.

- `pBindingFlags` is a pointer to an array of
  [VkDescriptorBindingFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBindingFlags.html) bitfields,
  one for each descriptor set layout binding.

## <a href="#_description" class="anchor"></a>Description

If `bindingCount` is zero or if this structure is not included in the
`pNext` chain, the
[VkDescriptorBindingFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBindingFlags.html) for each
descriptor set layout binding is considered to be zero. Otherwise, the
descriptor set layout binding at
[VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateInfo.html)::`pBindings`\[i\]
uses the flags in `pBindingFlags`\[i\].

Valid Usage

- <a
  href="#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-bindingCount-03002"
  id="VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-bindingCount-03002"></a>
  VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-bindingCount-03002  
  If `bindingCount` is not zero, `bindingCount` **must** equal
  [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateInfo.html)::`bindingCount`

- <a href="#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-flags-03003"
  id="VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-flags-03003"></a>
  VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-flags-03003  
  If
  [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateInfo.html)::`flags`
  includes `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT_KHR`,
  then all elements of `pBindingFlags` **must** not include
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`,
  `VK_DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT`, or
  `VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT`

- <a
  href="#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-pBindingFlags-03004"
  id="VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-pBindingFlags-03004"></a>
  VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-pBindingFlags-03004  
  If an element of `pBindingFlags` includes
  `VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT`, then it
  **must** be the element with the highest `binding` number

- <a
  href="#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingUniformBufferUpdateAfterBind-03005"
  id="VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingUniformBufferUpdateAfterBind-03005"></a>
  VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingUniformBufferUpdateAfterBind-03005  
  If
  [VkPhysicalDeviceDescriptorIndexingFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html)::`descriptorBindingUniformBufferUpdateAfterBind`
  is not enabled, all bindings with descriptor type
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER` **must** not use
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`

- <a
  href="#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingSampledImageUpdateAfterBind-03006"
  id="VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingSampledImageUpdateAfterBind-03006"></a>
  VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingSampledImageUpdateAfterBind-03006  
  If
  [VkPhysicalDeviceDescriptorIndexingFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html)::`descriptorBindingSampledImageUpdateAfterBind`
  is not enabled, all bindings with descriptor type
  `VK_DESCRIPTOR_TYPE_SAMPLER`,
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, or
  `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE` **must** not use
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`

- <a
  href="#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingStorageImageUpdateAfterBind-03007"
  id="VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingStorageImageUpdateAfterBind-03007"></a>
  VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingStorageImageUpdateAfterBind-03007  
  If
  [VkPhysicalDeviceDescriptorIndexingFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html)::`descriptorBindingStorageImageUpdateAfterBind`
  is not enabled, all bindings with descriptor type
  `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE` **must** not use
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`

- <a
  href="#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingStorageBufferUpdateAfterBind-03008"
  id="VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingStorageBufferUpdateAfterBind-03008"></a>
  VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingStorageBufferUpdateAfterBind-03008  
  If
  [VkPhysicalDeviceDescriptorIndexingFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html)::`descriptorBindingStorageBufferUpdateAfterBind`
  is not enabled, all bindings with descriptor type
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER` **must** not use
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`

- <a
  href="#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingUniformTexelBufferUpdateAfterBind-03009"
  id="VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingUniformTexelBufferUpdateAfterBind-03009"></a>
  VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingUniformTexelBufferUpdateAfterBind-03009  
  If
  [VkPhysicalDeviceDescriptorIndexingFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html)::`descriptorBindingUniformTexelBufferUpdateAfterBind`
  is not enabled, all bindings with descriptor type
  `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` **must** not use
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`

- <a
  href="#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingStorageTexelBufferUpdateAfterBind-03010"
  id="VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingStorageTexelBufferUpdateAfterBind-03010"></a>
  VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingStorageTexelBufferUpdateAfterBind-03010  
  If
  [VkPhysicalDeviceDescriptorIndexingFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html)::`descriptorBindingStorageTexelBufferUpdateAfterBind`
  is not enabled, all bindings with descriptor type
  `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` **must** not use
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`

- <a
  href="#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingInlineUniformBlockUpdateAfterBind-02211"
  id="VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingInlineUniformBlockUpdateAfterBind-02211"></a>
  VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingInlineUniformBlockUpdateAfterBind-02211  
  If
  [VkPhysicalDeviceInlineUniformBlockFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceInlineUniformBlockFeatures.html)::`descriptorBindingInlineUniformBlockUpdateAfterBind`
  is not enabled, all bindings with descriptor type
  `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` **must** not use
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`

- <a
  href="#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingAccelerationStructureUpdateAfterBind-03570"
  id="VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingAccelerationStructureUpdateAfterBind-03570"></a>
  VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingAccelerationStructureUpdateAfterBind-03570  
  If
  [VkPhysicalDeviceAccelerationStructureFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceAccelerationStructureFeaturesKHR.html)::`descriptorBindingAccelerationStructureUpdateAfterBind`
  is not enabled, all bindings with descriptor type
  `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR` or
  `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_NV` **must** not use
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`

- <a href="#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-None-03011"
  id="VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-None-03011"></a>
  VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-None-03011  
  All bindings with descriptor type
  `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`,
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`, or
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` **must** not use
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`

- <a
  href="#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingUpdateUnusedWhilePending-03012"
  id="VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingUpdateUnusedWhilePending-03012"></a>
  VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingUpdateUnusedWhilePending-03012  
  If
  [VkPhysicalDeviceDescriptorIndexingFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html)::`descriptorBindingUpdateUnusedWhilePending`
  is not enabled, all elements of `pBindingFlags` **must** not include
  `VK_DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT`

- <a
  href="#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingPartiallyBound-03013"
  id="VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingPartiallyBound-03013"></a>
  VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingPartiallyBound-03013  
  If
  [VkPhysicalDeviceDescriptorIndexingFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html)::`descriptorBindingPartiallyBound`
  is not enabled, all elements of `pBindingFlags` **must** not include
  `VK_DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT`

- <a
  href="#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingVariableDescriptorCount-03014"
  id="VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingVariableDescriptorCount-03014"></a>
  VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-descriptorBindingVariableDescriptorCount-03014  
  If
  [VkPhysicalDeviceDescriptorIndexingFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html)::`descriptorBindingVariableDescriptorCount`
  is not enabled, all elements of `pBindingFlags` **must** not include
  `VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT`

- <a
  href="#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-pBindingFlags-03015"
  id="VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-pBindingFlags-03015"></a>
  VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-pBindingFlags-03015  
  If an element of `pBindingFlags` includes
  `VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT`, that element’s
  `descriptorType` **must** not be
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` or
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`

Valid Usage (Implicit)

- <a href="#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-sType-sType"
  id="VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-sType-sType"></a>
  VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_BINDING_FLAGS_CREATE_INFO`

- <a
  href="#VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-pBindingFlags-parameter"
  id="VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-pBindingFlags-parameter"></a>
  VUID-VkDescriptorSetLayoutBindingFlagsCreateInfo-pBindingFlags-parameter  
  If `bindingCount` is not `0`, `pBindingFlags` **must** be a valid
  pointer to an array of `bindingCount` valid combinations of
  [VkDescriptorBindingFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBindingFlagBits.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_descriptor_indexing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_indexing.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkDescriptorBindingFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBindingFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDescriptorSetLayoutBindingFlagsCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
