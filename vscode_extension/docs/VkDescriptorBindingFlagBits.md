# VkDescriptorBindingFlagBits(3) Manual Page

## Name

VkDescriptorBindingFlagBits - Bitmask specifying descriptor set layout binding properties



## [](#_c_specification)C Specification

Bits which **can** be set in each element of [VkDescriptorSetLayoutBindingFlagsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutBindingFlagsCreateInfo.html)::`pBindingFlags`, specifying options for the corresponding descriptor set layout binding, are:

```c++
// Provided by VK_VERSION_1_2
typedef enum VkDescriptorBindingFlagBits {
    VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT = 0x00000001,
    VK_DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT = 0x00000002,
    VK_DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT = 0x00000004,
    VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT = 0x00000008,
  // Provided by VK_EXT_descriptor_indexing
    VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT_EXT = VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT,
  // Provided by VK_EXT_descriptor_indexing
    VK_DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT_EXT = VK_DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT,
  // Provided by VK_EXT_descriptor_indexing
    VK_DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT_EXT = VK_DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT,
  // Provided by VK_EXT_descriptor_indexing
    VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT_EXT = VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT,
} VkDescriptorBindingFlagBits;
```

or the equivalent

```c++
// Provided by VK_EXT_descriptor_indexing
typedef VkDescriptorBindingFlagBits VkDescriptorBindingFlagBitsEXT;
```

## [](#_description)Description

- `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` specifies that if descriptors in this binding are updated between when the descriptor set is bound in a command buffer and when that command buffer is submitted to a queue, then the submission will use the most recently set descriptors for this binding and the updates do not invalidate the command buffer. Descriptor bindings created with this flag are also partially exempt from the external synchronization requirement in [vkUpdateDescriptorSetWithTemplateKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUpdateDescriptorSetWithTemplateKHR.html) and [vkUpdateDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUpdateDescriptorSets.html). Multiple descriptors with this flag set **can** be updated concurrently in different threads, though the same descriptor **must** not be updated concurrently by two threads. Descriptors with this flag set **can** be updated concurrently with the set being bound to a command buffer in another thread, but not concurrently with the set being reset or freed.
- `VK_DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT` specifies that descriptors in this binding that are not *dynamically used* need not contain valid descriptors at the time the descriptors are consumed. A descriptor is dynamically used if any shader invocation executes an instruction that performs any memory access using the descriptor. If a descriptor is not dynamically used, any resource referenced by the descriptor is not considered to be referenced during command execution.
- `VK_DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT` specifies that descriptors in this binding **can** be updated after a command buffer has bound this descriptor set, or while a command buffer that uses this descriptor set is pending execution, as long as the descriptors that are updated are not used by those command buffers. Descriptor bindings created with this flag are also partially exempt from the external synchronization requirement in [vkUpdateDescriptorSetWithTemplateKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUpdateDescriptorSetWithTemplateKHR.html) and [vkUpdateDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUpdateDescriptorSets.html) in the same way as for `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`. If `VK_DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT` is also set, then descriptors **can** be updated as long as they are not dynamically used by any shader invocations. If `VK_DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT` is not set, then descriptors **can** be updated as long as they are not statically used by any shader invocations.
- `VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT` specifies that this is a *variable-sized descriptor binding* whose size will be specified when a descriptor set is allocated using this layout. The value of `descriptorCount` is treated as an upper bound on the size of the binding. This **must** only be used for the last binding in the descriptor set layout (i.e. the binding with the largest value of `binding`). For the purposes of counting against limits such as `maxDescriptorSet`* and `maxPerStageDescriptor`\*, the full value of `descriptorCount` is counted, except for descriptor bindings with a descriptor type of `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`, when [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateInfo.html)::`flags` does not contain `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`. In this case, `descriptorCount` specifies the upper bound on the byte size of the binding; thus it counts against the [`maxInlineUniformBlockSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxInlineUniformBlockSize) and [`maxInlineUniformTotalSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxInlineUniformTotalSize) limits instead.

Note

Note that while `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` and `VK_DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT` both involve updates to descriptor sets after they are bound, `VK_DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT` is a weaker requirement since it is only about descriptors that are not used, whereas `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` requires the implementation to observe updates to descriptors that are used.

## [](#_see_also)See Also

[VK\_EXT\_descriptor\_indexing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_descriptor_indexing.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkDescriptorBindingFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBindingFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDescriptorBindingFlagBits)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0