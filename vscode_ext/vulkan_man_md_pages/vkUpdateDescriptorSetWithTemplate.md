# vkUpdateDescriptorSetWithTemplate(3) Manual Page

## Name

vkUpdateDescriptorSetWithTemplate - Update the contents of a descriptor
set object using an update template



## <a href="#_c_specification" class="anchor"></a>C Specification

Once a `VkDescriptorUpdateTemplate` has been created, descriptor sets
**can** be updated by calling:

``` c
// Provided by VK_VERSION_1_1
void vkUpdateDescriptorSetWithTemplate(
    VkDevice                                    device,
    VkDescriptorSet                             descriptorSet,
    VkDescriptorUpdateTemplate                  descriptorUpdateTemplate,
    const void*                                 pData);
```

or the equivalent command

``` c
// Provided by VK_KHR_descriptor_update_template
void vkUpdateDescriptorSetWithTemplateKHR(
    VkDevice                                    device,
    VkDescriptorSet                             descriptorSet,
    VkDescriptorUpdateTemplate                  descriptorUpdateTemplate,
    const void*                                 pData);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that updates the descriptor set.

- `descriptorSet` is the descriptor set to update

- `descriptorUpdateTemplate` is a
  [VkDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplate.html) object
  specifying the update mapping between `pData` and the descriptor set
  to update.

- `pData` is a pointer to memory containing one or more
  [VkDescriptorImageInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorImageInfo.html),
  [VkDescriptorBufferInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBufferInfo.html), or
  [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html) structures or
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html) or
  [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureNV.html) handles
  used to write the descriptors.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkUpdateDescriptorSetWithTemplate-pData-01685"
  id="VUID-vkUpdateDescriptorSetWithTemplate-pData-01685"></a>
  VUID-vkUpdateDescriptorSetWithTemplate-pData-01685  
  `pData` **must** be a valid pointer to a memory containing one or more
  valid instances of
  [VkDescriptorImageInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorImageInfo.html),
  [VkDescriptorBufferInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBufferInfo.html), or
  [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html) in a layout defined by
  `descriptorUpdateTemplate` when it was created with
  [vkCreateDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDescriptorUpdateTemplate.html)

- <a href="#VUID-vkUpdateDescriptorSetWithTemplate-descriptorSet-06995"
  id="VUID-vkUpdateDescriptorSetWithTemplate-descriptorSet-06995"></a>
  VUID-vkUpdateDescriptorSetWithTemplate-descriptorSet-06995  
  Host access to `descriptorSet` **must** be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fundamentals-threadingbehavior"
  target="_blank" rel="noopener">externally synchronized</a> unless
  explicitly denoted otherwise for specific flags

Valid Usage (Implicit)

- <a href="#VUID-vkUpdateDescriptorSetWithTemplate-device-parameter"
  id="VUID-vkUpdateDescriptorSetWithTemplate-device-parameter"></a>
  VUID-vkUpdateDescriptorSetWithTemplate-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkUpdateDescriptorSetWithTemplate-descriptorSet-parameter"
  id="VUID-vkUpdateDescriptorSetWithTemplate-descriptorSet-parameter"></a>
  VUID-vkUpdateDescriptorSetWithTemplate-descriptorSet-parameter  
  `descriptorSet` **must** be a valid
  [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSet.html) handle

- <a
  href="#VUID-vkUpdateDescriptorSetWithTemplate-descriptorUpdateTemplate-parameter"
  id="VUID-vkUpdateDescriptorSetWithTemplate-descriptorUpdateTemplate-parameter"></a>
  VUID-vkUpdateDescriptorSetWithTemplate-descriptorUpdateTemplate-parameter  
  `descriptorUpdateTemplate` **must** be a valid
  [VkDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplate.html) handle

- <a href="#VUID-vkUpdateDescriptorSetWithTemplate-descriptorSet-parent"
  id="VUID-vkUpdateDescriptorSetWithTemplate-descriptorSet-parent"></a>
  VUID-vkUpdateDescriptorSetWithTemplate-descriptorSet-parent  
  `descriptorSet` **must** have been created, allocated, or retrieved
  from `device`

- <a
  href="#VUID-vkUpdateDescriptorSetWithTemplate-descriptorUpdateTemplate-parent"
  id="VUID-vkUpdateDescriptorSetWithTemplate-descriptorUpdateTemplate-parent"></a>
  VUID-vkUpdateDescriptorSetWithTemplate-descriptorUpdateTemplate-parent  
  `descriptorUpdateTemplate` **must** have been created, allocated, or
  retrieved from `device`

API example

``` c
struct AppBufferView {
    VkBufferView bufferView;
    uint32_t     applicationRelatedInformation;
};

struct AppDataStructure
{
    VkDescriptorImageInfo  imageInfo;          // a single image info
    VkDescriptorBufferInfo bufferInfoArray[3]; // 3 buffer infos in an array
    AppBufferView          bufferView[2];      // An application-defined structure containing a bufferView
    // ... some more application-related data
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
        .stride = 0         // stride not required if descriptorCount is 1
    },

    // binding to an array of buffer descriptors
    {
        .binding = 1,
        .dstArrayElement = 0,
        .descriptorCount = 3,
        .descriptorType = VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER,
        .offset = offsetof(AppDataStructure, bufferInfoArray),
        .stride = sizeof(VkDescriptorBufferInfo)    // descriptor buffer infos are compact
    },

    // binding to an array of buffer views
    {
        .binding = 2,
        .dstArrayElement = 0,
        .descriptorCount = 2,
        .descriptorType = VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER,
        .offset = offsetof(AppDataStructure, bufferView) +
                  offsetof(AppBufferView, bufferView),
        .stride = sizeof(AppBufferView)             // bufferViews do not have to be compact
    },
};

// create a descriptor update template for descriptor set updates
const VkDescriptorUpdateTemplateCreateInfo createInfo =
{
    .sType = VK_STRUCTURE_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_CREATE_INFO,
    .pNext = NULL,
    .flags = 0,
    .descriptorUpdateEntryCount = 3,
    .pDescriptorUpdateEntries = descriptorUpdateTemplateEntries,
    .templateType = VK_DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET,
    .descriptorSetLayout = myLayout,
    .pipelineBindPoint = 0,     // ignored by given templateType
    .pipelineLayout = 0,        // ignored by given templateType
    .set = 0,                   // ignored by given templateType
};

VkDescriptorUpdateTemplate myDescriptorUpdateTemplate;
myResult = vkCreateDescriptorUpdateTemplate(
    myDevice,
    &createInfo,
    NULL,
    &myDescriptorUpdateTemplate);

AppDataStructure appData;

// fill appData here or cache it in your engine
vkUpdateDescriptorSetWithTemplate(myDevice, myDescriptorSet, myDescriptorUpdateTemplate, &appData);
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSet.html),
[VkDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplate.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkUpdateDescriptorSetWithTemplate"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
