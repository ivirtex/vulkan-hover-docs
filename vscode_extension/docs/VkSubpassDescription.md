# VkSubpassDescription(3) Manual Page

## Name

VkSubpassDescription - Structure specifying a subpass description



## [](#_c_specification)C Specification

The `VkSubpassDescription` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.2](#versions-1.2). See [Deprecated Functionality](#deprecation-renderpass2) for more information.

```c++
// Provided by VK_VERSION_1_0
typedef struct VkSubpassDescription {
    VkSubpassDescriptionFlags       flags;
    VkPipelineBindPoint             pipelineBindPoint;
    uint32_t                        inputAttachmentCount;
    const VkAttachmentReference*    pInputAttachments;
    uint32_t                        colorAttachmentCount;
    const VkAttachmentReference*    pColorAttachments;
    const VkAttachmentReference*    pResolveAttachments;
    const VkAttachmentReference*    pDepthStencilAttachment;
    uint32_t                        preserveAttachmentCount;
    const uint32_t*                 pPreserveAttachments;
} VkSubpassDescription;
```

## [](#_members)Members

- `flags` is a bitmask of [VkSubpassDescriptionFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescriptionFlagBits.html) specifying usage of the subpass.
- `pipelineBindPoint` is a [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html) value specifying the pipeline type supported for this subpass.
- `inputAttachmentCount` is the number of input attachments.
- `pInputAttachments` is a pointer to an array of [VkAttachmentReference](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference.html) structures defining the input attachments for this subpass and their layouts.
- `colorAttachmentCount` is the number of color attachments.
- `pColorAttachments` is a pointer to an array of `colorAttachmentCount` [VkAttachmentReference](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference.html) structures defining the color attachments for this subpass and their layouts.
- `pResolveAttachments` is `NULL` or a pointer to an array of `colorAttachmentCount` [VkAttachmentReference](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference.html) structures defining the resolve attachments for this subpass and their layouts.
- `pDepthStencilAttachment` is a pointer to a [VkAttachmentReference](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference.html) structure specifying the depth/stencil attachment for this subpass and its layout.
- `preserveAttachmentCount` is the number of preserved attachments.
- `pPreserveAttachments` is a pointer to an array of `preserveAttachmentCount` render pass attachment indices identifying attachments that are not used by this subpass, but whose contents **must** be preserved throughout the subpass.

## [](#_description)Description

Each element of the `pInputAttachments` array corresponds to an input attachment index in a fragment shader, i.e. if a shader declares an image variable decorated with a `InputAttachmentIndex` value of **X**, then it uses the attachment provided in `pInputAttachments`\[**X**]. Input attachments **must** also be bound to the pipeline in a descriptor set. If the `attachment` member of any element of `pInputAttachments` is `VK_ATTACHMENT_UNUSED`, the application **must** not read from the corresponding input attachment index. Fragment shaders **can** use subpass input variables to access the contents of an input attachment at the fragment’s (xf,yf) framebuffer coordinates and layer. Input attachments **must** not be used by any subpasses within a render pass that enables [render pass transform](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vertexpostproc-renderpass-transform).

Each element of the `pColorAttachments` array corresponds to an output location in the shader, i.e. if the shader declares an output variable decorated with a `Location` value of **X**, then it uses the attachment provided in `pColorAttachments`\[**X**]. If the `attachment` member of any element of `pColorAttachments` is `VK_ATTACHMENT_UNUSED`, or if [Color Write Enable](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#framebuffer-color-write-enable) has been disabled for the corresponding attachment index, then writes to the corresponding location by a fragment shader are discarded.

If `flags` does not include `VK_SUBPASS_DESCRIPTION_SHADER_RESOLVE_BIT_QCOM`, and if `pResolveAttachments` is not `NULL`, each of its elements corresponds to a color attachment (the element in `pColorAttachments` at the same index), and a [multisample resolve operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-resolve-operations) is defined for each attachment unless the resolve attachment index is `VK_ATTACHMENT_UNUSED`.

Similarly, if `flags` does not include `VK_SUBPASS_DESCRIPTION_SHADER_RESOLVE_BIT_QCOM`, and [VkSubpassDescriptionDepthStencilResolve](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescriptionDepthStencilResolve.html)::`pDepthStencilResolveAttachment` is not `NULL` and does not have the value `VK_ATTACHMENT_UNUSED`, it corresponds to the depth/stencil attachment in `pDepthStencilAttachment`, and [multisample resolve operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-resolve-operations) for depth and stencil are defined by [VkSubpassDescriptionDepthStencilResolve](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescriptionDepthStencilResolve.html)::`depthResolveMode` and [VkSubpassDescriptionDepthStencilResolve](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescriptionDepthStencilResolve.html)::`stencilResolveMode`, respectively. If [VkSubpassDescriptionDepthStencilResolve](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescriptionDepthStencilResolve.html)::`depthResolveMode` is `VK_RESOLVE_MODE_NONE` or the `pDepthStencilResolveAttachment` does not have a depth aspect, no resolve operation is performed for the depth attachment. If [VkSubpassDescriptionDepthStencilResolve](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescriptionDepthStencilResolve.html)::`stencilResolveMode` is `VK_RESOLVE_MODE_NONE` or the `pDepthStencilResolveAttachment` does not have a stencil aspect, no resolve operation is performed for the stencil attachment.

If the image subresource range referenced by the depth/stencil attachment is created with `VK_IMAGE_CREATE_SAMPLE_LOCATIONS_COMPATIBLE_DEPTH_BIT_EXT`, then the [multisample resolve operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-resolve-operations) uses the sample locations state specified in the `sampleLocationsInfo` member of the element of the `VkRenderPassSampleLocationsBeginInfoEXT`::`pPostSubpassSampleLocations` for the subpass.

If `pDepthStencilAttachment` is `NULL`, or if its attachment index is `VK_ATTACHMENT_UNUSED`, it indicates that no depth/stencil attachment will be used in the subpass.

The contents of an attachment within the render area become undefined at the start of a subpass **S** if all of the following conditions are true:

- The attachment is used as a color, depth/stencil, or resolve attachment in any subpass in the render pass.
- There is a subpass **S1** that uses or preserves the attachment, and a subpass dependency from **S1** to **S**.
- The attachment is not used or preserved in subpass **S**.

In addition, the contents of an attachment within the render area become undefined at the start of a subpass **S** if all of the following conditions are true:

- `VK_SUBPASS_DESCRIPTION_SHADER_RESOLVE_BIT_QCOM` is set.
- The attachment is used as a color or depth/stencil in the subpass.

Once the contents of an attachment become undefined in subpass **S**, they remain undefined for subpasses in subpass dependency chains starting with subpass **S** until they are written again. However, they remain valid for subpasses in other subpass dependency chains starting with subpass **S1** if those subpasses use or preserve the attachment.

Valid Usage

- [](#VUID-VkSubpassDescription-attachment-06912)VUID-VkSubpassDescription-attachment-06912  
  If the `attachment` member of an element of `pInputAttachments` is not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL`
- [](#VUID-VkSubpassDescription-attachment-06913)VUID-VkSubpassDescription-attachment-06913  
  If the `attachment` member of an element of `pColorAttachments` is not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL`
- [](#VUID-VkSubpassDescription-attachment-06914)VUID-VkSubpassDescription-attachment-06914  
  If the `attachment` member of an element of `pResolveAttachments` is not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL`
- [](#VUID-VkSubpassDescription-attachment-06915)VUID-VkSubpassDescription-attachment-06915  
  If the `attachment` member of `pDepthStencilAttachment` is not `VK_ATTACHMENT_UNUSED`, ts `layout` member **must** not be `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL`
- [](#VUID-VkSubpassDescription-attachment-06916)VUID-VkSubpassDescription-attachment-06916  
  If the `attachment` member of an element of `pColorAttachments` is not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`
- [](#VUID-VkSubpassDescription-attachment-06917)VUID-VkSubpassDescription-attachment-06917  
  If the `attachment` member of an element of `pResolveAttachments` is not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`
- [](#VUID-VkSubpassDescription-attachment-06918)VUID-VkSubpassDescription-attachment-06918  
  If the `attachment` member of an element of `pInputAttachments` is not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`
- [](#VUID-VkSubpassDescription-attachment-06919)VUID-VkSubpassDescription-attachment-06919  
  If the `attachment` member of an element of `pColorAttachments` is not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`, `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkSubpassDescription-attachment-06920)VUID-VkSubpassDescription-attachment-06920  
  If the `attachment` member of an element of `pResolveAttachments` is not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`, `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkSubpassDescription-attachment-06921)VUID-VkSubpassDescription-attachment-06921  
  If the `attachment` member of an element of `pInputAttachments` is not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR`
- [](#VUID-VkSubpassDescription-attachment-06922)VUID-VkSubpassDescription-attachment-06922  
  If the `attachment` member of an element of `pColorAttachments` is not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`
- [](#VUID-VkSubpassDescription-attachment-06923)VUID-VkSubpassDescription-attachment-06923  
  If the `attachment` member of an element of `pResolveAttachments` is not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`
- [](#VUID-VkSubpassDescription-attachment-10755)VUID-VkSubpassDescription-attachment-10755  
  If the `attachment` member of an element of `pResolveAttachments` is not `VK_ATTACHMENT_UNUSED`, the underlying resource must not be bound to a `VkDeviceMemory` object allocated from a `VkMemoryHeap` with the `VK_MEMORY_HEAP_TILE_MEMORY_BIT_QCOM` property.
- [](#VUID-VkSubpassDescription-flags-10683)VUID-VkSubpassDescription-flags-10683  
  If `flags` includes `VK_SUBPASS_DESCRIPTION_TILE_SHADING_APRON_BIT_QCOM`, the render pass **must** have been created with a [VkRenderPassTileShadingCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTileShadingCreateInfoQCOM.html)::`tileApronSize` greater than `(0,0)`
- [](#VUID-VkSubpassDescription-pipelineBindPoint-04952)VUID-VkSubpassDescription-pipelineBindPoint-04952  
  `pipelineBindPoint` **must** be `VK_PIPELINE_BIND_POINT_GRAPHICS` or `VK_PIPELINE_BIND_POINT_SUBPASS_SHADING_HUAWEI`
- [](#VUID-VkSubpassDescription-colorAttachmentCount-00845)VUID-VkSubpassDescription-colorAttachmentCount-00845  
  `colorAttachmentCount` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxColorAttachments`
- [](#VUID-VkSubpassDescription-loadOp-00846)VUID-VkSubpassDescription-loadOp-00846  
  If the first use of an attachment in this render pass is as an input attachment, and the attachment is not also used as a color or depth/stencil attachment in the same subpass, then `loadOp` **must** not be `VK_ATTACHMENT_LOAD_OP_CLEAR`
- [](#VUID-VkSubpassDescription-pResolveAttachments-00847)VUID-VkSubpassDescription-pResolveAttachments-00847  
  If `pResolveAttachments` is not `NULL`, for each resolve attachment that is not `VK_ATTACHMENT_UNUSED`, the corresponding color attachment **must** not be `VK_ATTACHMENT_UNUSED`
- [](#VUID-VkSubpassDescription-pResolveAttachments-00848)VUID-VkSubpassDescription-pResolveAttachments-00848  
  If `pResolveAttachments` is not `NULL`, for each resolve attachment that is not `VK_ATTACHMENT_UNUSED`, the corresponding color attachment **must** not have a sample count of `VK_SAMPLE_COUNT_1_BIT`
- [](#VUID-VkSubpassDescription-pResolveAttachments-00849)VUID-VkSubpassDescription-pResolveAttachments-00849  
  If `pResolveAttachments` is not `NULL`, each resolve attachment that is not `VK_ATTACHMENT_UNUSED` **must** have a sample count of `VK_SAMPLE_COUNT_1_BIT`
- [](#VUID-VkSubpassDescription-pResolveAttachments-00850)VUID-VkSubpassDescription-pResolveAttachments-00850  
  If `pResolveAttachments` is not `NULL`, each resolve attachment that is not `VK_ATTACHMENT_UNUSED` **must** have the same [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) as its corresponding color attachment
- [](#VUID-VkSubpassDescription-pColorAttachments-09430)VUID-VkSubpassDescription-pColorAttachments-09430  
  All attachments in `pColorAttachments` that are not `VK_ATTACHMENT_UNUSED` **must** have the same sample count
- [](#VUID-VkSubpassDescription-pInputAttachments-02647)VUID-VkSubpassDescription-pInputAttachments-02647  
  All attachments in `pInputAttachments` that are not `VK_ATTACHMENT_UNUSED` **must** have image formats whose [potential format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#potential-format-features) contain at least `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT` or `VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-VkSubpassDescription-pColorAttachments-02648)VUID-VkSubpassDescription-pColorAttachments-02648  
  All attachments in `pColorAttachments` that are not `VK_ATTACHMENT_UNUSED` **must** have image formats whose [potential format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#potential-format-features) contain `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT`
- [](#VUID-VkSubpassDescription-pResolveAttachments-02649)VUID-VkSubpassDescription-pResolveAttachments-02649  
  All attachments in `pResolveAttachments` that are not `VK_ATTACHMENT_UNUSED` **must** have image formats whose [potential format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#potential-format-features) contain `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT`
- [](#VUID-VkSubpassDescription-pDepthStencilAttachment-02650)VUID-VkSubpassDescription-pDepthStencilAttachment-02650  
  If `pDepthStencilAttachment` is not `NULL` and the attachment is not `VK_ATTACHMENT_UNUSED` then it **must** have an image format whose [potential format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#potential-format-features) contain `VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-VkSubpassDescription-linearColorAttachment-06496)VUID-VkSubpassDescription-linearColorAttachment-06496  
  If the [`linearColorAttachment`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-linearColorAttachment) feature is enabled and the image is created with `VK_IMAGE_TILING_LINEAR`, all attachments in `pInputAttachments` that are not `VK_ATTACHMENT_UNUSED` **must** have image formats whose [potential format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#potential-format-features) **must** contain `VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV`
- [](#VUID-VkSubpassDescription-linearColorAttachment-06497)VUID-VkSubpassDescription-linearColorAttachment-06497  
  If the [`linearColorAttachment`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-linearColorAttachment) feature is enabled and the image is created with `VK_IMAGE_TILING_LINEAR`, all attachments in `pColorAttachments` that are not `VK_ATTACHMENT_UNUSED` **must** have image formats whose [potential format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#potential-format-features) **must** contain `VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV`
- [](#VUID-VkSubpassDescription-linearColorAttachment-06498)VUID-VkSubpassDescription-linearColorAttachment-06498  
  If the [`linearColorAttachment`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-linearColorAttachment) feature is enabled and the image is created with `VK_IMAGE_TILING_LINEAR`, all attachments in `pResolveAttachments` that are not `VK_ATTACHMENT_UNUSED` **must** have image formats whose [potential format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#potential-format-features) **must** contain `VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV`
- [](#VUID-VkSubpassDescription-None-09431)VUID-VkSubpassDescription-None-09431  
  If either of the following is enabled:
  
  - The `VK_AMD_mixed_attachment_samples` extension
  - The `VK_NV_framebuffer_mixed_samples` extension
  
  all attachments in `pColorAttachments` that are not `VK_ATTACHMENT_UNUSED` **must** have a sample count that is smaller than or equal to the sample count of `pDepthStencilAttachment` if it is not `VK_ATTACHMENT_UNUSED`
- [](#VUID-VkSubpassDescription-pDepthStencilAttachment-01418)VUID-VkSubpassDescription-pDepthStencilAttachment-01418  
  If `pDepthStencilAttachment` is not `VK_ATTACHMENT_UNUSED` and any attachments in `pColorAttachments` are not `VK_ATTACHMENT_UNUSED`, they **must** have the same sample count , if none of the following are enabled:
  
  - The `VK_AMD_mixed_attachment_samples` extension
  - The `VK_NV_framebuffer_mixed_samples` extension
- [](#VUID-VkSubpassDescription-attachment-00853)VUID-VkSubpassDescription-attachment-00853  
  Each element of `pPreserveAttachments` **must** not be `VK_ATTACHMENT_UNUSED`
- [](#VUID-VkSubpassDescription-pPreserveAttachments-00854)VUID-VkSubpassDescription-pPreserveAttachments-00854  
  Each element of `pPreserveAttachments` **must** not also be an element of any other member of the subpass description
- [](#VUID-VkSubpassDescription-layout-02519)VUID-VkSubpassDescription-layout-02519  
  If any attachment is used by more than one [VkAttachmentReference](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference.html) member, then each use **must** use the same `layout`
- [](#VUID-VkSubpassDescription-flags-00856)VUID-VkSubpassDescription-flags-00856  
  If `flags` includes `VK_SUBPASS_DESCRIPTION_PER_VIEW_POSITION_X_ONLY_BIT_NVX`, it **must** also include `VK_SUBPASS_DESCRIPTION_PER_VIEW_ATTRIBUTES_BIT_NVX`
- [](#VUID-VkSubpassDescription-flags-03341)VUID-VkSubpassDescription-flags-03341  
  If `flags` includes `VK_SUBPASS_DESCRIPTION_SHADER_RESOLVE_BIT_QCOM`, and if `pResolveAttachments` is not `NULL`, then each resolve attachment **must** be `VK_ATTACHMENT_UNUSED`
- [](#VUID-VkSubpassDescription-flags-03343)VUID-VkSubpassDescription-flags-03343  
  If `flags` includes `VK_SUBPASS_DESCRIPTION_SHADER_RESOLVE_BIT_QCOM`, then the subpass **must** be the last subpass in a subpass dependency chain
- [](#VUID-VkSubpassDescription-pInputAttachments-02868)VUID-VkSubpassDescription-pInputAttachments-02868  
  If the render pass is created with `VK_RENDER_PASS_CREATE_TRANSFORM_BIT_QCOM` each of the elements of `pInputAttachments` **must** be `VK_ATTACHMENT_UNUSED`
- [](#VUID-VkSubpassDescription-pDepthStencilAttachment-04438)VUID-VkSubpassDescription-pDepthStencilAttachment-04438  
  `pDepthStencilAttachment` and `pColorAttachments` **must** not contain references to the same attachment

Valid Usage (Implicit)

- [](#VUID-VkSubpassDescription-flags-parameter)VUID-VkSubpassDescription-flags-parameter  
  `flags` **must** be a valid combination of [VkSubpassDescriptionFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescriptionFlagBits.html) values
- [](#VUID-VkSubpassDescription-pipelineBindPoint-parameter)VUID-VkSubpassDescription-pipelineBindPoint-parameter  
  `pipelineBindPoint` **must** be a valid [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html) value
- [](#VUID-VkSubpassDescription-pInputAttachments-parameter)VUID-VkSubpassDescription-pInputAttachments-parameter  
  If `inputAttachmentCount` is not `0`, `pInputAttachments` **must** be a valid pointer to an array of `inputAttachmentCount` valid [VkAttachmentReference](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference.html) structures
- [](#VUID-VkSubpassDescription-pColorAttachments-parameter)VUID-VkSubpassDescription-pColorAttachments-parameter  
  If `colorAttachmentCount` is not `0`, `pColorAttachments` **must** be a valid pointer to an array of `colorAttachmentCount` valid [VkAttachmentReference](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference.html) structures
- [](#VUID-VkSubpassDescription-pResolveAttachments-parameter)VUID-VkSubpassDescription-pResolveAttachments-parameter  
  If `colorAttachmentCount` is not `0`, and `pResolveAttachments` is not `NULL`, `pResolveAttachments` **must** be a valid pointer to an array of `colorAttachmentCount` valid [VkAttachmentReference](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference.html) structures
- [](#VUID-VkSubpassDescription-pDepthStencilAttachment-parameter)VUID-VkSubpassDescription-pDepthStencilAttachment-parameter  
  If `pDepthStencilAttachment` is not `NULL`, `pDepthStencilAttachment` **must** be a valid pointer to a valid [VkAttachmentReference](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference.html) structure
- [](#VUID-VkSubpassDescription-pPreserveAttachments-parameter)VUID-VkSubpassDescription-pPreserveAttachments-parameter  
  If `preserveAttachmentCount` is not `0`, `pPreserveAttachments` **must** be a valid pointer to an array of `preserveAttachmentCount` `uint32_t` values

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAttachmentReference](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference.html), [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html), [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo.html), [VkSubpassDescriptionFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescriptionFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSubpassDescription)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0