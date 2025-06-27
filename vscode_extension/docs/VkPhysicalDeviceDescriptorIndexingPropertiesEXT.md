# VkPhysicalDeviceDescriptorIndexingProperties(3) Manual Page

## Name

VkPhysicalDeviceDescriptorIndexingProperties - Structure describing
descriptor indexing properties that can be supported by an
implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceDescriptorIndexingProperties` structure is defined
as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkPhysicalDeviceDescriptorIndexingProperties {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxUpdateAfterBindDescriptorsInAllPools;
    VkBool32           shaderUniformBufferArrayNonUniformIndexingNative;
    VkBool32           shaderSampledImageArrayNonUniformIndexingNative;
    VkBool32           shaderStorageBufferArrayNonUniformIndexingNative;
    VkBool32           shaderStorageImageArrayNonUniformIndexingNative;
    VkBool32           shaderInputAttachmentArrayNonUniformIndexingNative;
    VkBool32           robustBufferAccessUpdateAfterBind;
    VkBool32           quadDivergentImplicitLod;
    uint32_t           maxPerStageDescriptorUpdateAfterBindSamplers;
    uint32_t           maxPerStageDescriptorUpdateAfterBindUniformBuffers;
    uint32_t           maxPerStageDescriptorUpdateAfterBindStorageBuffers;
    uint32_t           maxPerStageDescriptorUpdateAfterBindSampledImages;
    uint32_t           maxPerStageDescriptorUpdateAfterBindStorageImages;
    uint32_t           maxPerStageDescriptorUpdateAfterBindInputAttachments;
    uint32_t           maxPerStageUpdateAfterBindResources;
    uint32_t           maxDescriptorSetUpdateAfterBindSamplers;
    uint32_t           maxDescriptorSetUpdateAfterBindUniformBuffers;
    uint32_t           maxDescriptorSetUpdateAfterBindUniformBuffersDynamic;
    uint32_t           maxDescriptorSetUpdateAfterBindStorageBuffers;
    uint32_t           maxDescriptorSetUpdateAfterBindStorageBuffersDynamic;
    uint32_t           maxDescriptorSetUpdateAfterBindSampledImages;
    uint32_t           maxDescriptorSetUpdateAfterBindStorageImages;
    uint32_t           maxDescriptorSetUpdateAfterBindInputAttachments;
} VkPhysicalDeviceDescriptorIndexingProperties;
```

or the equivalent

``` c
// Provided by VK_EXT_descriptor_indexing
typedef VkPhysicalDeviceDescriptorIndexingProperties VkPhysicalDeviceDescriptorIndexingPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-limits-maxUpdateAfterBindDescriptorsInAllPools"></span>
  `maxUpdateAfterBindDescriptorsInAllPools` is the maximum number of
  descriptors (summed over all descriptor types) that **can** be created
  across all pools that are created with the
  `VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT` bit set. Pool
  creation **may** fail when this limit is exceeded, or when the space
  this limit represents is unable to satisfy a pool creation due to
  fragmentation.

- <span id="extension-limits-shaderUniformBufferArrayNonUniformIndexingNative"></span>
  `shaderUniformBufferArrayNonUniformIndexingNative` is a boolean value
  indicating whether uniform buffer descriptors natively support
  nonuniform indexing. If this is `VK_FALSE`, then a single dynamic
  instance of an instruction that nonuniformly indexes an array of
  uniform buffers **may** execute multiple times in order to access all
  the descriptors.

- <span id="extension-limits-shaderSampledImageArrayNonUniformIndexingNative"></span>
  `shaderSampledImageArrayNonUniformIndexingNative` is a boolean value
  indicating whether sampler and image descriptors natively support
  nonuniform indexing. If this is `VK_FALSE`, then a single dynamic
  instance of an instruction that nonuniformly indexes an array of
  samplers or images **may** execute multiple times in order to access
  all the descriptors.

- <span id="extension-limits-shaderStorageBufferArrayNonUniformIndexingNative"></span>
  `shaderStorageBufferArrayNonUniformIndexingNative` is a boolean value
  indicating whether storage buffer descriptors natively support
  nonuniform indexing. If this is `VK_FALSE`, then a single dynamic
  instance of an instruction that nonuniformly indexes an array of
  storage buffers **may** execute multiple times in order to access all
  the descriptors.

- <span id="extension-limits-shaderStorageImageArrayNonUniformIndexingNative"></span>
  `shaderStorageImageArrayNonUniformIndexingNative` is a boolean value
  indicating whether storage image descriptors natively support
  nonuniform indexing. If this is `VK_FALSE`, then a single dynamic
  instance of an instruction that nonuniformly indexes an array of
  storage images **may** execute multiple times in order to access all
  the descriptors.

- <span id="extension-limits-shaderInputAttachmentArrayNonUniformIndexingNative"></span>
  `shaderInputAttachmentArrayNonUniformIndexingNative` is a boolean
  value indicating whether input attachment descriptors natively support
  nonuniform indexing. If this is `VK_FALSE`, then a single dynamic
  instance of an instruction that nonuniformly indexes an array of input
  attachments **may** execute multiple times in order to access all the
  descriptors.

- <span id="extension-limits-robustBufferAccessUpdateAfterBind"></span>
  `robustBufferAccessUpdateAfterBind` is a boolean value indicating
  whether <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustBufferAccess"
  target="_blank" rel="noopener"><code>robustBufferAccess</code></a>
  **can** be enabled on a device simultaneously with
  `descriptorBindingUniformBufferUpdateAfterBind`,
  `descriptorBindingStorageBufferUpdateAfterBind`,
  `descriptorBindingUniformTexelBufferUpdateAfterBind`, and/or
  `descriptorBindingStorageTexelBufferUpdateAfterBind`. If this is
  `VK_FALSE`, then either `robustBufferAccess` **must** be disabled or
  all of these update-after-bind features **must** be disabled.

- <span id="extension-limits-quadDivergentImplicitLod"></span>
  `quadDivergentImplicitLod` is a boolean value indicating whether
  implicit LOD calculations for image operations have well-defined
  results when the image and/or sampler objects used for the instruction
  are not uniform within a quad. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-derivative-image-operations"
  target="_blank" rel="noopener">Derivative Image Operations</a>.

- <span id="extension-limits-maxPerStageDescriptorUpdateAfterBindSamplers"></span>
  `maxPerStageDescriptorUpdateAfterBindSamplers` is similar to
  `maxPerStageDescriptorSamplers` but counts descriptors from descriptor
  sets created with or without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.

- <span id="extension-limits-maxPerStageDescriptorUpdateAfterBindUniformBuffers"></span>
  `maxPerStageDescriptorUpdateAfterBindUniformBuffers` is similar to
  `maxPerStageDescriptorUniformBuffers` but counts descriptors from
  descriptor sets created with or without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.

- <span id="extension-limits-maxPerStageDescriptorUpdateAfterBindStorageBuffers"></span>
  `maxPerStageDescriptorUpdateAfterBindStorageBuffers` is similar to
  `maxPerStageDescriptorStorageBuffers` but counts descriptors from
  descriptor sets created with or without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.

- <span id="extension-limits-maxPerStageDescriptorUpdateAfterBindSampledImages"></span>
  `maxPerStageDescriptorUpdateAfterBindSampledImages` is similar to
  `maxPerStageDescriptorSampledImages` but counts descriptors from
  descriptor sets created with or without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.

- <span id="extension-limits-maxPerStageDescriptorUpdateAfterBindStorageImages"></span>
  `maxPerStageDescriptorUpdateAfterBindStorageImages` is similar to
  `maxPerStageDescriptorStorageImages` but counts descriptors from
  descriptor sets created with or without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.

- <span id="extension-limits-maxPerStageDescriptorUpdateAfterBindInputAttachments"></span>
  `maxPerStageDescriptorUpdateAfterBindInputAttachments` is similar to
  `maxPerStageDescriptorInputAttachments` but counts descriptors from
  descriptor sets created with or without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.

- <span id="extension-limits-maxPerStageUpdateAfterBindResources"></span>
  `maxPerStageUpdateAfterBindResources` is similar to
  `maxPerStageResources` but counts descriptors from descriptor sets
  created with or without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.

- <span id="extension-limits-maxDescriptorSetUpdateAfterBindSamplers"></span>
  `maxDescriptorSetUpdateAfterBindSamplers` is similar to
  `maxDescriptorSetSamplers` but counts descriptors from descriptor sets
  created with or without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.

- <span id="extension-limits-maxDescriptorSetUpdateAfterBindUniformBuffers"></span>
  `maxDescriptorSetUpdateAfterBindUniformBuffers` is similar to
  `maxDescriptorSetUniformBuffers` but counts descriptors from
  descriptor sets created with or without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.

- <span id="extension-limits-maxDescriptorSetUpdateAfterBindUniformBuffersDynamic"></span>
  `maxDescriptorSetUpdateAfterBindUniformBuffersDynamic` is similar to
  `maxDescriptorSetUniformBuffersDynamic` but counts descriptors from
  descriptor sets created with or without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.
  While an application **can** allocate dynamic uniform buffer
  descriptors from a pool created with the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT`, bindings
  for these descriptors **must** not be present in any descriptor set
  layout that includes bindings created with
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`.

- <span id="extension-limits-maxDescriptorSetUpdateAfterBindStorageBuffers"></span>
  `maxDescriptorSetUpdateAfterBindStorageBuffers` is similar to
  `maxDescriptorSetStorageBuffers` but counts descriptors from
  descriptor sets created with or without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.

- <span id="extension-limits-maxDescriptorSetUpdateAfterBindStorageBuffersDynamic"></span>
  `maxDescriptorSetUpdateAfterBindStorageBuffersDynamic` is similar to
  `maxDescriptorSetStorageBuffersDynamic` but counts descriptors from
  descriptor sets created with or without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.
  While an application **can** allocate dynamic storage buffer
  descriptors from a pool created with the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT`, bindings
  for these descriptors **must** not be present in any descriptor set
  layout that includes bindings created with
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`.

- <span id="extension-limits-maxDescriptorSetUpdateAfterBindSampledImages"></span>
  `maxDescriptorSetUpdateAfterBindSampledImages` is similar to
  `maxDescriptorSetSampledImages` but counts descriptors from descriptor
  sets created with or without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.

- <span id="extension-limits-maxDescriptorSetUpdateAfterBindStorageImages"></span>
  `maxDescriptorSetUpdateAfterBindStorageImages` is similar to
  `maxDescriptorSetStorageImages` but counts descriptors from descriptor
  sets created with or without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.

- <span id="extension-limits-maxDescriptorSetUpdateAfterBindInputAttachments"></span>
  `maxDescriptorSetUpdateAfterBindInputAttachments` is similar to
  `maxDescriptorSetInputAttachments` but counts descriptors from
  descriptor sets created with or without the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.

If the `VkPhysicalDeviceDescriptorIndexingProperties` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceDescriptorIndexingProperties-sType-sType"
  id="VUID-VkPhysicalDeviceDescriptorIndexingProperties-sType-sType"></a>
  VUID-VkPhysicalDeviceDescriptorIndexingProperties-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_INDEXING_PROPERTIES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_descriptor_indexing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_indexing.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceDescriptorIndexingProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
