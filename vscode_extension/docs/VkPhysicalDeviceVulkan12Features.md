# VkPhysicalDeviceVulkan12Features(3) Manual Page

## Name

VkPhysicalDeviceVulkan12Features - Structure describing the Vulkan 1.2
features that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceVulkan12Features` structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkPhysicalDeviceVulkan12Features {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           samplerMirrorClampToEdge;
    VkBool32           drawIndirectCount;
    VkBool32           storageBuffer8BitAccess;
    VkBool32           uniformAndStorageBuffer8BitAccess;
    VkBool32           storagePushConstant8;
    VkBool32           shaderBufferInt64Atomics;
    VkBool32           shaderSharedInt64Atomics;
    VkBool32           shaderFloat16;
    VkBool32           shaderInt8;
    VkBool32           descriptorIndexing;
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
    VkBool32           samplerFilterMinmax;
    VkBool32           scalarBlockLayout;
    VkBool32           imagelessFramebuffer;
    VkBool32           uniformBufferStandardLayout;
    VkBool32           shaderSubgroupExtendedTypes;
    VkBool32           separateDepthStencilLayouts;
    VkBool32           hostQueryReset;
    VkBool32           timelineSemaphore;
    VkBool32           bufferDeviceAddress;
    VkBool32           bufferDeviceAddressCaptureReplay;
    VkBool32           bufferDeviceAddressMultiDevice;
    VkBool32           vulkanMemoryModel;
    VkBool32           vulkanMemoryModelDeviceScope;
    VkBool32           vulkanMemoryModelAvailabilityVisibilityChains;
    VkBool32           shaderOutputViewportIndex;
    VkBool32           shaderOutputLayer;
    VkBool32           subgroupBroadcastDynamicId;
} VkPhysicalDeviceVulkan12Features;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="features-samplerMirrorClampToEdge"></span>
  `samplerMirrorClampToEdge` indicates whether the implementation
  supports the `VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE` sampler
  address mode. If this feature is not enabled, the
  `VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE` sampler address mode
  **must** not be used.

- <span id="features-drawIndirectCount"></span> `drawIndirectCount`
  indicates whether the implementation supports the
  [vkCmdDrawIndirectCount](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndirectCount.html) and
  [vkCmdDrawIndexedIndirectCount](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDrawIndexedIndirectCount.html)
  functions. If this feature is not enabled, these functions **must**
  not be used.

- <span id="features-storageBuffer8BitAccess"></span>
  `storageBuffer8BitAccess` indicates whether objects in the
  `StorageBuffer`, `ShaderRecordBufferKHR`, or `PhysicalStorageBuffer`
  storage class with the `Block` decoration **can** have 8-bit integer
  members. If this feature is not enabled, 8-bit integer members
  **must** not be used in such objects. This also indicates whether
  shader modules **can** declare the `StorageBuffer8BitAccess`
  capability.

- <span id="features-uniformAndStorageBuffer8BitAccess"></span>
  `uniformAndStorageBuffer8BitAccess` indicates whether objects in the
  `Uniform` storage class with the `Block` decoration **can** have 8-bit
  integer members. If this feature is not enabled, 8-bit integer members
  **must** not be used in such objects. This also indicates whether
  shader modules **can** declare the `UniformAndStorageBuffer8BitAccess`
  capability.

- <span id="features-storagePushConstant8"></span>
  `storagePushConstant8` indicates whether objects in the `PushConstant`
  storage class **can** have 8-bit integer members. If this feature is
  not enabled, 8-bit integer members **must** not be used in such
  objects. This also indicates whether shader modules **can** declare
  the `StoragePushConstant8` capability.

- <span id="features-shaderBufferInt64Atomics"></span>
  `shaderBufferInt64Atomics` indicates whether shaders **can** perform
  64-bit unsigned and signed integer atomic operations on buffers.

- <span id="features-shaderSharedInt64Atomics"></span>
  `shaderSharedInt64Atomics` indicates whether shaders **can** perform
  64-bit unsigned and signed integer atomic operations on shared and
  payload memory.

- <span id="features-shaderFloat16"></span> `shaderFloat16` indicates
  whether 16-bit floats (halfs) are supported in shader code. This also
  indicates whether shader modules **can** declare the `Float16`
  capability. However, this only enables a subset of the storage classes
  that SPIR-V allows for the `Float16` SPIR-V capability: Declaring and
  using 16-bit floats in the `Private`, `Workgroup` (for non-Block
  variables), and `Function` storage classes is enabled, while declaring
  them in the interface storage classes (e.g., `UniformConstant`,
  `Uniform`, `StorageBuffer`, `Input`, `Output`, and `PushConstant`) is
  not enabled.

- <span id="features-shaderInt8"></span> `shaderInt8` indicates whether
  8-bit integers (signed and unsigned) are supported in shader code.
  This also indicates whether shader modules **can** declare the `Int8`
  capability. However, this only enables a subset of the storage classes
  that SPIR-V allows for the `Int8` SPIR-V capability: Declaring and
  using 8-bit integers in the `Private`, `Workgroup` (for non-Block
  variables), and `Function` storage classes is enabled, while declaring
  them in the interface storage classes (e.g., `UniformConstant`,
  `Uniform`, `StorageBuffer`, `Input`, `Output`, and `PushConstant`) is
  not enabled.

- <span id="features-descriptorIndexing"></span> `descriptorIndexing`
  indicates whether the implementation supports the minimum set of
  descriptor indexing features as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-requirements"
  target="_blank" rel="noopener">Feature Requirements</a> section.
  Enabling the `descriptorIndexing` member when
  [vkCreateDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDevice.html) is called does not imply the
  other minimum descriptor indexing features are also enabled. Those
  other descriptor indexing features **must** be enabled individually as
  needed by the application.

- <span id="features-shaderInputAttachmentArrayDynamicIndexing"></span>
  `shaderInputAttachmentArrayDynamicIndexing` indicates whether arrays
  of input attachments **can** be indexed by dynamically uniform integer
  expressions in shader code. If this feature is not enabled, resources
  with a descriptor type of `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`
  **must** be indexed only by constant integral expressions when
  aggregated into arrays in shader code. This also indicates whether
  shader modules **can** declare the
  `InputAttachmentArrayDynamicIndexing` capability.

- <span id="features-shaderUniformTexelBufferArrayDynamicIndexing"></span>
  `shaderUniformTexelBufferArrayDynamicIndexing` indicates whether
  arrays of uniform texel buffers **can** be indexed by dynamically
  uniform integer expressions in shader code. If this feature is not
  enabled, resources with a descriptor type of
  `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` **must** be indexed only by
  constant integral expressions when aggregated into arrays in shader
  code. This also indicates whether shader modules **can** declare the
  `UniformTexelBufferArrayDynamicIndexing` capability.

- <span id="features-shaderStorageTexelBufferArrayDynamicIndexing"></span>
  `shaderStorageTexelBufferArrayDynamicIndexing` indicates whether
  arrays of storage texel buffers **can** be indexed by dynamically
  uniform integer expressions in shader code. If this feature is not
  enabled, resources with a descriptor type of
  `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` **must** be indexed only by
  constant integral expressions when aggregated into arrays in shader
  code. This also indicates whether shader modules **can** declare the
  `StorageTexelBufferArrayDynamicIndexing` capability.

- <span id="features-shaderUniformBufferArrayNonUniformIndexing"></span>
  `shaderUniformBufferArrayNonUniformIndexing` indicates whether arrays
  of uniform buffers **can** be indexed by non-uniform integer
  expressions in shader code. If this feature is not enabled, resources
  with a descriptor type of `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER` or
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` **must** not be indexed by
  non-uniform integer expressions when aggregated into arrays in shader
  code. This also indicates whether shader modules **can** declare the
  `UniformBufferArrayNonUniformIndexing` capability.

- <span id="features-shaderSampledImageArrayNonUniformIndexing"></span>
  `shaderSampledImageArrayNonUniformIndexing` indicates whether arrays
  of samplers or sampled images **can** be indexed by non-uniform
  integer expressions in shader code. If this feature is not enabled,
  resources with a descriptor type of `VK_DESCRIPTOR_TYPE_SAMPLER`,
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, or
  `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE` **must** not be indexed by
  non-uniform integer expressions when aggregated into arrays in shader
  code. This also indicates whether shader modules **can** declare the
  `SampledImageArrayNonUniformIndexing` capability.

- <span id="features-shaderStorageBufferArrayNonUniformIndexing"></span>
  `shaderStorageBufferArrayNonUniformIndexing` indicates whether arrays
  of storage buffers **can** be indexed by non-uniform integer
  expressions in shader code. If this feature is not enabled, resources
  with a descriptor type of `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER` or
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` **must** not be indexed by
  non-uniform integer expressions when aggregated into arrays in shader
  code. This also indicates whether shader modules **can** declare the
  `StorageBufferArrayNonUniformIndexing` capability.

- <span id="features-shaderStorageImageArrayNonUniformIndexing"></span>
  `shaderStorageImageArrayNonUniformIndexing` indicates whether arrays
  of storage images **can** be indexed by non-uniform integer
  expressions in shader code. If this feature is not enabled, resources
  with a descriptor type of `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE` **must**
  not be indexed by non-uniform integer expressions when aggregated into
  arrays in shader code. This also indicates whether shader modules
  **can** declare the `StorageImageArrayNonUniformIndexing` capability.

- <span id="features-shaderInputAttachmentArrayNonUniformIndexing"></span>
  `shaderInputAttachmentArrayNonUniformIndexing` indicates whether
  arrays of input attachments **can** be indexed by non-uniform integer
  expressions in shader code. If this feature is not enabled, resources
  with a descriptor type of `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`
  **must** not be indexed by non-uniform integer expressions when
  aggregated into arrays in shader code. This also indicates whether
  shader modules **can** declare the
  `InputAttachmentArrayNonUniformIndexing` capability.

- <span id="features-shaderUniformTexelBufferArrayNonUniformIndexing"></span>
  `shaderUniformTexelBufferArrayNonUniformIndexing` indicates whether
  arrays of uniform texel buffers **can** be indexed by non-uniform
  integer expressions in shader code. If this feature is not enabled,
  resources with a descriptor type of
  `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` **must** not be indexed by
  non-uniform integer expressions when aggregated into arrays in shader
  code. This also indicates whether shader modules **can** declare the
  `UniformTexelBufferArrayNonUniformIndexing` capability.

- <span id="features-shaderStorageTexelBufferArrayNonUniformIndexing"></span>
  `shaderStorageTexelBufferArrayNonUniformIndexing` indicates whether
  arrays of storage texel buffers **can** be indexed by non-uniform
  integer expressions in shader code. If this feature is not enabled,
  resources with a descriptor type of
  `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` **must** not be indexed by
  non-uniform integer expressions when aggregated into arrays in shader
  code. This also indicates whether shader modules **can** declare the
  `StorageTexelBufferArrayNonUniformIndexing` capability.

- <span id="features-descriptorBindingUniformBufferUpdateAfterBind"></span>
  `descriptorBindingUniformBufferUpdateAfterBind` indicates whether the
  implementation supports updating uniform buffer descriptors after a
  set is bound. If this feature is not enabled,
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` **must** not be used
  with `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER`.

- <span id="features-descriptorBindingSampledImageUpdateAfterBind"></span>
  `descriptorBindingSampledImageUpdateAfterBind` indicates whether the
  implementation supports updating sampled image descriptors after a set
  is bound. If this feature is not enabled,
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` **must** not be used
  with `VK_DESCRIPTOR_TYPE_SAMPLER`,
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, or
  `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`.

- <span id="features-descriptorBindingStorageImageUpdateAfterBind"></span>
  `descriptorBindingStorageImageUpdateAfterBind` indicates whether the
  implementation supports updating storage image descriptors after a set
  is bound. If this feature is not enabled,
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` **must** not be used
  with `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`.

- <span id="features-descriptorBindingStorageBufferUpdateAfterBind"></span>
  `descriptorBindingStorageBufferUpdateAfterBind` indicates whether the
  implementation supports updating storage buffer descriptors after a
  set is bound. If this feature is not enabled,
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` **must** not be used
  with `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER`.

- <span id="features-descriptorBindingUniformTexelBufferUpdateAfterBind"></span>
  `descriptorBindingUniformTexelBufferUpdateAfterBind` indicates whether
  the implementation supports updating uniform texel buffer descriptors
  after a set is bound. If this feature is not enabled,
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` **must** not be used
  with `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER`.

- <span id="features-descriptorBindingStorageTexelBufferUpdateAfterBind"></span>
  `descriptorBindingStorageTexelBufferUpdateAfterBind` indicates whether
  the implementation supports updating storage texel buffer descriptors
  after a set is bound. If this feature is not enabled,
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` **must** not be used
  with `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER`.

- <span id="features-descriptorBindingUpdateUnusedWhilePending"></span>
  `descriptorBindingUpdateUnusedWhilePending` indicates whether the
  implementation supports updating descriptors while the set is in use.
  If this feature is not enabled,
  `VK_DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT` **must** not
  be used.

- <span id="features-descriptorBindingPartiallyBound"></span>
  `descriptorBindingPartiallyBound` indicates whether the implementation
  supports statically using a descriptor set binding in which some
  descriptors are not valid. If this feature is not enabled,
  `VK_DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT` **must** not be used.

- <span id="features-descriptorBindingVariableDescriptorCount"></span>
  `descriptorBindingVariableDescriptorCount` indicates whether the
  implementation supports descriptor sets with a variable-sized last
  binding. If this feature is not enabled,
  `VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT` **must** not be
  used.

- <span id="features-runtimeDescriptorArray"></span>
  `runtimeDescriptorArray` indicates whether the implementation supports
  the SPIR-V `RuntimeDescriptorArray` capability. If this feature is not
  enabled, descriptors **must** not be declared in runtime arrays.

- <span id="features-samplerFilterMinmax"></span> `samplerFilterMinmax`
  indicates whether the implementation supports a minimum set of
  required formats supporting min/max filtering as defined by the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-filterMinmaxSingleComponentFormats-minimum-requirements"
  target="_blank"
  rel="noopener"><code>filterMinmaxSingleComponentFormats</code></a>
  property minimum requirements. If this feature is not enabled, then
  [VkSamplerReductionModeCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionModeCreateInfo.html)
  **must** only use `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE`.

- <span id="features-scalarBlockLayout"></span> `scalarBlockLayout`
  indicates that the implementation supports the layout of resource
  blocks in shaders using [scalar
  alignment](#interfaces-alignment-requirements).

- <span id="features-imagelessFramebuffer"></span>
  `imagelessFramebuffer` indicates that the implementation supports
  specifying the image view for attachments at render pass begin time
  via
  [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassAttachmentBeginInfo.html).

- <span id="features-uniformBufferStandardLayout"></span>
  `uniformBufferStandardLayout` indicates that the implementation
  supports the same layouts for uniform buffers as for storage and other
  kinds of buffers. See [Standard Buffer
  Layout](#interfaces-resources-standard-layout).

- <span id="features-subgroup-extended-types"></span>
  `shaderSubgroupExtendedTypes` is a boolean specifying whether subgroup
  operations can use 8-bit integer, 16-bit integer, 64-bit integer,
  16-bit floating-point, and vectors of these types in [group
  operations](#shaders-group-operations) with [subgroup
  scope](#shaders-scope-subgroup), if the implementation supports the
  types.

- <span id="features-separateDepthStencilLayouts"></span>
  `separateDepthStencilLayouts` indicates whether the implementation
  supports a `VkImageMemoryBarrier` for a depth/stencil image with only
  one of `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`
  set, and whether `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL` can be used.

- <span id="features-hostQueryReset"></span> `hostQueryReset` indicates
  that the implementation supports resetting queries from the host with
  [vkResetQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkResetQueryPool.html).

- <span id="features-timelineSemaphore"></span> `timelineSemaphore`
  indicates whether semaphores created with a
  [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreType.html) of
  `VK_SEMAPHORE_TYPE_TIMELINE` are supported.

- <span id="features-bufferDeviceAddress"></span> `bufferDeviceAddress`
  indicates that the implementation supports accessing buffer memory in
  shaders as storage buffers via an address queried from
  [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddress.html).

- <span id="features-bufferDeviceAddressCaptureReplay"></span>
  `bufferDeviceAddressCaptureReplay` indicates that the implementation
  supports saving and reusing buffer and device addresses, e.g. for
  trace capture and replay.

- <span id="features-bufferDeviceAddressMultiDevice"></span>
  `bufferDeviceAddressMultiDevice` indicates that the implementation
  supports the `bufferDeviceAddress` , `rayTracingPipeline` and
  `rayQuery` features for logical devices created with multiple physical
  devices. If this feature is not supported, buffer and acceleration
  structure addresses **must** not be queried on a logical device
  created with more than one physical device.

- <span id="features-vulkanMemoryModel"></span> `vulkanMemoryModel`
  indicates whether shader modules **can** declare the
  `VulkanMemoryModel` capability.

- <span id="features-vulkanMemoryModelDeviceScope"></span>
  `vulkanMemoryModelDeviceScope` indicates whether the Vulkan Memory
  Model can use `Device` scope synchronization. This also indicates
  whether shader modules **can** declare the
  `VulkanMemoryModelDeviceScope` capability.

- <span id="features-vulkanMemoryModelAvailabilityVisibilityChains"></span>
  `vulkanMemoryModelAvailabilityVisibilityChains` indicates whether the
  Vulkan Memory Model can use [availability and visibility
  chains](#memory-model-availability-visibility) with more than one
  element.

- <span id="features-shaderOutputViewportIndex"></span>
  `shaderOutputViewportIndex` indicates whether the implementation
  supports the `ShaderViewportIndex` SPIR-V capability enabling
  variables decorated with the `ViewportIndex` built-in to be exported
  from mesh, vertex or tessellation evaluation shaders. If this feature
  is not enabled, the `ViewportIndex` built-in decoration **must** not
  be used on outputs in mesh, vertex or tessellation evaluation shaders.

- <span id="features-shaderOutputLayer"></span> `shaderOutputLayer`
  indicates whether the implementation supports the `ShaderLayer` SPIR-V
  capability enabling variables decorated with the `Layer` built-in to
  be exported from mesh, vertex or tessellation evaluation shaders. If
  this feature is not enabled, the `Layer` built-in decoration **must**
  not be used on outputs in mesh, vertex or tessellation evaluation
  shaders.

- <span id="features-subgroupBroadcastDynamicId"></span> If
  `subgroupBroadcastDynamicId` is `VK_TRUE`, the “Id” operand of
  `OpGroupNonUniformBroadcast` **can** be dynamically uniform within a
  subgroup, and the “Index” operand of `OpGroupNonUniformQuadBroadcast`
  **can** be dynamically uniform within the derivative group. If it is
  `VK_FALSE`, these operands **must** be constants.

If the `VkPhysicalDeviceVulkan12Features` structure is included in the
`pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceVulkan12Features` **can** also be used in the `pNext`
chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively
enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceVulkan12Features-sType-sType"
  id="VUID-VkPhysicalDeviceVulkan12Features-sType-sType"></a>
  VUID-VkPhysicalDeviceVulkan12Features-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_2_FEATURES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceVulkan12Features"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
