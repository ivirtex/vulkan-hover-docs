# VkPhysicalDeviceDescriptorBufferPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceDescriptorBufferPropertiesEXT - Structure describing
descriptor buffer properties supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceDescriptorBufferPropertiesEXT` structure is defined
as:

``` c
// Provided by VK_EXT_descriptor_buffer
typedef struct VkPhysicalDeviceDescriptorBufferPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           combinedImageSamplerDescriptorSingleArray;
    VkBool32           bufferlessPushDescriptors;
    VkBool32           allowSamplerImageViewPostSubmitCreation;
    VkDeviceSize       descriptorBufferOffsetAlignment;
    uint32_t           maxDescriptorBufferBindings;
    uint32_t           maxResourceDescriptorBufferBindings;
    uint32_t           maxSamplerDescriptorBufferBindings;
    uint32_t           maxEmbeddedImmutableSamplerBindings;
    uint32_t           maxEmbeddedImmutableSamplers;
    size_t             bufferCaptureReplayDescriptorDataSize;
    size_t             imageCaptureReplayDescriptorDataSize;
    size_t             imageViewCaptureReplayDescriptorDataSize;
    size_t             samplerCaptureReplayDescriptorDataSize;
    size_t             accelerationStructureCaptureReplayDescriptorDataSize;
    size_t             samplerDescriptorSize;
    size_t             combinedImageSamplerDescriptorSize;
    size_t             sampledImageDescriptorSize;
    size_t             storageImageDescriptorSize;
    size_t             uniformTexelBufferDescriptorSize;
    size_t             robustUniformTexelBufferDescriptorSize;
    size_t             storageTexelBufferDescriptorSize;
    size_t             robustStorageTexelBufferDescriptorSize;
    size_t             uniformBufferDescriptorSize;
    size_t             robustUniformBufferDescriptorSize;
    size_t             storageBufferDescriptorSize;
    size_t             robustStorageBufferDescriptorSize;
    size_t             inputAttachmentDescriptorSize;
    size_t             accelerationStructureDescriptorSize;
    VkDeviceSize       maxSamplerDescriptorBufferRange;
    VkDeviceSize       maxResourceDescriptorBufferRange;
    VkDeviceSize       samplerDescriptorBufferAddressSpaceSize;
    VkDeviceSize       resourceDescriptorBufferAddressSpaceSize;
    VkDeviceSize       descriptorBufferAddressSpaceSize;
} VkPhysicalDeviceDescriptorBufferPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-combinedImageSamplerDescriptorSingleArray"></span>
  `combinedImageSamplerDescriptorSingleArray` indicates that the
  implementation does not require an array of
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` descriptors to be written
  into a descriptor buffer as an array of image descriptors, immediately
  followed by an array of sampler descriptors.

- <span id="limits-bufferlessPushDescriptors"></span>
  `bufferlessPushDescriptors` indicates that the implementation does not
  require a buffer created with
  `VK_BUFFER_USAGE_PUSH_DESCRIPTORS_DESCRIPTOR_BUFFER_BIT_EXT` to be
  bound when using push descriptors.

- <span id="limits-allowSamplerImageViewPostSubmitCreation"></span>
  `allowSamplerImageViewPostSubmitCreation` indicates that the
  implementation does not restrict when the [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html)
  or [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) objects used to retrieve descriptor
  data **can** be created in relation to command buffer submission. If
  this value is `VK_FALSE`, then the application **must** create any
  [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) or [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) objects
  whose descriptor data is accessed during the execution of a command
  buffer, before the [vkQueueSubmit](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueSubmit.html) , or
  [vkQueueSubmit2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueSubmit2.html), call that submits that command
  buffer.

- <span id="limits-descriptorBufferOffsetAlignment"></span>
  `descriptorBufferOffsetAlignment` indicates the **required** alignment
  in bytes when setting offsets into the descriptor buffer.

- <span id="limits-maxDescriptorBufferBindings"></span>
  `maxDescriptorBufferBindings` indicates the maximum sum total number
  of descriptor buffers and embedded immutable sampler sets that **can**
  be bound.

- <span id="limits-maxResourceDescriptorBufferBindings"></span>
  `maxResourceDescriptorBufferBindings` indicates the maximum number of
  resource descriptor buffers that **can** be bound.

- <span id="limits-maxSamplerDescriptorBufferBindings"></span>
  `maxSamplerDescriptorBufferBindings` indicates the maximum number of
  sampler descriptor buffers that **can** be bound.

- <span id="limits-maxEmbeddedImmutableSamplerBindings"></span>
  `maxEmbeddedImmutableSamplerBindings` indicates the maximum number of
  embedded immutable sampler sets that **can** be bound.

- <span id="limits-maxEmbeddedImmutableSamplers"></span>
  `maxEmbeddedImmutableSamplers` indicates the maximum number of unique
  immutable samplers in descriptor set layouts created with
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_EMBEDDED_IMMUTABLE_SAMPLERS_BIT_EXT`,
  and pipeline layouts created from them, which **can** simultaneously
  exist on a device.

- <span id="limits-bufferCaptureReplayDescriptorDataSize"></span>
  `bufferCaptureReplayDescriptorDataSize` indicates the maximum size in
  bytes of the opaque data used for capture and replay with buffers.

- <span id="limits-imageCaptureReplayDescriptorDataSize"></span>
  `imageCaptureReplayDescriptorDataSize` indicates the maximum size in
  bytes of the opaque data used for capture and replay with images.

- <span id="limits-imageViewCaptureReplayDescriptorDataSize"></span>
  `imageViewCaptureReplayDescriptorDataSize` indicates the maximum size
  in bytes of the opaque data used for capture and replay with image
  views.

- <span id="limits-samplerCaptureReplayDescriptorDataSize"></span>
  `samplerCaptureReplayDescriptorDataSize` indicates the maximum size in
  bytes of the opaque data used for capture and replay with samplers.

- <span id="limits-accelerationStructureCaptureReplayDescriptorDataSize"></span>
  `accelerationStructureCaptureReplayDescriptorDataSize` indicates the
  maximum size in bytes of the opaque data used for capture and replay
  with acceleration structures.

- <span id="limits-samplerDescriptorSize"></span>
  `samplerDescriptorSize` indicates the size in bytes of a
  `VK_DESCRIPTOR_TYPE_SAMPLER` descriptor.

- <span id="limits-combinedImageSamplerDescriptorSize"></span>
  `combinedImageSamplerDescriptorSize` indicates the size in bytes of a
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` descriptor.

- <span id="limits-sampledImageDescriptorSize"></span>
  `sampledImageDescriptorSize` indicates the size in bytes of a
  `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE` descriptor.

- <span id="limits-storageImageDescriptorSize"></span>
  `storageImageDescriptorSize` indicates the size in bytes of a
  `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE` descriptor.

- <span id="limits-uniformTexelBufferDescriptorSize"></span>
  `uniformTexelBufferDescriptorSize` indicates the size in bytes of a
  `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` descriptor if the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustBufferAccess"
  target="_blank" rel="noopener"><code>robustBufferAccess</code></a>
  feature is not enabled.

- <span id="limits-robustUniformTexelBufferDescriptorSize"></span>
  `robustUniformTexelBufferDescriptorSize` indicates the size in bytes
  of a `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` descriptor if the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustBufferAccess"
  target="_blank" rel="noopener"><code>robustBufferAccess</code></a>
  feature is enabled.

- <span id="limits-storageTexelBufferDescriptorSize"></span>
  `storageTexelBufferDescriptorSize` indicates the size in bytes of a
  `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` descriptor if the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustBufferAccess"
  target="_blank" rel="noopener"><code>robustBufferAccess</code></a>
  feature is not enabled.

- <span id="limits-robustStorageTexelBufferDescriptorSize"></span>
  `robustStorageTexelBufferDescriptorSize` indicates the size in bytes
  of a `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` descriptor if the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustBufferAccess"
  target="_blank" rel="noopener"><code>robustBufferAccess</code></a>
  feature is enabled.

- <span id="limits-uniformBufferDescriptorSize"></span>
  `uniformBufferDescriptorSize` indicates the size in bytes of a
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER` descriptor.

- <span id="limits-robustUniformBufferDescriptorSize"></span>
  `robustUniformBufferDescriptorSize` indicates the size in bytes of a
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER` descriptor if the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustBufferAccess"
  target="_blank" rel="noopener"><code>robustBufferAccess</code></a>
  feature is enabled.

- <span id="limits-storageBufferDescriptorSize"></span>
  `storageBufferDescriptorSize` indicates the size in bytes of a
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER` descriptor.

- <span id="limits-robustStorageBufferDescriptorSize"></span>
  `robustStorageBufferDescriptorSize` indicates the size in bytes of a
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER` descriptor if the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustBufferAccess"
  target="_blank" rel="noopener"><code>robustBufferAccess</code></a>
  feature is enabled.

- <span id="limits-inputAttachmentDescriptorSize"></span>
  `inputAttachmentDescriptorSize` indicates the size in bytes of a
  `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT` descriptor.

- <span id="limits-accelerationStructureDescriptorSize"></span>
  `accelerationStructureDescriptorSize` indicates the size in bytes of a
  `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR` or
  `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_NV` descriptor.

- <span id="limits-maxSamplerDescriptorBufferRange"></span>
  `maxSamplerDescriptorBufferRange` indicates the maximum range in bytes
  from the address of a sampler descriptor buffer binding that is
  accessible to a shader.

- <span id="limits-maxResourceDescriptorBufferRange"></span>
  `maxResourceDescriptorBufferRange` indicates the maximum range in
  bytes from the address of a resource descriptor buffer binding that is
  accessible to a shader.

- <span id="limits-samplerDescriptorBufferAddressSpaceSize"></span>
  `samplerDescriptorBufferAddressSpaceSize` indicates the total size in
  bytes of the address space available for descriptor buffers created
  with `VK_BUFFER_USAGE_SAMPLER_DESCRIPTOR_BUFFER_BIT_EXT`.

- <span id="limits-resourceDescriptorBufferAddressSpaceSize"></span>
  `resourceDescriptorBufferAddressSpaceSize` indicates the total size in
  bytes of the address space available for descriptor buffers created
  with `VK_BUFFER_USAGE_RESOURCE_DESCRIPTOR_BUFFER_BIT_EXT`.

- <span id="limits-descriptorBufferAddressSpaceSize"></span>
  `descriptorBufferAddressSpaceSize` indicates the total size in bytes
  of the address space available for descriptor buffers created with
  both `VK_BUFFER_USAGE_SAMPLER_DESCRIPTOR_BUFFER_BIT_EXT` and
  `VK_BUFFER_USAGE_RESOURCE_DESCRIPTOR_BUFFER_BIT_EXT`.

## <a href="#_description" class="anchor"></a>Description

A descriptor binding with type `VK_DESCRIPTOR_TYPE_MUTABLE_VALVE` has a
descriptor size which is implied by the descriptor types included in the
[VkMutableDescriptorTypeCreateInfoVALVE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMutableDescriptorTypeCreateInfoVALVE.html)::`pDescriptorTypes`
list. The descriptor size is equal to the maximum size of any descriptor
type included in the `pDescriptorTypes` list.

As there is no way to request robust and non-robust descriptors
separately, or specify robust/non-robust descriptors in the set layout,
if <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustBufferAccess"
target="_blank" rel="noopener"><code>robustBufferAccess</code></a> is
enabled then robust descriptors are always used.

If the `VkPhysicalDeviceDescriptorBufferPropertiesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceDescriptorBufferPropertiesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceDescriptorBufferPropertiesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceDescriptorBufferPropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_BUFFER_PROPERTIES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_descriptor_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_buffer.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceDescriptorBufferPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
