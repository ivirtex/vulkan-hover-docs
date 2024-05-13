# vkCmdPushDescriptorSetWithTemplateKHR(3) Manual Page

## Name

vkCmdPushDescriptorSetWithTemplateKHR - Pushes descriptor updates into a
command buffer using a descriptor update template



## <a href="#_c_specification" class="anchor"></a>C Specification

It is also possible to use a descriptor update template to specify the
push descriptors to update. To do so, call:

``` c
// Provided by VK_VERSION_1_1 with VK_KHR_push_descriptor, VK_KHR_descriptor_update_template with VK_KHR_push_descriptor
void vkCmdPushDescriptorSetWithTemplateKHR(
    VkCommandBuffer                             commandBuffer,
    VkDescriptorUpdateTemplate                  descriptorUpdateTemplate,
    VkPipelineLayout                            layout,
    uint32_t                                    set,
    const void*                                 pData);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer that the descriptors will be
  recorded in.

- `descriptorUpdateTemplate` is a descriptor update template defining
  how to interpret the descriptor information in `pData`.

- `layout` is a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) object used to
  program the bindings. It **must** be compatible with the layout used
  to create the `descriptorUpdateTemplate` handle.

- `set` is the set number of the descriptor set in the pipeline layout
  that will be updated. This **must** be the same number used to create
  the `descriptorUpdateTemplate` handle.

- `pData` is a pointer to memory containing descriptors for the
  templated update.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-vkCmdPushDescriptorSetWithTemplateKHR-commandBuffer-00366"
  id="VUID-vkCmdPushDescriptorSetWithTemplateKHR-commandBuffer-00366"></a>
  VUID-vkCmdPushDescriptorSetWithTemplateKHR-commandBuffer-00366  
  The `pipelineBindPoint` specified during the creation of the
  descriptor update template **must** be supported by the
  `commandBuffer`’s parent `VkCommandPool`’s queue family

- <a href="#VUID-vkCmdPushDescriptorSetWithTemplateKHR-pData-01686"
  id="VUID-vkCmdPushDescriptorSetWithTemplateKHR-pData-01686"></a>
  VUID-vkCmdPushDescriptorSetWithTemplateKHR-pData-01686  
  `pData` **must** be a valid pointer to a memory containing one or more
  valid instances of
  [VkDescriptorImageInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorImageInfo.html),
  [VkDescriptorBufferInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBufferInfo.html), or
  [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html) in a layout defined by
  `descriptorUpdateTemplate` when it was created with
  [vkCreateDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDescriptorUpdateTemplate.html)

- <a href="#VUID-vkCmdPushDescriptorSetWithTemplateKHR-layout-07993"
  id="VUID-vkCmdPushDescriptorSetWithTemplateKHR-layout-07993"></a>
  VUID-vkCmdPushDescriptorSetWithTemplateKHR-layout-07993  
  `layout` **must** be compatible with the layout used to create
  `descriptorUpdateTemplate`

- <a
  href="#VUID-vkCmdPushDescriptorSetWithTemplateKHR-descriptorUpdateTemplate-07994"
  id="VUID-vkCmdPushDescriptorSetWithTemplateKHR-descriptorUpdateTemplate-07994"></a>
  VUID-vkCmdPushDescriptorSetWithTemplateKHR-descriptorUpdateTemplate-07994  
  `descriptorUpdateTemplate` **must** have been created with a
  `templateType` of
  `VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS_KHR`

- <a href="#VUID-vkCmdPushDescriptorSetWithTemplateKHR-set-07995"
  id="VUID-vkCmdPushDescriptorSetWithTemplateKHR-set-07995"></a>
  VUID-vkCmdPushDescriptorSetWithTemplateKHR-set-07995  
  `set` **must** be the same value used to create
  `descriptorUpdateTemplate`

- <a href="#VUID-vkCmdPushDescriptorSetWithTemplateKHR-set-07304"
  id="VUID-vkCmdPushDescriptorSetWithTemplateKHR-set-07304"></a>
  VUID-vkCmdPushDescriptorSetWithTemplateKHR-set-07304  
  `set` **must** be less than
  [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html)::`setLayoutCount`
  provided when `layout` was created

- <a href="#VUID-vkCmdPushDescriptorSetWithTemplateKHR-set-07305"
  id="VUID-vkCmdPushDescriptorSetWithTemplateKHR-set-07305"></a>
  VUID-vkCmdPushDescriptorSetWithTemplateKHR-set-07305  
  `set` **must** be the unique set number in the pipeline layout that
  uses a descriptor set layout that was created with
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT_KHR`

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdPushDescriptorSetWithTemplateKHR-commandBuffer-parameter"
  id="VUID-vkCmdPushDescriptorSetWithTemplateKHR-commandBuffer-parameter"></a>
  VUID-vkCmdPushDescriptorSetWithTemplateKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdPushDescriptorSetWithTemplateKHR-descriptorUpdateTemplate-parameter"
  id="VUID-vkCmdPushDescriptorSetWithTemplateKHR-descriptorUpdateTemplate-parameter"></a>
  VUID-vkCmdPushDescriptorSetWithTemplateKHR-descriptorUpdateTemplate-parameter  
  `descriptorUpdateTemplate` **must** be a valid
  [VkDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplate.html) handle

- <a href="#VUID-vkCmdPushDescriptorSetWithTemplateKHR-layout-parameter"
  id="VUID-vkCmdPushDescriptorSetWithTemplateKHR-layout-parameter"></a>
  VUID-vkCmdPushDescriptorSetWithTemplateKHR-layout-parameter  
  `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html)
  handle

- <a
  href="#VUID-vkCmdPushDescriptorSetWithTemplateKHR-commandBuffer-recording"
  id="VUID-vkCmdPushDescriptorSetWithTemplateKHR-commandBuffer-recording"></a>
  VUID-vkCmdPushDescriptorSetWithTemplateKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a
  href="#VUID-vkCmdPushDescriptorSetWithTemplateKHR-commandBuffer-cmdpool"
  id="VUID-vkCmdPushDescriptorSetWithTemplateKHR-commandBuffer-cmdpool"></a>
  VUID-vkCmdPushDescriptorSetWithTemplateKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, or compute operations

- <a href="#VUID-vkCmdPushDescriptorSetWithTemplateKHR-videocoding"
  id="VUID-vkCmdPushDescriptorSetWithTemplateKHR-videocoding"></a>
  VUID-vkCmdPushDescriptorSetWithTemplateKHR-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdPushDescriptorSetWithTemplateKHR-commonparent"
  id="VUID-vkCmdPushDescriptorSetWithTemplateKHR-commonparent"></a>
  VUID-vkCmdPushDescriptorSetWithTemplateKHR-commonparent  
  Each of `commandBuffer`, `descriptorUpdateTemplate`, and `layout`
  **must** have been created, allocated, or retrieved from the same
  [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized

- Host access to the `VkCommandPool` that `commandBuffer` was allocated
  from **must** be externally synchronized

Command Properties

<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
</colgroup>
<thead>
<tr class="header">
<th class="tableblock halign-left valign-top"><a
href="#VkCommandBufferLevel">Command Buffer Levels</a></th>
<th class="tableblock halign-left valign-top"><a
href="#vkCmdBeginRenderPass">Render Pass Scope</a></th>
<th class="tableblock halign-left valign-top"><a
href="#vkCmdBeginVideoCodingKHR">Video Coding Scope</a></th>
<th class="tableblock halign-left valign-top"><a
href="#VkQueueFlagBits">Supported Queue Types</a></th>
<th class="tableblock halign-left valign-top"><a
href="#fundamentals-queueoperation-command-types">Command Type</a></th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td class="tableblock halign-left valign-top"><p>Primary<br />
Secondary</p></td>
<td class="tableblock halign-left valign-top"><p>Both</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Graphics<br />
Compute</p></td>
<td class="tableblock halign-left valign-top"><p>State</p></td>
</tr>
</tbody>
</table>

API example

``` c
struct AppDataStructure
{
    VkDescriptorImageInfo  imageInfo;          // a single image info
    // ... some more application related data
};

const VkDescriptorUpdateTemplateEntry descriptorUpdateTemplateEntries[] =
{
    // binding to a single image descriptor
    {
        .binding = 0,
        .dstArrayElement = 0,
        .descriptorCount = 1,
        .descriptorType = VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER,
        .offset = offsetof(AppDataStructure, imageInfo),
        .stride = 0     // not required if descriptorCount is 1
    }
};

// create a descriptor update template for push descriptor set updates
const VkDescriptorUpdateTemplateCreateInfo createInfo =
{
    .sType = VK_STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO,
    .pNext = NULL,
    .flags = 0,
    .descriptorUpdateEntryCount = 1,
    .pDescriptorUpdateEntries = descriptorUpdateTemplateEntries,
    .templateType = VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS_KHR,
    .descriptorSetLayout = 0,   // ignored by given templateType
    .pipelineBindPoint = VK_PIPELINE_BIND_POINT_GRAPHICS,
    .pipelineLayout = myPipelineLayout,
    .set = 0,
};

VkDescriptorUpdateTemplate myDescriptorUpdateTemplate;
myResult = vkCreateDescriptorUpdateTemplate(
    myDevice,
    &createInfo,
    NULL,
    &myDescriptorUpdateTemplate);

AppDataStructure appData;
// fill appData here or cache it in your engine
vkCmdPushDescriptorSetWithTemplateKHR(myCmdBuffer, myDescriptorUpdateTemplate, myPipelineLayout, 0,&appData);
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_descriptor_update_template](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_descriptor_update_template.html),
[VK_KHR_push_descriptor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_push_descriptor.html),
[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplate.html),
[VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdPushDescriptorSetWithTemplateKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
