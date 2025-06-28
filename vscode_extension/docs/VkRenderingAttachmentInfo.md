# VkRenderingAttachmentInfo(3) Manual Page

## Name

VkRenderingAttachmentInfo - Structure specifying attachment information



## [](#_c_specification)C Specification

The `VkRenderingAttachmentInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_3
typedef struct VkRenderingAttachmentInfo {
    VkStructureType          sType;
    const void*              pNext;
    VkImageView              imageView;
    VkImageLayout            imageLayout;
    VkResolveModeFlagBits    resolveMode;
    VkImageView              resolveImageView;
    VkImageLayout            resolveImageLayout;
    VkAttachmentLoadOp       loadOp;
    VkAttachmentStoreOp      storeOp;
    VkClearValue             clearValue;
} VkRenderingAttachmentInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_dynamic_rendering
typedef VkRenderingAttachmentInfo VkRenderingAttachmentInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `imageView` is the image view that will be used for rendering.
- `imageLayout` is the layout that `imageView` will be in during rendering.
- `resolveMode` is a [VkResolveModeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResolveModeFlagBits.html) value defining how data written to `imageView` will be resolved into `resolveImageView`.
- `resolveImageView` is an image view used to write resolved data at the end of rendering.
- `resolveImageLayout` is the layout that `resolveImageView` will be in during rendering.
- `loadOp` is a [VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentLoadOp.html) value defining the [load operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-load-operations) for the attachment.
- `storeOp` is a [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentStoreOp.html) value defining the [store operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-store-operations) for the attachment.
- `clearValue` is a [VkClearValue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearValue.html) structure defining values used to clear `imageView` when `loadOp` is `VK_ATTACHMENT_LOAD_OP_CLEAR`.

## [](#_description)Description

Values in `imageView` are loaded and stored according to the values of `loadOp` and `storeOp`, within the render area for each device specified in [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html). If `imageView` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), and `resolveMode` is not `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_BIT_ANDROID`, other members of this structure are ignored; writes to this attachment will be discarded, and no [load](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-load-operations), [store](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-store-operations), or [multisample resolve](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-resolve-operations) operations will be performed.

If `resolveMode` is `VK_RESOLVE_MODE_NONE`, then `resolveImageView` is ignored. If `resolveMode` is not `VK_RESOLVE_MODE_NONE`, and `resolveImageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), a [render pass multisample resolve operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-resolve-operations) is defined for the attachment subresource. If `resolveMode` is `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_BIT_ANDROID`, and the [`nullColorAttachmentWithExternalFormatResolve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve) limit is `VK_TRUE`, values are only undefined once [load operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-load-operations) have completed.

Note

The resolve mode and store operation are independent; it is valid to write both resolved and unresolved values, and equally valid to discard the unresolved values while writing the resolved ones.

Store and resolve operations are only performed at the end of a render pass instance that does not specify the `VK_RENDERING_SUSPENDING_BIT_KHR` flag.

Load operations are only performed at the beginning of a render pass instance that does not specify the `VK_RENDERING_RESUMING_BIT_KHR` flag.

Image contents at the end of a suspended render pass instance remain defined for access by a resuming render pass instance.

If the [`nullColorAttachmentWithExternalFormatResolve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve) limit is `VK_TRUE`, and `resolveMode` is `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_BIT_ANDROID`, values in the color attachment will be loaded from the resolve attachment at the start of rendering, and **may** also be reloaded any time after a resolve occurs or the resolve attachment is written to; if this occurs it **must** happen-before any writes to the color attachment are performed which happen-after the resolve that triggers this. If any color component in the external format is subsampled, values will be read from the nearest sample in the image when they are loaded.

Valid Usage

- [](#VUID-VkRenderingAttachmentInfo-imageView-06129)VUID-VkRenderingAttachmentInfo-imageView-06129  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and has a non-integer color format, `resolveMode` **must** be `VK_RESOLVE_MODE_NONE` or `VK_RESOLVE_MODE_AVERAGE_BIT`
- [](#VUID-VkRenderingAttachmentInfo-imageView-06130)VUID-VkRenderingAttachmentInfo-imageView-06130  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and has an integer color format, `resolveMode` **must** be `VK_RESOLVE_MODE_NONE` or `VK_RESOLVE_MODE_SAMPLE_ZERO_BIT`
- [](#VUID-VkRenderingAttachmentInfo-imageView-06861)VUID-VkRenderingAttachmentInfo-imageView-06861  
  `imageView` **must** not have a sample count of `VK_SAMPLE_COUNT_1_BIT` if all of the following hold:
  
  - `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
  - `resolveMode` is not `VK_RESOLVE_MODE_NONE`
  - the `pNext` chain of [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html) does not include a [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html) structure with the `multisampledRenderToSingleSampledEnable` field equal to `VK_TRUE`
- [](#VUID-VkRenderingAttachmentInfo-imageView-06862)VUID-VkRenderingAttachmentInfo-imageView-06862  
  `resolveImageView` **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) if all of the following hold:
  
  - `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
  - `resolveMode` is not `VK_RESOLVE_MODE_NONE`
  - the `pNext` chain of [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html) does not include a [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html) structure with the `multisampledRenderToSingleSampledEnable` field equal to `VK_TRUE`
- [](#VUID-VkRenderingAttachmentInfo-imageView-06863)VUID-VkRenderingAttachmentInfo-imageView-06863  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `resolveMode` is not `VK_RESOLVE_MODE_NONE`, the `pNext` chain of [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html) includes a [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html) structure with the `multisampledRenderToSingleSampledEnable` field equal to `VK_TRUE`, and `imageView` has a sample count of `VK_SAMPLE_COUNT_1_BIT`, `resolveImageView` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkRenderingAttachmentInfo-imageView-06864)VUID-VkRenderingAttachmentInfo-imageView-06864  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `resolveImageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), and `resolveMode` is not `VK_RESOLVE_MODE_NONE`, `resolveImageView` **must** have a sample count of `VK_SAMPLE_COUNT_1_BIT`
- [](#VUID-VkRenderingAttachmentInfo-imageView-06865)VUID-VkRenderingAttachmentInfo-imageView-06865  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `resolveImageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), and `resolveMode` is not `VK_RESOLVE_MODE_NONE`, `imageView` and `resolveImageView` **must** have the same [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html)
- [](#VUID-VkRenderingAttachmentInfo-imageView-06135)VUID-VkRenderingAttachmentInfo-imageView-06135  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `imageLayout` **must** not be `VK_IMAGE_LAYOUT_UNDEFINED`, `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL`, `VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL`, `VK_IMAGE_LAYOUT_ZERO_INITIALIZED_EXT`, `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL`, or `VK_IMAGE_LAYOUT_PREINITIALIZED`
- [](#VUID-VkRenderingAttachmentInfo-imageView-06136)VUID-VkRenderingAttachmentInfo-imageView-06136  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and `resolveMode` is not `VK_RESOLVE_MODE_NONE`, `resolveImageLayout` **must** not be `VK_IMAGE_LAYOUT_UNDEFINED`, `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL`, `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL`, `VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL`, `VK_IMAGE_LAYOUT_ZERO_INITIALIZED_EXT`, `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL`, or `VK_IMAGE_LAYOUT_PREINITIALIZED`
- [](#VUID-VkRenderingAttachmentInfo-imageView-06137)VUID-VkRenderingAttachmentInfo-imageView-06137  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and `resolveMode` is not `VK_RESOLVE_MODE_NONE`, `resolveImageLayout` **must** not be `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL` or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkRenderingAttachmentInfo-imageView-06138)VUID-VkRenderingAttachmentInfo-imageView-06138  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `imageLayout` **must** not be `VK_IMAGE_LAYOUT_SHADING_RATE_OPTIMAL_NV`
- [](#VUID-VkRenderingAttachmentInfo-imageView-06139)VUID-VkRenderingAttachmentInfo-imageView-06139  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and `resolveMode` is not `VK_RESOLVE_MODE_NONE`, `resolveImageLayout` **must** not be `VK_IMAGE_LAYOUT_SHADING_RATE_OPTIMAL_NV`
- [](#VUID-VkRenderingAttachmentInfo-imageView-06140)VUID-VkRenderingAttachmentInfo-imageView-06140  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `imageLayout` **must** not be `VK_IMAGE_LAYOUT_FRAGMENT_DENSITY_MAP_OPTIMAL_EXT`
- [](#VUID-VkRenderingAttachmentInfo-imageView-06141)VUID-VkRenderingAttachmentInfo-imageView-06141  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and `resolveMode` is not `VK_RESOLVE_MODE_NONE`, `resolveImageLayout` **must** not be `VK_IMAGE_LAYOUT_FRAGMENT_DENSITY_MAP_OPTIMAL_EXT`
- [](#VUID-VkRenderingAttachmentInfo-imageView-06142)VUID-VkRenderingAttachmentInfo-imageView-06142  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and `resolveMode` is not `VK_RESOLVE_MODE_NONE`, `resolveImageLayout` **must** not be `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`
- [](#VUID-VkRenderingAttachmentInfo-imageView-06143)VUID-VkRenderingAttachmentInfo-imageView-06143  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `imageLayout` **must** not be `VK_IMAGE_LAYOUT_FRAGMENT_SHADING_RATE_ATTACHMENT_OPTIMAL_KHR`
- [](#VUID-VkRenderingAttachmentInfo-imageView-06144)VUID-VkRenderingAttachmentInfo-imageView-06144  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and `resolveMode` is not `VK_RESOLVE_MODE_NONE`, `resolveImageLayout` **must** not be `VK_IMAGE_LAYOUT_FRAGMENT_SHADING_RATE_ATTACHMENT_OPTIMAL_KHR`
- [](#VUID-VkRenderingAttachmentInfo-imageView-10780)VUID-VkRenderingAttachmentInfo-imageView-10780  
  If [feedback loop is enabled](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-feedbackloop) for the attachment identified by `imageView`, then `imageView` **must** have been created with a `usage` value including `VK_IMAGE_USAGE_ATTACHMENT_FEEDBACK_LOOP_BIT_EXT`, either `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT` or `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`, and either `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT` or `VK_IMAGE_USAGE_SAMPLED_BIT`
- [](#VUID-VkRenderingAttachmentInfo-imageView-06145)VUID-VkRenderingAttachmentInfo-imageView-06145  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `imageLayout` **must** not be `VK_IMAGE_LAYOUT_PRESENT_SRC_KHR`
- [](#VUID-VkRenderingAttachmentInfo-imageView-06146)VUID-VkRenderingAttachmentInfo-imageView-06146  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and `resolveMode` is not `VK_RESOLVE_MODE_NONE`, `resolveImageLayout` **must** not be `VK_IMAGE_LAYOUT_PRESENT_SRC_KHR`
- [](#VUID-VkRenderingAttachmentInfo-externalFormatResolve-09323)VUID-VkRenderingAttachmentInfo-externalFormatResolve-09323  
  If the [`externalFormatResolve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-externalFormatResolve) feature is not enabled, `resolveMode` **must** not be `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_BIT_ANDROID`
- [](#VUID-VkRenderingAttachmentInfo-resolveMode-09324)VUID-VkRenderingAttachmentInfo-resolveMode-09324  
  If `resolveMode` is `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_BIT_ANDROID`, `resolveImageView` **must** be a valid image view
- [](#VUID-VkRenderingAttachmentInfo-nullColorAttachmentWithExternalFormatResolve-09325)VUID-VkRenderingAttachmentInfo-nullColorAttachmentWithExternalFormatResolve-09325  
  If the [`nullColorAttachmentWithExternalFormatResolve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve) property is `VK_TRUE` and `resolveMode` is `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_BIT_ANDROID`, `resolveImageView` **must** have been created with an image with a `samples` value of `VK_SAMPLE_COUNT_1_BIT`
- [](#VUID-VkRenderingAttachmentInfo-resolveMode-09326)VUID-VkRenderingAttachmentInfo-resolveMode-09326  
  If `resolveMode` is `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_BIT_ANDROID`, `resolveImageView` **must** have been created with an external format specified by [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFormatANDROID.html)
- [](#VUID-VkRenderingAttachmentInfo-resolveMode-09327)VUID-VkRenderingAttachmentInfo-resolveMode-09327  
  If `resolveMode` is `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_BIT_ANDROID`, `resolveImageView` **must** have been created with a `subresourceRange.layerCount` of `1`
- [](#VUID-VkRenderingAttachmentInfo-resolveMode-09328)VUID-VkRenderingAttachmentInfo-resolveMode-09328  
  If `resolveMode` is `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_BIT_ANDROID` and [`nullColorAttachmentWithExternalFormatResolve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve) is `VK_TRUE`, `imageView` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkRenderingAttachmentInfo-resolveMode-09329)VUID-VkRenderingAttachmentInfo-resolveMode-09329  
  If `resolveMode` is `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_BIT_ANDROID` and [`nullColorAttachmentWithExternalFormatResolve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve) is `VK_FALSE`, `imageView` **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html)
- [](#VUID-VkRenderingAttachmentInfo-resolveMode-09330)VUID-VkRenderingAttachmentInfo-resolveMode-09330  
  If `resolveMode` is `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_BIT_ANDROID` and [`nullColorAttachmentWithExternalFormatResolve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve) is `VK_FALSE`, `imageView` **must** have a format equal to the value of [VkAndroidHardwareBufferFormatResolvePropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferFormatResolvePropertiesANDROID.html)::`colorAttachmentFormat` as returned by a call to [vkGetAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAndroidHardwareBufferPropertiesANDROID.html) for the Android hardware buffer that was used to create `resolveImageView`
- [](#VUID-VkRenderingAttachmentInfo-resolveImageView-10728)VUID-VkRenderingAttachmentInfo-resolveImageView-10728  
  If `resolveImageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the underlying resource must not be bound to a `VkDeviceMemory` object allocated from a `VkMemoryHeap` with the `VK_MEMORY_HEAP_TILE_MEMORY_BIT_QCOM` property.

Valid Usage (Implicit)

- [](#VUID-VkRenderingAttachmentInfo-sType-sType)VUID-VkRenderingAttachmentInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDERING_ATTACHMENT_INFO`
- [](#VUID-VkRenderingAttachmentInfo-pNext-pNext)VUID-VkRenderingAttachmentInfo-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkAttachmentFeedbackLoopInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentFeedbackLoopInfoEXT.html)
- [](#VUID-VkRenderingAttachmentInfo-sType-unique)VUID-VkRenderingAttachmentInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkRenderingAttachmentInfo-imageView-parameter)VUID-VkRenderingAttachmentInfo-imageView-parameter  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `imageView` **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) handle
- [](#VUID-VkRenderingAttachmentInfo-imageLayout-parameter)VUID-VkRenderingAttachmentInfo-imageLayout-parameter  
  `imageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-VkRenderingAttachmentInfo-resolveMode-parameter)VUID-VkRenderingAttachmentInfo-resolveMode-parameter  
  If `resolveMode` is not `0`, `resolveMode` **must** be a valid [VkResolveModeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResolveModeFlagBits.html) value
- [](#VUID-VkRenderingAttachmentInfo-resolveImageView-parameter)VUID-VkRenderingAttachmentInfo-resolveImageView-parameter  
  If `resolveImageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `resolveImageView` **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) handle
- [](#VUID-VkRenderingAttachmentInfo-resolveImageLayout-parameter)VUID-VkRenderingAttachmentInfo-resolveImageLayout-parameter  
  `resolveImageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-VkRenderingAttachmentInfo-loadOp-parameter)VUID-VkRenderingAttachmentInfo-loadOp-parameter  
  `loadOp` **must** be a valid [VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentLoadOp.html) value
- [](#VUID-VkRenderingAttachmentInfo-storeOp-parameter)VUID-VkRenderingAttachmentInfo-storeOp-parameter  
  `storeOp` **must** be a valid [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentStoreOp.html) value
- [](#VUID-VkRenderingAttachmentInfo-clearValue-parameter)VUID-VkRenderingAttachmentInfo-clearValue-parameter  
  `clearValue` **must** be a valid [VkClearValue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearValue.html) union
- [](#VUID-VkRenderingAttachmentInfo-commonparent)VUID-VkRenderingAttachmentInfo-commonparent  
  Both of `imageView`, and `resolveImageView` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentLoadOp.html), [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentStoreOp.html), [VkClearValue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearValue.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html), [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html), [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html), [VkResolveModeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResolveModeFlagBits.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRenderingAttachmentInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0