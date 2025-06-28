# VkPhysicalDeviceDescriptorBufferPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceDescriptorBufferPropertiesEXT - Structure describing descriptor buffer properties supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceDescriptorBufferPropertiesEXT` structure is defined as:

```c++
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

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`combinedImageSamplerDescriptorSingleArray` indicates that the implementation does not require an array of `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` descriptors to be written into a descriptor buffer as an array of image descriptors, immediately followed by an array of sampler descriptors.
- []()`bufferlessPushDescriptors` indicates that the implementation does not require a buffer created with `VK_BUFFER_USAGE_PUSH_DESCRIPTORS_DESCRIPTOR_BUFFER_BIT_EXT` to be bound when using push descriptors.
- []()`allowSamplerImageViewPostSubmitCreation` indicates that the implementation does not restrict when the [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) or [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) objects used to retrieve descriptor data **can** be created in relation to command buffer submission. If this value is `VK_FALSE`, then the application **must** create any [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) or [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) objects whose descriptor data is accessed during the execution of a command buffer, before the [vkQueueSubmit](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueSubmit.html) , or [vkQueueSubmit2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueSubmit2.html), call that submits that command buffer.
- []()`descriptorBufferOffsetAlignment` indicates the **required** alignment in bytes when setting offsets into the descriptor buffer.
- []()`maxDescriptorBufferBindings` indicates the maximum number of descriptor buffer bindings.
- []()`maxResourceDescriptorBufferBindings` indicates the maximum number of descriptor buffer bindings with `VK_BUFFER_USAGE_RESOURCE_DESCRIPTOR_BUFFER_BIT_EXT` that **can** be used.
- []()`maxSamplerDescriptorBufferBindings` indicates the maximum number of descriptor buffer bindings with `VK_BUFFER_USAGE_SAMPLER_DESCRIPTOR_BUFFER_BIT_EXT` that **can** be used.
- []()`maxEmbeddedImmutableSamplerBindings` indicates the maximum number of embedded immutable sampler sets that **can** be bound.
- []()`maxEmbeddedImmutableSamplers` indicates the maximum number of unique immutable samplers in descriptor set layouts created with `VK_DESCRIPTOR_SET_LAYOUT_CREATE_EMBEDDED_IMMUTABLE_SAMPLERS_BIT_EXT`, and pipeline layouts created from them, which **can** simultaneously exist on a device.
- []()`bufferCaptureReplayDescriptorDataSize` indicates the maximum size in bytes of the opaque data used for capture and replay with buffers.
- []()`imageCaptureReplayDescriptorDataSize` indicates the maximum size in bytes of the opaque data used for capture and replay with images.
- []()`imageViewCaptureReplayDescriptorDataSize` indicates the maximum size in bytes of the opaque data used for capture and replay with image views.
- []()`samplerCaptureReplayDescriptorDataSize` indicates the maximum size in bytes of the opaque data used for capture and replay with samplers.
- []()`accelerationStructureCaptureReplayDescriptorDataSize` indicates the maximum size in bytes of the opaque data used for capture and replay with acceleration structures.
- []()`samplerDescriptorSize` indicates the size in bytes of a `VK_DESCRIPTOR_TYPE_SAMPLER` descriptor.
- []()`combinedImageSamplerDescriptorSize` indicates the size in bytes of a `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` descriptor.
- []()`sampledImageDescriptorSize` indicates the size in bytes of a `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE` descriptor.
- []()`storageImageDescriptorSize` indicates the size in bytes of a `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE` descriptor.
- []()`uniformTexelBufferDescriptorSize` indicates the size in bytes of a `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` descriptor if the [`robustBufferAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-robustBufferAccess) feature is not enabled.
- []()`robustUniformTexelBufferDescriptorSize` indicates the size in bytes of a `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` descriptor if the [`robustBufferAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-robustBufferAccess) feature is enabled.
- []()`storageTexelBufferDescriptorSize` indicates the size in bytes of a `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` descriptor if the [`robustBufferAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-robustBufferAccess) feature is not enabled.
- []()`robustStorageTexelBufferDescriptorSize` indicates the size in bytes of a `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` descriptor if the [`robustBufferAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-robustBufferAccess) feature is enabled.
- []()`uniformBufferDescriptorSize` indicates the size in bytes of a `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER` descriptor.
- []()`robustUniformBufferDescriptorSize` indicates the size in bytes of a `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER` descriptor if the [`robustBufferAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-robustBufferAccess) feature is enabled.
- []()`storageBufferDescriptorSize` indicates the size in bytes of a `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER` descriptor.
- []()`robustStorageBufferDescriptorSize` indicates the size in bytes of a `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER` descriptor if the [`robustBufferAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-robustBufferAccess) feature is enabled.
- []()`inputAttachmentDescriptorSize` indicates the size in bytes of a `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT` descriptor.
- []()`accelerationStructureDescriptorSize` indicates the size in bytes of a `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR` or `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_NV` descriptor.
- []()`maxSamplerDescriptorBufferRange` indicates the maximum range in bytes from the address of a sampler descriptor buffer binding that is accessible to a shader.
- []()`maxResourceDescriptorBufferRange` indicates the maximum range in bytes from the address of a resource descriptor buffer binding that is accessible to a shader.
- []()`samplerDescriptorBufferAddressSpaceSize` indicates the total size in bytes of the address space available for descriptor buffers created with `VK_BUFFER_USAGE_SAMPLER_DESCRIPTOR_BUFFER_BIT_EXT`.
- []()`resourceDescriptorBufferAddressSpaceSize` indicates the total size in bytes of the address space available for descriptor buffers created with `VK_BUFFER_USAGE_RESOURCE_DESCRIPTOR_BUFFER_BIT_EXT`.
- []()`descriptorBufferAddressSpaceSize` indicates the total size in bytes of the address space available for descriptor buffers created with both `VK_BUFFER_USAGE_SAMPLER_DESCRIPTOR_BUFFER_BIT_EXT` and `VK_BUFFER_USAGE_RESOURCE_DESCRIPTOR_BUFFER_BIT_EXT`.

## [](#_description)Description

A descriptor binding with type `VK_DESCRIPTOR_TYPE_MUTABLE_VALVE` has a descriptor size which is implied by the descriptor types included in the [VkMutableDescriptorTypeCreateInfoVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMutableDescriptorTypeCreateInfoVALVE.html)::`pDescriptorTypes` list. The descriptor size is equal to the maximum size of any descriptor type included in the `pDescriptorTypes` list.

As there is no way to request robust and non-robust descriptors separately, or specify robust/non-robust descriptors in the set layout, if the [`robustBufferAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-robustBufferAccess) feature is enabled then robust descriptors are always used.

If the `VkPhysicalDeviceDescriptorBufferPropertiesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceDescriptorBufferPropertiesEXT-sType-sType)VUID-VkPhysicalDeviceDescriptorBufferPropertiesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DESCRIPTOR_BUFFER_PROPERTIES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_descriptor\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_buffer.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceDescriptorBufferPropertiesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0