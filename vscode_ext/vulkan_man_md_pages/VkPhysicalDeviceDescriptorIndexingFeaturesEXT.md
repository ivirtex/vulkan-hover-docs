# VkPhysicalDeviceDescriptorIndexingFeatures(3) Manual Page

## Name

VkPhysicalDeviceDescriptorIndexingFeatures - Structure describing
descriptor indexing features that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceDescriptorIndexingFeatures` structure is defined
as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkPhysicalDeviceDescriptorIndexingFeatures {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           shaderInputAttachmentArrayDynamicIndexing;
    VkBool32           shaderUniformTexelBufferArrayDynamicIndexing;
    VkBool32           shaderStorageTexelBufferArrayDynamicIndexing;
    VkBool32           shaderUniformBufferArrayNonUniformIndexing;
    VkBool32           shaderSampledImageArrayNonUniformIndexing;
    VkBool32           shaderStorageBufferArrayNonUniformIndexing;
    VkBool32           shaderStorageImageArrayNonUniformIndexing;
    VkBool32           shaderInputAttachmentArrayNonUniformIndexing;
    VkBool32           shaderUniformTexelBufferArrayNonUniformIndexing;
    VkBool32           shaderStorageTexelBufferArrayNonUniformIndexing;
    VkBool32           descriptorBindingUniformBufferUpdateAfterBind;
    VkBool32           descriptorBindingSampledImageUpdateAfterBind;
    VkBool32           descriptorBindingStorageImageUpdateAfterBind;
    VkBool32           descriptorBindingStorageBufferUpdateAfterBind;
    VkBool32           descriptorBindingUniformTexelBufferUpdateAfterBind;
    VkBool32           descriptorBindingStorageTexelBufferUpdateAfterBind;
    VkBool32           descriptorBindingUpdateUnusedWhilePending;
    VkBool32           descriptorBindingPartiallyBound;
    VkBool32           descriptorBindingVariableDescriptorCount;
    VkBool32           runtimeDescriptorArray;
} VkPhysicalDeviceDescriptorIndexingFeatures;
```

or the equivalent

``` c
// Provided by VK_EXT_descriptor_indexing
typedef VkPhysicalDeviceDescriptorIndexingFeatures VkPhysicalDeviceDescriptorIndexingFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-features-shaderInputAttachmentArrayDynamicIndexing"></span>
  `shaderInputAttachmentArrayDynamicIndexing` indicates whether arrays
  of input attachments **can** be indexed by dynamically uniform integer
  expressions in shader code. If this feature is not enabled, resources
  with a descriptor type of `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`
  **must** be indexed only by constant integral expressions when
  aggregated into arrays in shader code. This also indicates whether
  shader modules **can** declare the
  `InputAttachmentArrayDynamicIndexing` capability.

- <span id="extension-features-shaderUniformTexelBufferArrayDynamicIndexing"></span>
  `shaderUniformTexelBufferArrayDynamicIndexing` indicates whether
  arrays of uniform texel buffers **can** be indexed by dynamically
  uniform integer expressions in shader code. If this feature is not
  enabled, resources with a descriptor type of
  `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` **must** be indexed only by
  constant integral expressions when aggregated into arrays in shader
  code. This also indicates whether shader modules **can** declare the
  `UniformTexelBufferArrayDynamicIndexing` capability.

- <span id="extension-features-shaderStorageTexelBufferArrayDynamicIndexing"></span>
  `shaderStorageTexelBufferArrayDynamicIndexing` indicates whether
  arrays of storage texel buffers **can** be indexed by dynamically
  uniform integer expressions in shader code. If this feature is not
  enabled, resources with a descriptor type of
  `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` **must** be indexed only by
  constant integral expressions when aggregated into arrays in shader
  code. This also indicates whether shader modules **can** declare the
  `StorageTexelBufferArrayDynamicIndexing` capability.

- <span id="extension-features-shaderUniformBufferArrayNonUniformIndexing"></span>
  `shaderUniformBufferArrayNonUniformIndexing` indicates whether arrays
  of uniform buffers **can** be indexed by non-uniform integer
  expressions in shader code. If this feature is not enabled, resources
  with a descriptor type of `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER` or
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` **must** not be indexed by
  non-uniform integer expressions when aggregated into arrays in shader
  code. This also indicates whether shader modules **can** declare the
  `UniformBufferArrayNonUniformIndexing` capability.

- <span id="extension-features-shaderSampledImageArrayNonUniformIndexing"></span>
  `shaderSampledImageArrayNonUniformIndexing` indicates whether arrays
  of samplers or sampled images **can** be indexed by non-uniform
  integer expressions in shader code. If this feature is not enabled,
  resources with a descriptor type of `VK_DESCRIPTOR_TYPE_SAMPLER`,
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, or
  `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE` **must** not be indexed by
  non-uniform integer expressions when aggregated into arrays in shader
  code. This also indicates whether shader modules **can** declare the
  `SampledImageArrayNonUniformIndexing` capability.

- <span id="extension-features-shaderStorageBufferArrayNonUniformIndexing"></span>
  `shaderStorageBufferArrayNonUniformIndexing` indicates whether arrays
  of storage buffers **can** be indexed by non-uniform integer
  expressions in shader code. If this feature is not enabled, resources
  with a descriptor type of `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER` or
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` **must** not be indexed by
  non-uniform integer expressions when aggregated into arrays in shader
  code. This also indicates whether shader modules **can** declare the
  `StorageBufferArrayNonUniformIndexing` capability.

- <span id="extension-features-shaderStorageImageArrayNonUniformIndexing"></span>
  `shaderStorageImageArrayNonUniformIndexing` indicates whether arrays
  of storage images **can** be indexed by non-uniform integer
  expressions in shader code. If this feature is not enabled, resources
  with a descriptor type of `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE` **must**
  not be indexed by non-uniform integer expressions when aggregated into
  arrays in shader code. This also indicates whether shader modules
  **can** declare the `StorageImageArrayNonUniformIndexing` capability.

- <span id="extension-features-shaderInputAttachmentArrayNonUniformIndexing"></span>
  `shaderInputAttachmentArrayNonUniformIndexing` indicates whether
  arrays of input attachments **can** be indexed by non-uniform integer
  expressions in shader code. If this feature is not enabled, resources
  with a descriptor type of `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`
  **must** not be indexed by non-uniform integer expressions when
  aggregated into arrays in shader code. This also indicates whether
  shader modules **can** declare the
  `InputAttachmentArrayNonUniformIndexing` capability.

- <span id="extension-features-shaderUniformTexelBufferArrayNonUniformIndexing"></span>
  `shaderUniformTexelBufferArrayNonUniformIndexing` indicates whether
  arrays of uniform texel buffers **can** be indexed by non-uniform
  integer expressions in shader code. If this feature is not enabled,
  resources with a descriptor type of
  `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` **must** not be indexed by
  non-uniform integer expressions when aggregated into arrays in shader
  code. This also indicates whether shader modules **can** declare the
  `UniformTexelBufferArrayNonUniformIndexing` capability.

- <span id="extension-features-shaderStorageTexelBufferArrayNonUniformIndexing"></span>
  `shaderStorageTexelBufferArrayNonUniformIndexing` indicates whether
  arrays of storage texel buffers **can** be indexed by non-uniform
  integer expressions in shader code. If this feature is not enabled,
  resources with a descriptor type of
  `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` **must** not be indexed by
  non-uniform integer expressions when aggregated into arrays in shader
  code. This also indicates whether shader modules **can** declare the
  `StorageTexelBufferArrayNonUniformIndexing` capability.

- <span id="extension-features-descriptorBindingUniformBufferUpdateAfterBind"></span>
  `descriptorBindingUniformBufferUpdateAfterBind` indicates whether the
  implementation supports updating uniform buffer descriptors after a
  set is bound. If this feature is not enabled,
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` **must** not be used
  with `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER`.

- <span id="extension-features-descriptorBindingSampledImageUpdateAfterBind"></span>
  `descriptorBindingSampledImageUpdateAfterBind` indicates whether the
  implementation supports updating sampled image descriptors after a set
  is bound. If this feature is not enabled,
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` **must** not be used
  with `VK_DESCRIPTOR_TYPE_SAMPLER`,
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, or
  `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`.

- <span id="extension-features-descriptorBindingStorageImageUpdateAfterBind"></span>
  `descriptorBindingStorageImageUpdateAfterBind` indicates whether the
  implementation supports updating storage image descriptors after a set
  is bound. If this feature is not enabled,
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` **must** not be used
  with `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`.

- <span id="extension-features-descriptorBindingStorageBufferUpdateAfterBind"></span>
  `descriptorBindingStorageBufferUpdateAfterBind` indicates whether the
  implementation supports updating storage buffer descriptors after a
  set is bound. If this feature is not enabled,
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` **must** not be used
  with `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER`.

- <span id="extension-features-descriptorBindingUniformTexelBufferUpdateAfterBind"></span>
  `descriptorBindingUniformTexelBufferUpdateAfterBind` indicates whether
  the implementation supports updating uniform texel buffer descriptors
  after a set is bound. If this feature is not enabled,
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` **must** not be used
  with `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER`.

- <span id="extension-features-descriptorBindingStorageTexelBufferUpdateAfterBind"></span>
  `descriptorBindingStorageTexelBufferUpdateAfterBind` indicates whether
  the implementation supports updating storage texel buffer descriptors
  after a set is bound. If this feature is not enabled,
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` **must** not be used
  with `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER`.

- <span id="extension-features-descriptorBindingUpdateUnusedWhilePending"></span>
  `descriptorBindingUpdateUnusedWhilePending` indicates whether the
  implementation supports updating descriptors while the set is in use.
  If this feature is not enabled,
  `VK_DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT` **must** not
  be used.

- <span id="extension-features-descriptorBindingPartiallyBound"></span>
  `descriptorBindingPartiallyBound` indicates whether the implementation
  supports statically using a descriptor set binding in which some
  descriptors are not valid. If this feature is not enabled,
  `VK_DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT` **must** not be used.

- <span id="extension-features-descriptorBindingVariableDescriptorCount"></span>
  `descriptorBindingVariableDescriptorCount` indicates whether the
  implementation supports descriptor sets with a variable-sized last
  binding. If this feature is not enabled,
  `VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT` **must** not be
  used.

- <span id="extension-features-runtimeDescriptorArray"></span>
  `runtimeDescriptorArray` indicates whether the implementation supports
  the SPIR-V `RuntimeDescriptorArray` capability. If this feature is not
  enabled, descriptors **must** not be declared in runtime arrays.

If the `VkPhysicalDeviceDescriptorIndexingFeatures` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceDescriptorIndexingFeatures` **can** also be used in the
`pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceDescriptorIndexingFeatures-sType-sType"
  id="VUID-VkPhysicalDeviceDescriptorIndexingFeatures-sType-sType"></a>
  VUID-VkPhysicalDeviceDescriptorIndexingFeatures-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_FEATURES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_descriptor_indexing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_indexing.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceDescriptorIndexingFeatures"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
