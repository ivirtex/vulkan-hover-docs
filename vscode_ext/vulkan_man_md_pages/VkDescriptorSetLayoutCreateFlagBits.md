# VkDescriptorSetLayoutCreateFlagBits(3) Manual Page

## Name

VkDescriptorSetLayoutCreateFlagBits - Bitmask specifying descriptor set
layout properties



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateInfo.html)::`flags`,
specifying options for descriptor set layout, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkDescriptorSetLayoutCreateFlagBits {
  // Provided by VK_VERSION_1_2
    VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT = 0x00000002,
  // Provided by VK_KHR_push_descriptor
    VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT_KHR = 0x00000001,
  // Provided by VK_EXT_descriptor_buffer
    VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT = 0x00000010,
  // Provided by VK_EXT_descriptor_buffer
    VK_DESCRIPTOR_SET_LAYOUT_CREATE_EMBEDDED_IMMUTABLE_SAMPLERS_BIT_EXT = 0x00000020,
  // Provided by VK_NV_device_generated_commands_compute
    VK_DESCRIPTOR_SET_LAYOUT_CREATE_INDIRECT_BINDABLE_BIT_NV = 0x00000080,
  // Provided by VK_EXT_mutable_descriptor_type
    VK_DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_EXT = 0x00000004,
  // Provided by VK_NV_per_stage_descriptor_set
    VK_DESCRIPTOR_SET_LAYOUT_CREATE_PER_STAGE_BIT_NV = 0x00000040,
  // Provided by VK_EXT_descriptor_indexing
    VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT_EXT = VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT,
  // Provided by VK_VALVE_mutable_descriptor_type
    VK_DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_VALVE = VK_DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_EXT,
} VkDescriptorSetLayoutCreateFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT_KHR` specifies
  that descriptor sets **must** not be allocated using this layout, and
  descriptors are instead pushed by
  [vkCmdPushDescriptorSetKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPushDescriptorSetKHR.html).

- `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` specifies
  that descriptor sets using this layout **must** be allocated from a
  descriptor pool created with the
  `VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT` bit set. Descriptor
  set layouts created with this bit set have alternate limits for the
  maximum number of descriptors per-stage and per-pipeline layout. The
  non-UpdateAfterBind limits only count descriptors in sets created
  without this flag. The UpdateAfterBind limits count all descriptors,
  but the limits **may** be higher than the non-UpdateAfterBind limits.

- `VK_DESCRIPTOR_SET_LAYOUT_CREATE_INDIRECT_BINDABLE_BIT_NV` specifies
  that descriptor sets using this layout allows them to be bound with
  compute pipelines that are created with
  `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV` flag set to be used in
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#device-generated-commands"
  target="_blank" rel="noopener">Device-Generated Commands</a>.

- `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT` specifies
  that this layout **must** only be used with descriptor buffers.

- `VK_DESCRIPTOR_SET_LAYOUT_CREATE_EMBEDDED_IMMUTABLE_SAMPLERS_BIT_EXT`
  specifies that this is a layout only containing immutable samplers
  that **can** be bound by
  [vkCmdBindDescriptorBufferEmbeddedSamplersEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindDescriptorBufferEmbeddedSamplersEXT.html).
  Unlike normal immutable samplers, embedded immutable samplers do not
  require the application to provide them in a descriptor buffer.

- `VK_DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_EXT` specifies
  that descriptor sets using this layout **must** be allocated from a
  descriptor pool created with the
  `VK_DESCRIPTOR_POOL_CREATE_HOST_ONLY_BIT_EXT` bit set. Descriptor set
  layouts created with this bit have no expressible limit for maximum
  number of descriptors per-stage. Host descriptor sets are limited only
  by available host memory, but **may** be limited for implementation
  specific reasons. Implementations **may** limit the number of
  supported descriptors to UpdateAfterBind limits or non-UpdateAfterBind
  limits, whichever is larger.

- `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PER_STAGE_BIT_NV` specifies that
  binding numbers in descriptor sets using this layout **may** represent
  different resources and/or types of resources in each stage.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDescriptorSetLayoutCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDescriptorSetLayoutCreateFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
