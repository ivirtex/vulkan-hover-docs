# VkDescriptorSetLayoutCreateInfo(3) Manual Page

## Name

VkDescriptorSetLayoutCreateInfo - Structure specifying parameters of a
newly created descriptor set layout



## <a href="#_c_specification" class="anchor"></a>C Specification

Information about the descriptor set layout is passed in a
`VkDescriptorSetLayoutCreateInfo` structure:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkDescriptorSetLayoutCreateInfo {
    VkStructureType                        sType;
    const void*                            pNext;
    VkDescriptorSetLayoutCreateFlags       flags;
    uint32_t                               bindingCount;
    const VkDescriptorSetLayoutBinding*    pBindings;
} VkDescriptorSetLayoutCreateInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkDescriptorSetLayoutCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateFlagBits.html)
  specifying options for descriptor set layout creation.

- `bindingCount` is the number of elements in `pBindings`.

- `pBindings` is a pointer to an array of
  [VkDescriptorSetLayoutBinding](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutBinding.html)
  structures.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkDescriptorSetLayoutCreateInfo-binding-00279"
  id="VUID-VkDescriptorSetLayoutCreateInfo-binding-00279"></a>
  VUID-VkDescriptorSetLayoutCreateInfo-binding-00279  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-perStageDescriptorSet"
  target="_blank" rel="noopener"><code>perStageDescriptorSet</code></a>
  feature is not enabled, or `flags` does not contain
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PER_STAGE_BIT_NV`, then the
  [VkDescriptorSetLayoutBinding](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutBinding.html)::`binding`
  members of the elements of the `pBindings` array **must** each have
  different values

- <a href="#VUID-VkDescriptorSetLayoutCreateInfo-flags-00280"
  id="VUID-VkDescriptorSetLayoutCreateInfo-flags-00280"></a>
  VUID-VkDescriptorSetLayoutCreateInfo-flags-00280  
  If `flags` contains
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT_KHR`, then all
  elements of `pBindings` **must** not have a `descriptorType` of
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` or
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`

- <a href="#VUID-VkDescriptorSetLayoutCreateInfo-flags-02208"
  id="VUID-VkDescriptorSetLayoutCreateInfo-flags-02208"></a>
  VUID-VkDescriptorSetLayoutCreateInfo-flags-02208  
  If `flags` contains
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT_KHR`, then all
  elements of `pBindings` **must** not have a `descriptorType` of
  `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`

- <a href="#VUID-VkDescriptorSetLayoutCreateInfo-flags-00281"
  id="VUID-VkDescriptorSetLayoutCreateInfo-flags-00281"></a>
  VUID-VkDescriptorSetLayoutCreateInfo-flags-00281  
  If `flags` contains
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT_KHR`, then the
  total number of elements of all bindings **must** be less than or
  equal to
  [VkPhysicalDevicePushDescriptorPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePushDescriptorPropertiesKHR.html)::`maxPushDescriptors`

- <a href="#VUID-VkDescriptorSetLayoutCreateInfo-flags-04590"
  id="VUID-VkDescriptorSetLayoutCreateInfo-flags-04590"></a>
  VUID-VkDescriptorSetLayoutCreateInfo-flags-04590  
  If `flags` contains
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT_KHR`, `flags`
  **must** not contain
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_EXT`

- <a href="#VUID-VkDescriptorSetLayoutCreateInfo-flags-04591"
  id="VUID-VkDescriptorSetLayoutCreateInfo-flags-04591"></a>
  VUID-VkDescriptorSetLayoutCreateInfo-flags-04591  
  If `flags` contains
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT_KHR`, `pBindings`
  **must** not have a `descriptorType` of
  `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`

- <a href="#VUID-VkDescriptorSetLayoutCreateInfo-flags-03000"
  id="VUID-VkDescriptorSetLayoutCreateInfo-flags-03000"></a>
  VUID-VkDescriptorSetLayoutCreateInfo-flags-03000  
  If any binding has the `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`
  bit set, `flags` **must** include
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT`

- <a href="#VUID-VkDescriptorSetLayoutCreateInfo-descriptorType-03001"
  id="VUID-VkDescriptorSetLayoutCreateInfo-descriptorType-03001"></a>
  VUID-VkDescriptorSetLayoutCreateInfo-descriptorType-03001  
  If any binding has the `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`
  bit set, then all bindings **must** not have `descriptorType` of
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` or
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`

- <a href="#VUID-VkDescriptorSetLayoutCreateInfo-flags-04592"
  id="VUID-VkDescriptorSetLayoutCreateInfo-flags-04592"></a>
  VUID-VkDescriptorSetLayoutCreateInfo-flags-04592  
  If `flags` contains
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT`, `flags`
  **must** not contain
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_EXT`

- <a href="#VUID-VkDescriptorSetLayoutCreateInfo-pBindings-07303"
  id="VUID-VkDescriptorSetLayoutCreateInfo-pBindings-07303"></a>
  VUID-VkDescriptorSetLayoutCreateInfo-pBindings-07303  
  If any element `pBindings`\[i\] has a `descriptorType` of
  `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`, then the `pNext` chain **must**
  include a
  [VkMutableDescriptorTypeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMutableDescriptorTypeCreateInfoEXT.html)
  structure, and `mutableDescriptorTypeListCount` **must** be greater
  than i

- <a href="#VUID-VkDescriptorSetLayoutCreateInfo-descriptorType-04594"
  id="VUID-VkDescriptorSetLayoutCreateInfo-descriptorType-04594"></a>
  VUID-VkDescriptorSetLayoutCreateInfo-descriptorType-04594  
  If a binding has a `descriptorType` value of
  `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`, then `pImmutableSamplers` **must**
  be `NULL`

- <a
  href="#VUID-VkDescriptorSetLayoutCreateInfo-mutableDescriptorType-04595"
  id="VUID-VkDescriptorSetLayoutCreateInfo-mutableDescriptorType-04595"></a>
  VUID-VkDescriptorSetLayoutCreateInfo-mutableDescriptorType-04595  
  If
  [VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT.html)::`mutableDescriptorType`
  is not enabled, `pBindings` **must** not contain a `descriptorType` of
  `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`

- <a href="#VUID-VkDescriptorSetLayoutCreateInfo-flags-04596"
  id="VUID-VkDescriptorSetLayoutCreateInfo-flags-04596"></a>
  VUID-VkDescriptorSetLayoutCreateInfo-flags-04596  
  If `flags` contains
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_EXT`,
  [VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT.html)::`mutableDescriptorType`
  **must** be enabled

- <a href="#VUID-VkDescriptorSetLayoutCreateInfo-flags-08000"
  id="VUID-VkDescriptorSetLayoutCreateInfo-flags-08000"></a>
  VUID-VkDescriptorSetLayoutCreateInfo-flags-08000  
  If `flags` contains
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`, then all
  elements of `pBindings` **must** not have a `descriptorType` of
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` or
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`

- <a href="#VUID-VkDescriptorSetLayoutCreateInfo-flags-08001"
  id="VUID-VkDescriptorSetLayoutCreateInfo-flags-08001"></a>
  VUID-VkDescriptorSetLayoutCreateInfo-flags-08001  
  If `flags` contains
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_EMBEDDED_IMMUTABLE_SAMPLERS_BIT_EXT`,
  `flags` **must** also contain
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`

- <a href="#VUID-VkDescriptorSetLayoutCreateInfo-flags-08002"
  id="VUID-VkDescriptorSetLayoutCreateInfo-flags-08002"></a>
  VUID-VkDescriptorSetLayoutCreateInfo-flags-08002  
  If `flags` contains
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`, then
  `flags` **must** not contain
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT`

- <a href="#VUID-VkDescriptorSetLayoutCreateInfo-flags-08003"
  id="VUID-VkDescriptorSetLayoutCreateInfo-flags-08003"></a>
  VUID-VkDescriptorSetLayoutCreateInfo-flags-08003  
  If `flags` contains
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`, then
  `flags` **must** not contain
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_VALVE`

- <a href="#VUID-VkDescriptorSetLayoutCreateInfo-flags-09463"
  id="VUID-VkDescriptorSetLayoutCreateInfo-flags-09463"></a>
  VUID-VkDescriptorSetLayoutCreateInfo-flags-09463  
  If `flags` contains
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PER_STAGE_BIT_NV`, then <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-perStageDescriptorSet"
  target="_blank" rel="noopener"><code>perStageDescriptorSet</code></a>
  **must** be enabled

- <a href="#VUID-VkDescriptorSetLayoutCreateInfo-flags-09464"
  id="VUID-VkDescriptorSetLayoutCreateInfo-flags-09464"></a>
  VUID-VkDescriptorSetLayoutCreateInfo-flags-09464  
  If `flags` contains
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PER_STAGE_BIT_NV`, then there
  **must** not be any two elements of the `pBindings` array with the
  same
  [VkDescriptorSetLayoutBinding](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutBinding.html)::`binding`
  value and their
  [VkDescriptorSetLayoutBinding](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutBinding.html)::`stageFlags`
  containing the same bit

Valid Usage (Implicit)

- <a href="#VUID-VkDescriptorSetLayoutCreateInfo-sType-sType"
  id="VUID-VkDescriptorSetLayoutCreateInfo-sType-sType"></a>
  VUID-VkDescriptorSetLayoutCreateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_CREATE_INFO`

- <a href="#VUID-VkDescriptorSetLayoutCreateInfo-pNext-pNext"
  id="VUID-VkDescriptorSetLayoutCreateInfo-pNext-pNext"></a>
  VUID-VkDescriptorSetLayoutCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkDescriptorSetLayoutBindingFlagsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutBindingFlagsCreateInfo.html)
  or
  [VkMutableDescriptorTypeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMutableDescriptorTypeCreateInfoEXT.html)

- <a href="#VUID-VkDescriptorSetLayoutCreateInfo-sType-unique"
  id="VUID-VkDescriptorSetLayoutCreateInfo-sType-unique"></a>
  VUID-VkDescriptorSetLayoutCreateInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkDescriptorSetLayoutCreateInfo-flags-parameter"
  id="VUID-VkDescriptorSetLayoutCreateInfo-flags-parameter"></a>
  VUID-VkDescriptorSetLayoutCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of
  [VkDescriptorSetLayoutCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateFlagBits.html)
  values

- <a href="#VUID-VkDescriptorSetLayoutCreateInfo-pBindings-parameter"
  id="VUID-VkDescriptorSetLayoutCreateInfo-pBindings-parameter"></a>
  VUID-VkDescriptorSetLayoutCreateInfo-pBindings-parameter  
  If `bindingCount` is not `0`, `pBindings` **must** be a valid pointer
  to an array of `bindingCount` valid
  [VkDescriptorSetLayoutBinding](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutBinding.html)
  structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDescriptorSetLayoutBinding](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutBinding.html),
[VkDescriptorSetLayoutCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDescriptorSetLayout.html),
[vkGetDescriptorSetLayoutSupport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDescriptorSetLayoutSupport.html),
[vkGetDescriptorSetLayoutSupportKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDescriptorSetLayoutSupportKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDescriptorSetLayoutCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
